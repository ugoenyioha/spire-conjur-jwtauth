# spire-conjur-jwtauth <!-- omit in toc -->

[![Go Report Card](https://goreportcard.com/badge/github.com/infamousjoeg/spire-conjur-jwtauth)](https://goreportcard.com/report/github.com/infamousjoeg/spire-conjur-jwtauth)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://github.com/infamousjoeg/spire-conjur-jwtauth/actions/workflows/build.yml/badge.svg)](https://github.com/infamousjoeg/spire-conjur-jwtauth/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/infamousjoeg/spire-conjur-jwtauth.svg)](https://pkg.go.dev/github.com/infamousjoeg/spire-conjur-jwtauth)
[![Coverage Status](https://coveralls.io/repos/github/infamousjoeg/spire-conjur-jwtauth/badge.svg?branch=main)](https://coveralls.io/github/infamousjoeg/spire-conjur-jwtauth?branch=main)

## Table of Contents <!-- omit in toc -->

- [Overview](#overview)
- [What is SPIFFE and SPIRE?](#what-is-spiffe-and-spire)
  - [SPIFFE (Secure Production Identity Framework For Everyone)](#spiffe-secure-production-identity-framework-for-everyone)
  - [SPIRE (SPIFFE Runtime Environment)](#spire-spiffe-runtime-environment)
- [CredentialComposer Plugin](#credentialcomposer-plugin)
  - [What is a CredentialComposer?](#what-is-a-credentialcomposer)
  - [Why Use a CredentialComposer?](#why-use-a-credentialcomposer)
  - [How This Plugin Works](#how-this-plugin-works)
  - [Installation and Configuration](#installation-and-configuration)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [Usage](#usage)
- [Related Documentation](#related-documentation)
- [Development](#development)
  - [Running Unit Tests](#running-unit-tests)
  - [Contributing](#contributing)
- [License](#license)


## Overview

This repository contains a `CredentialComposer` plugin for SPIRE that extends the functionality of JWT-SVIDs, adding custom claims based on the workload's SPIFFE identity. Specifically, this plugin adds claims such as `"spiffe-id"`, `"trust-domain"`, and `"workload"` to the JWT-SVID, allowing better integration with systems that require additional identity information.

## What is SPIFFE and SPIRE?

### SPIFFE (Secure Production Identity Framework For Everyone)
[**SPIFFE**](https://spiffe.io/) is an open standard for securely identifying and authenticating services in dynamic and heterogeneous environments. It defines the **SPIFFE ID** format (`spiffe://<trust-domain>/<workload-path>`) that uniquely identifies a workload within a trust domain.

### SPIRE (SPIFFE Runtime Environment)
[**SPIRE**](https://spiffe.io/spire/) is an open-source implementation of the SPIFFE standards. It provides a secure identity infrastructure for workloads by managing the issuance of **SPIFFE IDs** and their corresponding credentials, such as **X.509 SVIDs** and **JWT-SVIDs**.

- **SPIFFE ID**: A unique identifier for workloads, used to establish and enforce trust.
- **SVID (SPIFFE Verifiable Identity Document)**: Represents the credentials associated with a SPIFFE ID, which can be an **X.509 certificate** or a **JWT**.

SPIRE's main role is to assign SPIFFE identities to workloads, manage trust relationships, and provide a robust mechanism for workload attestation.

## CredentialComposer Plugin

### What is a CredentialComposer?
A **CredentialComposer** is a plugin interface in SPIRE that allows users to customize the **SVIDs** issued by SPIRE, such as adding additional metadata or modifying the claims. Specifically, a `CredentialComposer` can be used to add custom claims to **JWT-SVIDs** that SPIRE issues to workloads.

The goal of a `CredentialComposer` plugin is to extend the default attributes associated with a SPIFFE ID, allowing workloads to convey additional information that may be required for specific authorization or logging purposes.

### Why Use a CredentialComposer?
The `CredentialComposer` in this repository is used to:
- Add claims such as `"spiffe-id"`, `"trust-domain"`, and `"workload"` to the **JWT-SVID**.
- Enable downstream services to utilize these claims for decision-making, logging, or visibility.
- Extend the standard attributes of the **JWT-SVID** issued by SPIRE.

### How This Plugin Works
This `CredentialComposer` plugin retrieves information from the workload's SPIFFE ID and uses it to create additional claims:
- **SPIFFE ID (`spiffe-id`)**: The full SPIFFE identifier for the workload.
- **Trust Domain (`trust-domain`)**: The trust domain extracted from the SPIFFE ID.
- **Workload Path (`workload`)**: The workload path extracted from the SPIFFE ID.

The plugin is implemented using the SPIRE Plugin SDK and integrates seamlessly with SPIRE Server to enhance the information conveyed by **JWT-SVIDs**.

### Installation and Configuration

#### Prerequisites
- **SPIRE**: A running SPIRE server that you want to extend using this plugin. [Get Started with SPIRE](https://spiffe.io/spire/getting-started/).
- **Go**: This plugin is written in Go, and Go v1.20+ is required to build the plugin.

#### Installation
1. **Clone this repository**:
   ```sh
   git clone https://github.com/infamousjoeg/spire-conjur-jwtauth.git
   ```
2. **Build the Plugin**:
   ```sh
   go build -o credentialcomposer-plugin
   ```
   This command will create an executable binary that can be used as the plugin.

3. **Update SPIRE Server Configuration**:
   Add the plugin configuration to your SPIRE server's `server.conf` file:

   ```hcl
   plugins {
       CredentialComposer "conjur_jwtauth_composer" {
           plugin_data {
               command = "/path/to/credentialcomposer-plugin"
           }
       }
   }
   ```

4. **Restart the SPIRE Server** to load the new plugin configuration:
   ```sh
   systemctl restart spire-server
   ```

### Usage

Once configured, this plugin will modify the **JWT-SVIDs** issued by SPIRE by adding the custom claims. You can verify the additional claims by checking the contents of the JWT-SVID:

- **SPIFFE ID**: The unique identity of the workload.
- **Trust Domain**: Represents the logical security domain that the workload belongs to.
- **Workload**: Provides more granular information about the specific workload.

These additional claims can be used for:
- **Authorization Policies**: Enabling fine-grained access controls by including specific workload and domain information.
- **Logging and Auditing**: Providing richer identity information for observability and auditing.

## Related Documentation

- [**SPIFFE Documentation**](https://spiffe.io/docs/): Learn more about SPIFFE and its specifications.
- [**SPIRE Documentation**](https://spiffe.io/spire/docs/): Get in-depth information about how SPIRE works, how to deploy it, and how to write plugins.
- [**SPIRE Plugin SDK**](https://github.com/spiffe/spire-plugin-sdk): The SDK used to write SPIRE plugins, including **CredentialComposers**.
- [**SPIFFE JWT-SVID Standard**](https://github.com/spiffe/spiffe/blob/main/standards/JWT-SVID.md): Information on the **JWT-SVID** format and standard claims.

## Development

### Running Unit Tests
The repository includes unit tests to validate the functionality of the plugin. Run the tests with:

```sh
go test -v ./...
```

The tests include:
- Validation of **SPIFFE ID** parsing.
- Verification of correct claims added to the **JWT-SVID**.
- Handling of edge cases and invalid SPIFFE IDs.

### Contributing
Contributions are welcome! Please feel free to submit issues or pull requests for any changes, additions, or suggestions you have.

## License
This project is licensed under the [MIT License](LICENSE).

---

This plugin extends SPIRE's capability to provide richer identity claims in JWT-SVIDs, making it a powerful addition for environments that require custom identity metadata for authorization and auditing purposes. For more information about SPIFFE and SPIRE, please visit the official [SPIFFE Project](https://spiffe.io/).