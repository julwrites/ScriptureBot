# Architecture Documentation

## Overview
ScriptureBot is a Go-based Telegram bot built on the [BotPlatform](https://github.com/julwrites/BotPlatform) framework. It follows a layered architecture to separate concerns between transport, translation, and business logic.

## High-Level Architecture
1.  **Web App (Entry Point)**: `main.go` listens for incoming webhooks from Telegram.
2.  **Translation Layer**: `BotPlatform` handles parsing incoming JSON updates and formatting outgoing responses.
3.  **Logic Layer (`pkg/app`)**: Contains the core business logic for Bible passage retrieval, search, Q&A, and devotionals.
4.  **Data Layer**:
    - **Firestore (Datastore mode)**: Stores user preferences and subscription data.
    - **Secret Manager**: Securely stores API keys and credentials.
    - **External APIs**:
        - **Bible AI API**: Provides scripture text, search results, and LLM-based answers.
        - **BibleGateway (Legacy)**: Fallback scraping for passages.

## Directory Structure
- `cmd/`: Command-line tools (e.g., `migrate`, `webhook`).
- `pkg/app/`: Application logic (Passage, Search, Ask, Devo).
- `pkg/bot/`: Bot interface implementation.
- `pkg/secrets/`: Secret management logic.
- `pkg/utils/`: Shared utilities (HTML parsing, database helpers).
- `resource/`: YAML data files for devotionals and TMS.

## Infrastructure
- **Hosting**: Google Cloud Run (Containerized).
- **CI/CD**: GitHub Actions (`automation.yml` for tests, `deployment.yml` for deploy).
- **Container**: Docker (Multi-stage build).
- **Region**: `asia-southeast1`.
