# AI Features and Security Review

## Summary
This document summarizes the findings from a review of the ScriptureBot's AI features and security implementation.

## AI Features

### 1. Limited Access
Currently, AI features (`/ask` command and context-aware queries) are strictly gated to the `TELEGRAM_ADMIN_ID`. Regular users cannot access these capabilities.

### 2. Statelessness
The AI interaction is largely stateless. While the `QueryContext` structure supports a `History` field, it is currently not populated with conversation history, meaning the AI treats each request in isolation.

### 3. Basic Prompting
The prompt sent to the AI is simply the user's raw message. There is no "System Prompt" or instruction tuning to guide the AI's persona, tone, or formatting, relying entirely on the model's default behavior.

### 4. Natural Language Fallback
The natural language processing pipeline (`ProcessNaturalLanguage`) has a "catch-all" case that currently does nothing for unrecognized inputs. This leaves users without feedback if their query doesn't match a specific pattern (like a Bible reference or search command).

## Security Implementation

### 1. Inefficient Secret Management
The `pkg/secrets` package loads secrets (including making calls to Google Secret Manager) on every request in `main.go`. This is highly inefficient and could lead to API rate limiting and increased costs.

### 2. Hardcoded Admin Checks
The logic to check if a user is an admin (`env.User.Id == adminID`) is duplicated across multiple files (`pkg/app/admin.go`, `pkg/app/ask.go`, `pkg/app/natural_language.go`). This makes it difficult to maintain or modify admin permissions (e.g., adding multiple admins).

### 3. HTML Sanitization
The bot uses `ParseToTelegramHTML` to sanitize HTML output from the AI. This implementation appears robust, correctly escaping text content while allowing a safe subset of tags (`<b>`, `<i>`, etc.).

### 4. Error Handling
In some cases, raw API error messages might be returned to the user. While sensitive headers like `X-API-KEY` are handled separately, exposing internal error details is generally discouraged.

## Recommendations

1.  **Refactor Secret Management:** Implement caching for secrets to prevent redundant external calls.
2.  **Centralize Admin Logic:** Create a `utils.IsAdmin(env)` helper to unify permission checks.
3.  **Enhance Natural Language Processing:** Provide a meaningful fallback response for unrecognized queries, and consider enabling AI features for all users (potentially with rate limits).
4.  **Improve AI Context:** populate the conversation history in `QueryContext` to enable multi-turn conversations.
5.  **Implement System Prompts:** Add a system prompt to guide the AI's behavior and ensure consistent, high-quality responses.
