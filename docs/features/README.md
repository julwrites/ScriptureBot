# Features Documentation

## Core Features

### 1. Bible Passage Retrieval
- **Command**: Direct reference (e.g., `John 3:16`) or `/passage`.
- **Functionality**: Retrieves and displays Bible verses.
- **Implementation**:
    - Primary: Bible AI API (`SubmitQuery`).
    - Fallback: Scraping `classic.biblegateway.com`.

### 2. Bible Search
- **Command**: `/search <phrase>` or short phrases (< 5 words).
- **Functionality**: Finds verses containing the search terms.
- **Implementation**: Uses Bible AI API `query.words` endpoint.

### 3. Bible Ask (AI Q&A)
- **Command**: `/ask <question>`.
- **Restriction**: Currently limited to Administrators.
- **Functionality**: Uses an LLM to answer questions about the Bible, optionally using context verses.
- **Implementation**: Uses Bible AI API `query.prompt` endpoint.

### 4. Devotionals
- **Command**: `/devo`.
- **Functionality**: Provides daily devotional readings from various sources (Daily NT, DJBR, N5BR).
- **Data**: Sources defined in `resource/*.yaml`.

### 5. Topical Memory System (TMS)
- **Command**: `/tms`.
- **Functionality**: retrieves verses for memorization based on topics or IDs.
- **Data**: Defined in `resource/tms_data.yaml`.

### 6. User Management
- **Subscriptions**: Users can subscribe to daily updates.
- **Preferences**: Stores translation version preference (default: NET).
