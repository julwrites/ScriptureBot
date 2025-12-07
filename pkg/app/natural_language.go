package app

import (
	"strings"

	"github.com/julwrites/BotPlatform/pkg/def"
)

func ProcessNaturalLanguage(env def.SessionData) def.SessionData {
	msg := env.Msg.Message

	// 1. Check if it is a Bible Reference (Only a verse)
	// ParseBibleReference checks for exact match of reference syntax
	if _, ok := ParseBibleReference(msg); ok {
		return GetBiblePassage(env)
	}

	// 2. Check if it contains references
	// If it contains references, we assume it's a query about them, so we Ask.
	refs := ExtractBibleReferences(msg)
	if len(refs) > 0 {
		return GetBibleAskWithContext(env, refs)
	}

	// 3. Check for "short phrase" (Search)
	// Definition: < 5 words and no question mark?
	words := strings.Fields(msg)
	if len(words) < 5 && !strings.Contains(msg, "?") {
		return GetBibleSearch(env)
	}

	// 4. Assume Query Prompt (Ask)
	return env
}
