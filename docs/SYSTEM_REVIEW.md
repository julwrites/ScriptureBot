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
  - Used for persisting user data and configuration. Integration is handled via `BotPlatform` types and local utilities.
- **Secrets Management**:
  - Uses **Google Secret Manager** for production secrets (`TELEGRAM_ID`, `TELEGRAM_ADMIN_ID`, etc.).
  - Supports local environment variables (`.env`) for development.
  - *Note*: The `deployment.yml` generates a `secrets.yaml` file which appears to be unused by the application logic.

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

### Bible AI API Integration
- **Status**: **Healthy & Loose Coupling**.
- **Client**: The integration is encapsulated within `pkg/app/api_client.go`.
- **Abstraction**: The application defines its own Request/Response structs (`QueryRequest`, `VerseResponse`) in `pkg/app/api_models.go`, which mirror the API contract.
- **Resilience**: The system implements a fallback to web scraping (`GetBiblePassageFallback`) for passage retrieval, ensuring basic functionality remains even if the API is down.
- **Configuration**: Uses a thread-safe singleton pattern for API credentials.

### Infrastructure Integration
- **Google Cloud**:
  - **Cloud Run**: Hosting.
  - **Secret Manager**: Configuration.
  - **Artifact Registry**: Docker image storage.
- **CI/CD**:
  - GitHub Actions pipelines for `Build & Test` (`automation.yml`) and `Build, Stage & Deploy` (`deployment.yml`).
  - Deployment automatically triggers a webhook update via `cmd/webhook`.

---

## 3. Deep Dive: BotPlatform Coupling

### Overview
While the BibleAIAPI integration is well-encapsulated, the integration with `BotPlatform` exhibits **significant tight coupling**. The application logic (`pkg/app`) and data persistence layer (`pkg/utils`) rely directly on the framework's internal data structures (`def.SessionData`, `def.UserData`) and interfaces (`platform.Platform`).

### Critical Coupling Points

#### 1. Domain Layer Contamination
The core business logic functions in `pkg/app` accept and return Framework types.
- **File**: `pkg/app/command.go`
- **Function**: `ProcessCommand(env def.SessionData, bot platform.Platform) def.SessionData`
- **Impact**: The "Domain" layer cannot be used without the "Framework" layer. You cannot easily extract the Bible search logic to a CLI tool or a different web server without importing the entire `BotPlatform` dependency.

#### 2. Persistence Layer Leakage
The database utility functions use Framework types as Data Transfer Objects (DTOs).
- **File**: `pkg/utils/database.go`
- **Function**: `GetUser(user def.UserData, project string) def.UserData`
- **Impact**: The database schema is effectively tied to the `BotPlatform`'s `UserData` struct definition. If `BotPlatform` modifies `UserData` (e.g., changing field tags or types), it could break database compatibility or require migration of the `ScriptureBot` database.

#### 3. Side Effects in Logic
Some logic functions bypass the return-value flow and interact directly with the bot platform to send messages.
- **File**: `pkg/app/devo.go`
- **Function**: `GetDevo`
- **Code**: `bot.Post(env)` is called to send an intermediate "Just a moment..." message.
- **Impact**: This violation of the Command-Query Separation principle makes the function impure and difficult to test. A unit test for `GetDevo` will fail unless the `bot` interface is mocked to handle the `Post` call.

#### 4. Presentation Logic in Application
The Application layer constructs platform-specific formatting.
- **File**: `pkg/app/passage.go`
- **Code**: Uses `platform.TelegramBold(...)`, `platform.TelegramItalics(...)`.
- **Impact**: The application logic "knows" it is rendering for Telegram. Supporting another platform (e.g., Discord, Slack) would require modifying the core `passage.go` logic to handle different formatting strategies.

### Recommendations (Refactoring Strategy)

To solve this, a **Ports and Adapters (Hexagonal)** architecture is recommended:

1.  **Define Domain Entities**: Create `app.User` and `app.Session` structs in `pkg/app` (or `pkg/domain`) that contain *only* the data needed for ScriptureBot's logic.
2.  **Create Adapters**:
    - In `pkg/bot`, write translation functions: `AdapterToDomain(def.SessionData) app.Session` and `AdapterToPlatform(app.Session) def.SessionData`.
3.  **Abstract Formatting**:
    - Define a `Formatter` interface in the domain: `type Formatter interface { Bold(string) string }`.
    - Implement a `TelegramFormatter` in `pkg/bot` and pass it to the application logic.
4.  **Isolate Persistence**:
    - Ensure `pkg/utils/database.go` maps `app.User` to the Firestore entity, rather than storing `def.UserData` directly.

---

## 4. Code Review

### Structure
- The code is well-organized into logical packages:
  - `pkg/app`: Core business logic (Passage, Search, Ask).
  - `pkg/bot`: Telegram handler and bot logic wiring.
  - `pkg/secrets`: Secrets retrieval strategy.
  - `pkg/utils`: Helper functions (HTML parsing, Database).

### Key Observations & Findings
- **Secrets Confusion**: `deployment.yml` generates a `secrets.yaml` file and copies it into the Docker image, but the application code (`pkg/secrets`) does not appear to read this file format. It relies on environment variables or Google Secret Manager. This `secrets.yaml` might be dead code.
- **Search Output Format**: The `GetBibleSearch` function in `pkg/app/search.go` formats results as a simple text list (`- Verse`). However, project documentation/memory suggests it should return HTML anchor tags. This indicates a potential regression or incomplete implementation.
- **Fuzzy Matching**: `pkg/app/bible_reference.go` implements a sophisticated Levenshtein distance algorithm for identifying Bible book names, handling misspellings and abbreviations effectively.
- **Access Control**: The `/ask` command (`GetBibleAsk`) is currently restricted to the `TELEGRAM_ADMIN_ID`, falling back to natural language processing for other users.
