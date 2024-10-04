package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl"
	"github.com/spiffe/spire-plugin-sdk/pluginmain"
	"github.com/spiffe/spire-plugin-sdk/pluginsdk"
	credentialcomposerv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/server/credentialcomposer/v1"
	configv1 "github.com/spiffe/spire-plugin-sdk/proto/spire/service/common/config/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

// Plugin implements the CredentialComposer plugin for SPIRE
type Plugin struct {
	credentialcomposerv1.UnimplementedCredentialComposerServer
	configv1.UnimplementedConfigServer

	configMtx sync.RWMutex
	config    *Config
	logger    hclog.Logger
}

// Config is used to configure the plugin
type Config struct {
}

// Ensure Plugin implements the necessary interfaces
var (
	_ pluginsdk.NeedsLogger       = (*Plugin)(nil)
	_ pluginsdk.NeedsHostServices = (*Plugin)(nil)
)

func (p *Plugin) ComposeWorkloadJWTSVID(ctx context.Context, req *credentialcomposerv1.ComposeWorkloadJWTSVIDRequest) (*credentialcomposerv1.ComposeWorkloadJWTSVIDResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	// Get the configuration
	config, err := p.getConfig()
	if err != nil {
		return nil, err
	}

	// Silence the linter by using the config variable
	_ = config

	// Access the SPIFFE ID of the workload
	spiffeID := req.SpiffeId
	if spiffeID == "" {
		return nil, status.Error(codes.InvalidArgument, "SPIFFE ID is missing in the request")
	}

	// Extract the trust domain and workload path from the SPIFFE ID
	trustDomain, workloadPath, err := parseTrustDomainAndWorkload(spiffeID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to parse trust domain and workload from SPIFFE ID: %v", err)
	}

	// Log the extracted information for visibility
	p.logger.Info("Retrieved workload information", "spiffe_id", spiffeID, "trust_domain", trustDomain, "workload", workloadPath)

	// Extract and modify claims
	updatedClaims, err := updateClaims(req.Attributes.GetClaims(), map[string]interface{}{
		"spiffe-id":    spiffeID,
		"trust-domain": trustDomain,
		"workload":     workloadPath,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create updated claims: %v", err)
	}

	// Create and return the response with the modified claims
	return &credentialcomposerv1.ComposeWorkloadJWTSVIDResponse{
		Attributes: &credentialcomposerv1.JWTSVIDAttributes{
			Claims: updatedClaims,
		},
	}, nil
}

// updateClaims adds or updates the provided key-value pairs in the given claims
func updateClaims(claims *structpb.Struct, updates map[string]interface{}) (*structpb.Struct, error) {
	if claims == nil {
		claims = &structpb.Struct{}
	}
	claimsMap := claims.AsMap()

	// Apply updates to the claims map
	for key, value := range updates {
		claimsMap[key] = value
	}

	// Convert the updated map back to a structpb.Struct
	return structpb.NewStruct(claimsMap)
}

// parseTrustDomainAndWorkload extracts the trust domain and workload path from a SPIFFE ID.
func parseTrustDomainAndWorkload(spiffeID string) (string, string, error) {
	// Check if the SPIFFE ID has the correct prefix.
	if !strings.HasPrefix(spiffeID, "spiffe://") {
		return "", "", fmt.Errorf("invalid SPIFFE ID format")
	}

	// Remove the "spiffe://" prefix.
	spiffeID = strings.TrimPrefix(spiffeID, "spiffe://")

	// Split the SPIFFE ID into trust domain and workload path.
	idParts := strings.SplitN(spiffeID, "/", 2)
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return "", "", fmt.Errorf("SPIFFE ID must contain both a trust domain and a workload path")
	}

	trustDomain := idParts[0]
	workloadPath := idParts[1]

	return trustDomain, workloadPath, nil
}

// Configure configures the plugin when it is first loaded
func (p *Plugin) Configure(ctx context.Context, req *configv1.ConfigureRequest) (*configv1.ConfigureResponse, error) {
	config := new(Config)
	if err := hcl.Decode(config, req.HclConfiguration); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to decode configuration: %v", err)
	}

	// Update configuration
	p.setConfig(config)
	return &configv1.ConfigureResponse{}, nil
}

// BrokerHostServices is called to obtain clients to SPIRE host services
func (p *Plugin) BrokerHostServices(broker pluginsdk.ServiceBroker) error {
	return nil
}

// SetLogger sets the logger provided by SPIRE to the plugin
func (p *Plugin) SetLogger(logger hclog.Logger) {
	p.logger = logger
}

// setConfig replaces the configuration atomically under a write lock
func (p *Plugin) setConfig(config *Config) {
	p.configMtx.Lock()
	defer p.configMtx.Unlock()
	p.config = config
}

// getConfig gets the configuration under a read lock
func (p *Plugin) getConfig() (*Config, error) {
	p.configMtx.RLock()
	defer p.configMtx.RUnlock()
	if p.config == nil {
		return nil, status.Error(codes.FailedPrecondition, "not configured")
	}
	return p.config, nil
}

func main() {
	plugin := new(Plugin)
	pluginmain.Serve(
		credentialcomposerv1.CredentialComposerPluginServer(plugin),
		configv1.ConfigServiceServer(plugin),
	)
}
