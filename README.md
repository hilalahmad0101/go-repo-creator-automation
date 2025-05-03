# GitHub Repository Creation Tool

A simple Go tool to automate the creation of GitHub repositories via API.

## Overview

This tool allows you to quickly create new GitHub repositories with customizable settings, either under your personal account or within an organization. It uses the GitHub API through the official Go client library.

## Features

- Create repositories in personal accounts or organizations
- Set repository visibility (public/private)
- Initialize with README file
- Add description automatically
- Simple environment variable configuration

## Prerequisites

- Go 1.16 or higher
- GitHub Personal Access Token with appropriate permissions (repo, admin:org)
- Basic understanding of Go and GitHub

## Installation

1. Clone this repository or copy the source code
2. Install the required dependencies:

```bash
go get github.com/google/go-github/v50/github
go get golang.org/x/oauth2
```

## Configuration

This tool uses environment variables for configuration:

| Environment Variable | Description | Required | Default |
|---------------------|-------------|----------|---------|
| `GITHUB_TOKEN` | Your GitHub Personal Access Token | Yes | - |
| `REPO_NAME` | Name for the new repository | Yes | - |
| `GITHUB_ORG` | Organization name (leave empty for personal account) | No | empty (personal account) |
| `REPO_PRIVATE` | Set to "true" for private repository | No | false (public) |

## Usage

1. Set the required environment variables:

```bash
export GITHUB_TOKEN="your_github_token"
export REPO_NAME="my-new-repo"
export GITHUB_ORG="your-organization" # Optional
export REPO_PRIVATE="true" # Optional
```

2. Run the tool:

```bash
go run main.go
```

3. On successful execution, the URL of your newly created repository will be printed to the console.

## Example

Creating a private repository under your personal account:

```bash
export GITHUB_TOKEN="ghp_1234567890abcdef"
export REPO_NAME="my-awesome-project"
export REPO_PRIVATE="true"
go run main.go
```

Creating a public repository under an organization:

```bash
export GITHUB_TOKEN="ghp_1234567890abcdef"
export REPO_NAME="team-project"
export GITHUB_ORG="my-organization"
go run main.go
```

## Error Handling

The tool provides basic error messages if repository creation fails. Common issues include:
- Invalid token
- Insufficient permissions
- Repository name already exists
- Invalid organization name

## Extending the Tool

To extend this tool with additional functionality:

1. Explore the [go-github documentation](https://pkg.go.dev/github.com/google/go-github/v50/github) for additional API capabilities
2. Add more repository settings in the `repo` object
3. Implement additional error handling for more robust operation
4. Consider adding command-line flags as an alternative to environment variables

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This tool is provided as-is without any guarantees. Use at your own risk.
