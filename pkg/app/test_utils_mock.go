package app

// MockSubmitQuery is a helper to mock SubmitQuery for testing purposes.
// It captures the request in a closure and allows verifying it.
func MockSubmitQuery(t HelperT, callback func(QueryRequest)) func(QueryRequest, interface{}) error {
	return func(req QueryRequest, result interface{}) error {
		callback(req)

		// Return dummy success data to prevent nil pointer dereferences in handlers
		switch r := result.(type) {
		case *WordSearchResponse:
			*r = WordSearchResponse{
				{Verse: "John 3:16", URL: "https://example.com/John3:16"},
			}
		case *OQueryResponse:
			*r = OQueryResponse{
				Text: "This is a mock response.",
				References: []SearchResult{
					{Verse: "John 3:16", URL: "https://example.com/John3:16"},
				},
			}
		case *VerseResponse:
			*r = VerseResponse{
				Verse: "For God so loved the world...",
			}
		}
		return nil
	}
}

// HelperT is an interface to allow passing *testing.T
type HelperT interface {
	Helper()
	Errorf(format string, args ...interface{})
}
