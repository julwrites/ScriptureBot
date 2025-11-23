package app

// Request Models

type UserContext struct {
	Version string `json:"version,omitempty"`
}

type QueryContext struct {
	History []string    `json:"history,omitempty"`
	Schema  string      `json:"schema,omitempty"`
	Verses  []string    `json:"verses,omitempty"`
	Words   []string    `json:"words,omitempty"`
	User    UserContext `json:"user,omitempty"`
}

type QueryObject struct {
	Verses []string `json:"verses,omitempty"`
	Words  []string `json:"words,omitempty"`
	Prompt string   `json:"prompt,omitempty"`
}

type QueryRequest struct {
	Query   QueryObject  `json:"query"`
	Context QueryContext `json:"context,omitempty"`
}

// Response Models

type VerseResponse struct {
	Verse string `json:"verse"`
}

type SearchResult struct {
	Verse string `json:"verse"`
	URL   string `json:"url"`
}

// WordSearchResponse is a list of SearchResults
type WordSearchResponse []SearchResult

type OQueryResponse struct {
	Text       string         `json:"text"`
	References []SearchResult `json:"references"`
}

type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
