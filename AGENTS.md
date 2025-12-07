# ScriptureBot Developer Guide

**CURRENT STATUS: MAINTENANCE MODE**

## Helper Scripts
- `scripts/tasks.py`: Manage development tasks.
  - `python3 scripts/tasks.py list`: List tasks.
  - `python3 scripts/tasks.py create <category> <title>`: Create a task.
  - `python3 scripts/tasks.py update <id> <status>`: Update task status.

## Documentation
- `docs/architecture/`: System architecture and directory structure.
- `docs/features/`: Feature specifications.
- `docs/tasks/`: Active and pending tasks.

## Project Specific Instructions

### Core Directives
- **API First**: The Bible AI API is the primary source for data. Scraping (`pkg/app/passage.go` fallback) is deprecated and should be avoided for new features.
- **Secrets**: Do not commit secrets. Use `pkg/secrets` to retrieve them from Environment or Google Secret Manager.
- **Testing**: Run tests from the root using `go test ./pkg/...`.

### Code Guidelines
- **Go Version**: 1.24+
- **Naming**:
  - Variables: `camelCase`
  - Functions: `PascalCase` (exported), `camelCase` (internal)
  - Packages: `underscore_case`
- **Structure**:
  - `pkg/app`: Business logic.
  - `pkg/bot`: Platform integration.
  - `pkg/utils`: Shared utilities.

### Local Development
- **Setup**: Create a `.env` file with `TELEGRAM_ID` and `TELEGRAM_ADMIN_ID`.
- **Run**: `go run main.go`
- **Testing**: Use `ngrok` to tunnel webhooks or send mock HTTP requests.
