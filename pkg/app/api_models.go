package app

// Request Models

type UserOptions struct {
	Version    string `json:"version,omitempty"`
	AIProvider string `json:"ai_provider,omitempty"`
}

type Options struct {
	Stream bool `json:"stream,omitempty"`
}

type QueryContext struct {
	History []string `json:"history,omitempty"`
	Schema  string   `json:"schema,omitempty"`
	Verses  []string `json:"verses,omitempty"`
	Words   []string `json:"words,omitempty"`
}

type QueryObject struct {
	Verses  []string     `json:"verses,omitempty"`
	Words   []string     `json:"words,omitempty"`
	Prompt  string       `json:"prompt,omitempty"`
	Context QueryContext `json:"context,omitempty"`
}

type QueryRequest struct {
	Query   QueryObject `json:"query"`
	User    UserOptions `json:"user,omitempty"`
	Options Options     `json:"options,omitempty"`
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

type Meta struct {
	AIProvider string `json:"ai_provider"`
}

type PromptResponse struct {
	Data OQueryResponse `json:"data"`
	Meta Meta           `json:"meta"`
}

type ErrorResponse struct {
	Error ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Version struct {
	Name      string            `json:"name"`
	Code      string            `json:"code"`
	Language  string            `json:"language"`
	Providers map[string]string `json:"providers"`
}

type VersionsResponse struct {
	Data  []Version `json:"data"`
	Total int       `json:"total"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}
