# Copilot Instructions for copilot-cli-starter

## Overview

This is an AI-ready CLI tool built with Copilot support. The project follows predictable, machine-readable patterns optimized for AI agent workflows.

## Key Files

- **AGENTS.md** - Core guidelines for contributions and agent workflows
- **main.go** - Entry point
- **cmd/** - Command implementations
- **skills/** - AI agent skills
- **Makefile** - Build and test automation

## Quick Start

1. Clone the repository
2. Run `make build` to compile
3. Run `make test` to verify

## Contributing

When making changes:

1. Keep PRs focused on a single goal
2. Maintain machine-readable behavior
3. Add regression tests for changes
4. Document decision ceilings
5. Preserve existing contracts

## Agent Workflows

This CLI supports AI agent skills through the `skills/` directory. Skills are composable units of functionality that agents can invoke.

## Testing

- `make test` - Run unit tests
- `make live-skills-test` - Run tests with network access
