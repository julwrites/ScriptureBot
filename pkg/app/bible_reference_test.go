package app

import (
	"testing"
)

func TestParseBibleReference(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		valid    bool
	}{
		// Valid References (Exact)
		{"gen 1", "Genesis 1", true},
		{"Genesis 1", "Genesis 1", true},
		{"mk 2", "Mark 2", true},
		{"mat 5:16", "Matthew 5:16", true},
		{"heb 13:3-4", "Hebrews 13:3-4", true},
		{"Exodus 2-3", "Exodus 2-3", true},
		{"1jn 3:16", "1 John 3:16", true},
		{"Jude", "Jude", true},
		{"jude", "Jude", true},
		{"3 John", "3 John", true},
		{"3 jn", "3 John", true},
		{"Genesis", "Genesis 1", true}, // Multi-chapter default
		{"Obadiah", "Obadiah", true},
		{"Phil 1", "Philippians 1", true},
		{"Phlm 1", "Philemon 1", true},

		// Fuzzy Matches (Typos)
		{"Gensis 1", "Genesis 1", true},         // Missing 'e', dist 1
		{"Genisis 1", "Genesis 1", true},        // 'i' instead of 'e', dist 1
		{"Mathew 5", "Matthew 5", true},         // Missing 't', dist 1
		{"Revalation 3", "Revelation 3", true},  // 'a' instead of 'e', dist 1
		{"Philipians 4", "Philippians 4", true}, // Missing 'p', dist 1
		{"1 Jhn 3", "1 John 3", true},           // Missing 'o', dist 1. "1 Jhn" vs "1 John".

		// Thresholds / False Positives
		{"Genius 1", "", false},    // Dist to Genesis is > threshold? "Genius" (6) vs "Genesis" (7). Dist 3. Threshold 1. False.
		{"Mary 1", "", false},      // "Mary" (4). Threshold 0. "Mark" (4). Dist 1. No fuzzy allowed for len < 5.
		{"Mark 1", "Mark 1", true}, // Exact match.
		{"Luke 1", "Luke 1", true}, // Exact match.
		{"Luke", "Luke 1", true},   // Exact match.
		{"Luek 1", "", false},      // "Luek" (4). Threshold 0. No match.

		// Invalid References
		{"John is here", "", false},
		{"Exiting in 2", "", false},
		{"I am thinking of...", "", false},
		{"General", "", false},
		{"Hello World", "", false},
		{"", "", false},
		{"1", "", false},
		{"Genesis is great", "", false},
		{"Jude is short", "", false},
		{"NotABook 1", "", false},
	}

	for _, tt := range tests {
		result, valid := ParseBibleReference(tt.input)
		if valid != tt.valid {
			t.Errorf("ParseBibleReference(%q) valid = %v, want %v", tt.input, valid, tt.valid)
		}
		if valid && result != tt.expected {
			t.Errorf("ParseBibleReference(%q) result = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestExtractBibleReferences(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"Read John 3:16", []string{"John 3:16"}},
		{"Compare Gen 1:1 and Ex 20", []string{"Genesis 1:1", "Exodus 20"}},
		{"What does it say in Mark 5?", []string{"Mark 5"}},
		{"I like Genesis", []string{"Genesis 1"}}, // Defaults to 1? Yes, per current logic.
		{"No references here", nil},
		{"John said hello", []string{"John 1"}}, // False positive risk, but per logic.
		{"Read 1 John 3 and 2 John", []string{"1 John 3", "2 John"}},
	}

	for _, tt := range tests {
		result := ExtractBibleReferences(tt.input)
		if len(result) != len(tt.expected) {
			t.Errorf("ExtractBibleReferences(%q) length = %d, want %d", tt.input, len(result), len(tt.expected))
			continue
		}
		for i, ref := range result {
			if ref != tt.expected[i] {
				t.Errorf("ExtractBibleReferences(%q)[%d] = %q, want %q", tt.input, i, ref, tt.expected[i])
			}
		}
	}
}
