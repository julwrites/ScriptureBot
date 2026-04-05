# AI Agent Instructions

You are an expert Software Engineer working on this project. Your primary responsibility is to implement features and fixes while strictly adhering to the Agent Harness workflows.

## Core Directives
- **API First**: The Bible AI API is the primary source for data. Scraping (`pkg/app/passage.go` fallback) is deprecated and should be avoided for new features.
- **Secrets**: Do not commit secrets. Use `pkg/secrets` to retrieve them from Environment or Google Secret Manager.
- **Testing**: Run tests from the root using `go test ./pkg/...`.

## Code Guidelines
- **Go Version**: 1.24+
- **Naming**:
  - Variables: `camelCase`
  - Functions: `PascalCase` (exported), `camelCase` (internal)
  - Packages: `underscore_case`
- **Structure**:
  - `pkg/app`: Business logic.
  - `pkg/bot`: Platform integration.
  - `pkg/utils`: Shared utilities.

## Local Development
- **Setup**: Create a `.env` file with `TELEGRAM_ID` and `TELEGRAM_ADMIN_ID`.
- **Run**: `go run main.go`
- **Testing**: Use `ngrok` to tunnel webhooks or send mock HTTP requests.
