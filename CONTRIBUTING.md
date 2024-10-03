# Contributing to spire-conjur-jwtauth <!-- omit in toc -->

Thank you for considering contributing to the **spire-conjur-jwtauth** project! Contributions are what make the open-source community thrive, and we greatly value and welcome any help you can offer. Whether it's reporting bugs, suggesting features, writing code, or improving documentation, all contributions are appreciated.

## Table of Contents <!-- omit in toc -->

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Features](#suggesting-features)
  - [Submitting Code Changes](#submitting-code-changes)
  - [Improving Documentation](#improving-documentation)
- [Development Workflow](#development-workflow)
  - [Setting Up Your Environment](#setting-up-your-environment)
  - [Commit Message Guidelines](#commit-message-guidelines)
  - [Pull Request Guidelines](#pull-request-guidelines)
- [Style Guide](#style-guide)
- [License](#license)

## Code of Conduct

To ensure a welcoming environment for all contributors, we have adopted a Code of Conduct. Please read and adhere to the [Code of Conduct](https://github.com/infamousjoeg/spire-conjur-jwtauth/CODE_OF_CONDUCT.md) before participating in this project.

## How Can I Contribute?

### Reporting Bugs

If you find a bug, please open an issue. Before you do so, please make sure the issue hasn't already been reported by searching through the [existing issues](https://github.com/infamousjoeg/spire-conjur-jwtauth/issues).

When creating a new bug report, include as much information as possible:
- A detailed description of the issue.
- Steps to reproduce the issue.
- Expected and actual behavior.
- The environment (OS, version, etc.) in which you experienced the problem.

### Suggesting Features

We are always open to new ideas! To suggest a new feature:
1. Open a [new issue](https://github.com/infamousjoeg/spire-conjur-jwtauth/issues).
2. Use the "Feature Request" template.
3. Provide a detailed explanation of your suggestion and its use case.

### Submitting Code Changes

If you would like to make a code contribution:
1. **Fork the repository** to your own GitHub account.
2. **Clone your fork** to your local machine.
3. **Create a branch** for your changes:
   ```sh
   git checkout -b feature/my-new-feature
   ```
4. **Make your changes**, and **commit** them with descriptive commit messages.
5. **Run the tests** to make sure everything is still working properly:
   ```sh
   go test ./...
   ```
6. **Push your branch** to your GitHub repository.
7. **Open a Pull Request** (PR) to the main repository. Provide a description of your changes and the reasons behind them.

We strive to review and respond to Pull Requests in a timely manner. Please be patient, and feel free to engage in discussions about your PR.

### Improving Documentation

Improving documentation is one of the easiest yet most impactful ways to contribute. Whether you see a typo, outdated information, or missing details, your input is valuable! You can either:
- Create an issue to highlight the problem.
- Directly edit the documentation and submit a Pull Request.

## Development Workflow

### Setting Up Your Environment

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/infamousjoeg/spire-conjur-jwtauth.git
   ```
2. **Install Dependencies**:
   This project requires [Go](https://golang.org/doc/install) v1.20+.
3. **Build the Plugin**:
   ```sh
   go build -o credentialcomposer-plugin
   ```
4. **Run Tests**:
   Ensure that all unit tests pass before making a Pull Request:
   ```sh
   go test -v ./...
   ```

### Commit Message Guidelines

Please use clear and descriptive commit messages. Format your commit messages as follows:
- Use the present tense ("Add feature" not "Added feature").
- Start with a capital letter.
- Be concise but descriptive.

### Pull Request Guidelines

- Make sure your code is thoroughly tested.
- Ensure your code follows the project’s style guidelines.
- Link related issues in your Pull Request description.
- Address feedback promptly during code review.

## Style Guide

- **Formatting**: This project uses `gofmt` for code formatting. Run `gofmt` before committing your changes.
- **Linting**: Lint your code using `golangci-lint` to catch potential issues before submitting a Pull Request.

## License

By contributing to this project, you agree that your contributions will be licensed under the [Apache License 2.0](LICENSE).