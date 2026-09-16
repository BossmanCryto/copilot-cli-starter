# Copilot CLI - Agent Guidelines

This CLI is built for both humans and AI agents with predictable, machine-readable behavior.

## Core Principles

1. **One Goal Per PR**: Each pull request should have a single, clear objective
2. **Contract First**: Establish the current contract before choosing a solution
3. **YAGNI**: Don't add features that aren't needed yet
4. **Reuse Machinery**: Leverage existing project infrastructure
5. **Root Cause Fixes**: Fix at the narrowest shared boundary
6. **Smallest Complete Change**: Ship minimal but complete functionality

## Source Layout

```
.
├── cmd/                    # Command implementations
│   ├── auth/              # Authentication commands
│   ├── config/            # Configuration commands
│   └── core/              # Core commands
├── pkg/                   # Internal packages
│   ├── api/               # API client
│   ├── config/            # Configuration management
│   └── utils/             # Utilities
├── skills/                # AI agent skills
│   ├── workflows/         # Workflow automation
│   └── catalog/           # Skill catalog
├── shortcuts/             # Domain-specific shortcuts
│   └── domains/           # Domain implementations
├── main.go                # Entry point
└── Makefile               # Build automation
```

## Build & Test

```bash
# Build locally
make build

# Run tests
make test

# Test with network
make live-skills-test
```

## For AI Agents

- Focus on predictable, machine-readable output
- Minimize unrelated cleanup
- Keep PRs focused and atomic
- Preserve contracts unless explicitly breaking
- Write regression tests for behavioral changes
- Document decision ceilings and expansion conditions
