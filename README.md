# Scripture Bot

![status: active](https://img.shields.io/badge/status-active-green.svg)

A Telegram bot to make the Bible more accessible, providing passages, search, and Q&A.

## Features
*   **Bible Passage**: Get any Bible passage (e.g., "John 3:16").
*   **Bible Search**: Search for words (`/search grace`).
*   **Bible Ask**: AI-powered Q&A (`/ask Who is Moses?`) [Admin Only].
*   **Devotionals**: Daily reading material (`/devo`).
*   **TMS**: Topical Memory System verses (`/tms`).

## Project Status
**Current Phase**: Transition & Migration
- Moving from legacy web scraping to a modern Bible AI API.
- Migrating infrastructure to Google Cloud `asia-southeast1`.

## Local Development

### Prerequisites
- Go 1.24+
- Docker (optional)

### Setup
1.  Clone the repository.
2.  Create a `.env` file in the root directory:
    ```env
    TELEGRAM_ID=your_bot_token
    TELEGRAM_ADMIN_ID=your_user_id
    BIBLE_API_URL=https://api.example.com (optional)
    BIBLE_API_KEY=your_key (optional)
    ```
3.  Run the bot:
    ```bash
    go run main.go
    ```

### Testing
- To receive Telegram updates locally, use a tool like `ngrok` to expose port 8080 and set your bot's webhook.
- Or send a mock HTTP POST request to `http://localhost:8080/<TELEGRAM_ID>`.

## Contributing
See [AGENTS.md](AGENTS.md) for architecture details and development guidelines.

## License
See [LICENSE](LICENSE)
