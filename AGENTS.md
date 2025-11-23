# ScriptureBot Architecture & Guidelines

## Architecture Overview

ScriptureBot is a Go-based bot application designed to provide Bible passages and related resources. It is built on top of `github.com/julwrites/BotPlatform`.

### Directory Structure

- `pkg/app`: Contains the core application logic.
  - `passage.go`: Currently handles Bible passage retrieval via web scraping (classic.biblegateway.com).
  - `devo*.go`: Handles devotionals.
  - `command.go`: Command handling logic.
- `pkg/bot`: Contains bot interface implementations (e.g., Telegram).
- `pkg/utils`: Shared utility functions.

### Key Dependencies

- `github.com/julwrites/BotPlatform`: The underlying bot framework.
- `golang.org/x/net/html`: Used for parsing HTML (currently used for scraping).
- `cloud.google.com/go/datastore`: Used for data persistence.

## Development Guidelines

- **Passage Retrieval**: The current scraping mechanism in `pkg/app/passage.go` is being replaced by a new Bible AI API service.
- **New Features**:
  - Word Search: Search for words in the Bible.
  - Bible Query: Ask questions using natural language (LLM-backed).
- **Code Style**: Follow standard Go idioms. Ensure error handling is robust.

## API Integration

The new Bible AI API exposes a `/query` endpoint.
- **Verses**: `query.verses`
- **Word Search**: `query.words`
- **Prompt/Query**: `query.prompt`

Refer to `openapi.yaml` for the full specification.
