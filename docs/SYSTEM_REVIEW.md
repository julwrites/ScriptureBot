# ScriptureBot System Review

## 1. Architectural Review

### Overview
ScriptureBot is a Telegram bot hosted on Google Cloud Run, designed to provide Bible passages, search functionality, and AI-driven answers. It operates as a client to the **Bible AI API** while leveraging the **BotPlatform** framework for platform-agnostic bot logic.

### Components
- **Frontend (Telegram)**: The user interface is the Telegram chat.
- **Application Server (ScriptureBot)**:
  - Written in Go.
  - Hosted on **Google Cloud Run**.
  - Serves as the middle layer, processing natural language and routing commands.
- **Framework (BotPlatform)**:
  - A library (`github.com/julwrites/BotPlatform`) that handles Telegram API communication, user session management, and message translation.
- **Backend Service (BibleAIAPI)**:
  - An external API service used for retrieving Bible verses, searching, and AI "Ask" functionality.
- **Database (Firestore)**:
  - Used for persisting user data and configuration. Integration is handled via `BotPlatform` and local utilities.
- **Secrets Management**:
  - Uses **Google Secret Manager** for production secrets (`TELEGRAM_ID`, `TELEGRAM_ADMIN_ID`, etc.).
  - Supports local environment variables (`.env`) for development.
  - *Note*: The `deployment.yml` generates a `secrets.yaml` file which appears to be unused by the application logic, as `pkg/secrets` relies on `godotenv` (for `.env`) or Google Secret Manager.

### Data Flow
1. **Ingestion**: Telegram Webhook -> `main.go` -> `bot.TelegramHandler`.
2. **Translation**: `BotPlatform` translates the raw HTTP request into a standardized `SessionData` object.
3. **Logic**: `app.ProcessCommand` determines the intent:
   - **Passage**: Calls `app.GetBiblePassage`.
   - **Search**: Calls `app.GetBibleSearch`.
   - **Ask**: Calls `app.GetBibleAsk` (Admin only) or `GetBibleAskWithContext`.
4. **External Call**:
   - Primary: `app.SubmitQuery` sends JSON requests to `BibleAIAPI`.
   - Fallback: `app.GetBiblePassageFallback` scrapes `classic.biblegateway.com` if the API fails for passages.
5. **Response**: The `SessionData` with the response message is sent back to Telegram via `BotPlatform`.

---

## 2. Integration Review

### BotPlatform Integration
- **Coupling**: The application is tightly coupled with `BotPlatform`. It imports `pkg/platform` and `pkg/def` heavily.
- **Session Management**: It relies on `BotPlatform` to manage user state (`env.User`).
- **Formatting**: HTML formatting functions (e.g., `platform.TelegramBold`) are used in `pkg/app/passage.go`.
- **Versioning**: `go.mod` pins `BotPlatform` to a specific commit. Changes in `BotPlatform` would require updates here.

### Bible AI API Integration
- **Client**: Custom client implementation in `pkg/app/api_client.go`.
- **Protocol**: JSON over HTTP.
- **Authentication**: Uses `BIBLE_API_KEY` from secrets.
- **Endpoints**:
  - `/query` for Verse retrieval, Word Search, and AI Prompts.
- **Resilience**: The system implements a fallback to web scraping (`GetBiblePassageFallback`) for passage retrieval, ensuring basic functionality remains even if the API is down.

### Infrastructure Integration
- **Google Cloud**:
  - **Cloud Run**: Hosting.
  - **Secret Manager**: Configuration.
  - **Artifact Registry**: Docker image storage.
- **CI/CD**:
  - GitHub Actions pipelines for `Build & Test` (`automation.yml`) and `Build, Stage & Deploy` (`deployment.yml`).
  - Deployment automatically triggers a webhook update via `cmd/webhook`.

---

## 3. Code Review

### Structure
- The code is well-organized into logical packages:
  - `pkg/app`: Core business logic (Passage, Search, Ask).
  - `pkg/bot`: Telegram handler and bot logic wiring.
  - `pkg/secrets`: Secrets retrieval strategy.
  - `pkg/utils`: Helper functions (HTML parsing, Database).

### Key Observations & Findings
- **Secrets Confusion**: `deployment.yml` generates a `secrets.yaml` file and copies it into the Docker image, but the application code (`pkg/secrets`) does not appear to read this file format. It relies on environment variables or Google Secret Manager. This `secrets.yaml` might be dead code.
- **Search Output Format**: The `GetBibleSearch` function in `pkg/app/search.go` formats results as a simple text list (`- Verse`). However, project documentation/memory suggests it should return HTML anchor tags. This indicates a potential regression or incomplete implementation.
- **HTML Parsing**: `pkg/app/passage.go` contains robust HTML parsing logic to convert BibleGateway HTML into Telegram-friendly format, preserving superscripts and formatting.
- **Fuzzy Matching**: `pkg/app/bible_reference.go` implements a sophisticated Levenshtein distance algorithm for identifying Bible book names, handling misspellings and abbreviations effectively.
- **Access Control**: The `/ask` command (`GetBibleAsk`) is currently restricted to the `TELEGRAM_ADMIN_ID`, falling back to natural language processing for other users.
- **Concurrency**: `api_client.go` uses `sync.Mutex` to manage global configuration state, ensuring thread safety during configuration loads.

### Recommendations
1. **Cleanup Secrets**: Verify if `secrets.yaml` is needed. If not, remove its generation from `deployment.yml` to avoid confusion.
2. **Standardize Search Output**: Align `GetBibleSearch` output with the intended design (HTML anchor tags) if that is still the requirement.
3. **Test Coverage**: The integration tests (`database_integration_test.go`) and mock utilities (`test_utils_mock.go`) provide a good foundation. Continue expanding unit tests for edge cases in `natural_language.go`.
