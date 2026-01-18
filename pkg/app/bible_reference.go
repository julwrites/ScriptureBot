package app

import (
	"sort"
	"strings"
)

var BibleBooks = map[string]string{
	// Old Testament
	"genesis": "Genesis", "gen": "Genesis", "gn": "Genesis",
	"exodus": "Exodus", "ex": "Exodus", "exo": "Exodus", "exod": "Exodus",
	"leviticus": "Leviticus", "lev": "Leviticus", "le": "Leviticus", "lv": "Leviticus",
	"numbers": "Numbers", "num": "Numbers", "nu": "Numbers", "nm": "Numbers",
	"deuteronomy": "Deuteronomy", "deut": "Deuteronomy", "de": "Deuteronomy", "dt": "Deuteronomy",
	"joshua": "Joshua", "josh": "Joshua", "jos": "Joshua",
	"judges": "Judges", "judg": "Judges", "jdg": "Judges",
	"ruth": "Ruth", "ru": "Ruth", "rth": "Ruth",
	"1 samuel": "1 Samuel", "1 sam": "1 Samuel", "1 sa": "1 Samuel", "1sam": "1 Samuel", "1sa": "1 Samuel", "i sam": "1 Samuel", "i samuel": "1 Samuel", "1st samuel": "1 Samuel",
	"2 samuel": "2 Samuel", "2 sam": "2 Samuel", "2 sa": "2 Samuel", "2sam": "2 Samuel", "2sa": "2 Samuel", "ii sam": "2 Samuel", "ii samuel": "2 Samuel", "2nd samuel": "2 Samuel",
	"1 kings": "1 Kings", "1 kgs": "1 Kings", "1 ki": "1 Kings", "1kings": "1 Kings", "i kgs": "1 Kings", "i kings": "1 Kings", "1st kings": "1 Kings",
	"2 kings": "2 Kings", "2 kgs": "2 Kings", "2 ki": "2 Kings", "2kings": "2 Kings", "ii kgs": "2 Kings", "ii kings": "2 Kings", "2nd kings": "2 Kings",
	"1 chronicles": "1 Chronicles", "1 chron": "1 Chronicles", "1 chr": "1 Chronicles", "1 ch": "1 Chronicles", "1chronicles": "1 Chronicles", "i chron": "1 Chronicles", "i chronicles": "1 Chronicles", "1st chronicles": "1 Chronicles",
	"2 chronicles": "2 Chronicles", "2 chron": "2 Chronicles", "2 chr": "2 Chronicles", "2 ch": "2 Chronicles", "2chronicles": "2 Chronicles", "ii chron": "2 Chronicles", "ii chronicles": "2 Chronicles", "2nd chronicles": "2 Chronicles",
	"ezra": "Ezra", "ezr": "Ezra",
	"nehemiah": "Nehemiah", "neh": "Nehemiah", "ne": "Nehemiah",
	"esther": "Esther", "est": "Esther", "esth": "Esther",
	"job": "Job", "jb": "Job",
	"psalms": "Psalms", "psalm": "Psalms", "ps": "Psalms", "psa": "Psalms",
	"proverbs": "Proverbs", "prov": "Proverbs", "pro": "Proverbs", "pr": "Proverbs",
	"ecclesiastes": "Ecclesiastes", "eccl": "Ecclesiastes", "ecc": "Ecclesiastes",
	"song of solomon": "Song of Solomon", "song": "Song of Solomon", "songs": "Song of Solomon", "sos": "Song of Solomon", "song of songs": "Song of Solomon",
	"isaiah": "Isaiah", "isa": "Isaiah",
	"jeremiah": "Jeremiah", "jer": "Jeremiah", "je": "Jeremiah",
	"lamentations": "Lamentations", "lam": "Lamentations",
	"ezekiel": "Ezekiel", "ezek": "Ezekiel", "eze": "Ezekiel", "ezk": "Ezekiel",
	"daniel": "Daniel", "dan": "Daniel", "dn": "Daniel",
	"hosea": "Hosea", "hos": "Hosea",
	"joel": "Joel", "jl": "Joel",
	"amos": "Amos", "am": "Amos",
	"obadiah": "Obadiah", "obad": "Obadiah", "ob": "Obadiah",
	"jonah": "Jonah", "jon": "Jonah", "jnh": "Jonah",
	"micah": "Micah", "mic": "Micah",
	"nahum": "Nahum", "nah": "Nahum",
	"habakkuk": "Habakkuk", "hab": "Habakkuk",
	"zephaniah": "Zephaniah", "zeph": "Zephaniah", "zep": "Zephaniah",
	"haggai": "Haggai", "hag": "Haggai", "hg": "Haggai",
	"zechariah": "Zechariah", "zech": "Zechariah", "zec": "Zechariah",
	"malachi": "Malachi", "mal": "Malachi",

	// New Testament
	"matthew": "Matthew", "matt": "Matthew", "mat": "Matthew", "mt": "Matthew",
	"mark": "Mark", "mrk": "Mark", "mk": "Mark", "mr": "Mark",
	"luke": "Luke", "luk": "Luke", "lk": "Luke",
	"john": "John", "jn": "John", "jhn": "John", "joh": "John",
	"acts": "Acts", "ac": "Acts", "act": "Acts",
	"romans": "Romans", "rom": "Romans", "ro": "Romans", "rm": "Romans",
	"1 corinthians": "1 Corinthians", "1 cor": "1 Corinthians", "1 co": "1 Corinthians", "1cor": "1 Corinthians", "i cor": "1 Corinthians", "i corinthians": "1 Corinthians", "1st corinthians": "1 Corinthians",
	"2 corinthians": "2 Corinthians", "2 cor": "2 Corinthians", "2 co": "2 Corinthians", "2cor": "2 Corinthians", "ii cor": "2 Corinthians", "ii corinthians": "2 Corinthians", "2nd corinthians": "2 Corinthians",
	"galatians": "Galatians", "gal": "Galatians", "ga": "Galatians",
	"ephesians": "Ephesians", "eph": "Ephesians", "ep": "Ephesians",
	"philippians": "Philippians", "phil": "Philippians", "php": "Philippians",
	"colossians": "Colossians", "col": "Colossians",
	"1 thessalonians": "1 Thessalonians", "1 thess": "1 Thessalonians", "1 th": "1 Thessalonians", "1thess": "1 Thessalonians", "i thess": "1 Thessalonians", "i thessalonians": "1 Thessalonians", "1st thessalonians": "1 Thessalonians",
	"2 thessalonians": "2 Thessalonians", "2 thess": "2 Thessalonians", "2 th": "2 Thessalonians", "2thess": "2 Thessalonians", "ii thess": "2 Thessalonians", "ii thessalonians": "2 Thessalonians", "2nd thessalonians": "2 Thessalonians",
	"1 timothy": "1 Timothy", "1 tim": "1 Timothy", "1 ti": "1 Timothy", "1tim": "1 Timothy", "i tim": "1 Timothy", "i timothy": "1 Timothy", "1st timothy": "1 Timothy",
	"2 timothy": "2 Timothy", "2 tim": "2 Timothy", "2 ti": "2 Timothy", "2tim": "2 Timothy", "ii tim": "2 Timothy", "ii timothy": "2 Timothy", "2nd timothy": "2 Timothy",
	"titus": "Titus", "tit": "Titus", "ti": "Titus",
	"philemon": "Philemon", "philem": "Philemon", "phlm": "Philemon", "phm": "Philemon",
	"hebrews": "Hebrews", "heb": "Hebrews",
	"james": "James", "jas": "James", "jm": "James",
	"1 peter": "1 Peter", "1 pet": "1 Peter", "1 pe": "1 Peter", "1 pt": "1 Peter", "1peter": "1 Peter", "i pet": "1 Peter", "i peter": "1 Peter", "1st peter": "1 Peter",
	"2 peter": "2 Peter", "2 pet": "2 Peter", "2 pe": "2 Peter", "2 pt": "2 Peter", "2peter": "2 Peter", "ii pet": "2 Peter", "ii peter": "2 Peter", "2nd peter": "2 Peter",
	"1 john": "1 John", "1 jn": "1 John", "1jn": "1 John", "1john": "1 John", "i jn": "1 John", "i john": "1 John", "1st john": "1 John",
	"2 john": "2 John", "2 jn": "2 John", "2jn": "2 John", "2john": "2 John", "ii jn": "2 John", "ii john": "2 John", "2nd john": "2 John",
	"3 john": "3 John", "3 jn": "3 John", "3jn": "3 John", "3john": "3 John", "iii jn": "3 John", "iii john": "3 John", "3rd john": "3 John",
	"jude": "Jude", "jud": "Jude", "jd": "Jude",
	"revelation": "Revelation", "rev": "Revelation",
}

var SingleChapterBooks = map[string]bool{
	"Obadiah":  true,
	"Philemon": true,
	"2 John":   true,
	"3 John":   true,
	"Jude":     true,
}

// sortedBookKeys caches the sorted keys of BibleBooks for efficient matching
var sortedBookKeys []string

// uniqueCanonicalBooks stores unique canonical book names for fuzzy matching
var uniqueCanonicalBooks []string

func init() {
	// Initialize sorted keys sorted by length (descending) to ensure greedy matching
	// e.g. "1 John" matches before "John"
	uniqueMap := make(map[string]bool)
	for k, v := range BibleBooks {
		sortedBookKeys = append(sortedBookKeys, k)
		if !uniqueMap[v] {
			uniqueCanonicalBooks = append(uniqueCanonicalBooks, v)
			uniqueMap[v] = true
		}
	}
	sort.Slice(sortedBookKeys, func(i, j int) bool {
		return len(sortedBookKeys[i]) > len(sortedBookKeys[j])
	})
}

// ParseBibleReference parses a string to identify and normalize a Bible reference.
// It returns the normalized reference string and a boolean indicating validity.
func ParseBibleReference(input string) (string, bool) {
	ref, consumedLen, ok := ParseBibleReferenceFromStart(input)
	if !ok {
		return "", false
	}
	// Verify that we consumed the entire string (ignoring whitespace)
	if len(strings.TrimSpace(input[consumedLen:])) > 0 {
		return "", false
	}
	return ref, true
}

func findFuzzyMatch(input string) (string, int) {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", 0
	}

	// Try checking the first 1, 2, or 3 words as potential book names
	// e.g. "1 John", "Song of Solomon"
	maxWords := 3
	if len(tokens) < 3 {
		maxWords = len(tokens)
	}

	bestBook := ""
	bestMatchLen := 0
	minDist := 100

	// Check longest candidate first? Or shortest?
	// If input is "Gensis 1", tokens=["Gensis", "1"].
	// 2 words: "Gensis 1". Dist to "Genesis" (7)? "Gensis 1" (8). Dist is large.
	// 1 word: "Gensis". Dist to "Genesis". Dist 1. Match!
	// We should probably prioritize the match with the smallest distance relative to length?
	// Or just smallest distance.

	for i := maxWords; i >= 1; i-- {
		candidate := strings.Join(tokens[:i], " ")
		lowerCand := strings.ToLower(candidate)

		for _, canonical := range uniqueCanonicalBooks {
			lowerCanon := strings.ToLower(canonical)

			// Threshold Logic
			threshold := 0
			canonLen := len(lowerCanon)
			if canonLen >= 7 {
				threshold = 2
			} else if canonLen >= 5 {
				threshold = 1
			}
			// Length < 5: threshold 0 (Exact match only).
			// Since exact match is handled before this function, we can effectively skip if threshold is 0.
			if threshold == 0 {
				continue
			}

			dist := levenshteinDistance(lowerCand, lowerCanon)

			if dist <= threshold {
				// Found a match. Since we want "best" match, we should continue?
				// Or since we iterate decreasing word count, finding a match on more words is better?
				// But "Gensis 1" -> 2 words "Gensis 1". No match.
				// 1 word "Gensis". Match.
				// If we have "Song of Solmon", 3 words. Match.
				// So we should return immediately if we find a valid match?
				// But what if multiple books match?
				// "Mathew" -> "Matthew" (dist 1).
				// We can track the best match (lowest distance).

				if dist < minDist {
					minDist = dist
					bestBook = canonical
					bestMatchLen = len(candidate)
				}
			}
		}
	}

	if bestBook != "" {
		return bestBook, bestMatchLen
	}
	return "", 0
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// ExtractBibleReferences extracts all valid Bible references from a text.
func ExtractBibleReferences(input string) []string {
	var refs []string

	startIdx := 0
	length := len(input)

	for startIdx < length {
		// If we are at a whitespace, advance.
		if input[startIdx] == ' ' || input[startIdx] == '\t' || input[startIdx] == '\n' || input[startIdx] == '\r' {
			startIdx++
			continue
		}

		ref, consumed, ok := ParseBibleReferenceFromStart(input[startIdx:])
		if ok {
			refs = append(refs, ref)
			startIdx += consumed
		} else {
			// Advance to next word
			nextSpace := strings.IndexAny(input[startIdx:], " \t\n\r")
			if nextSpace == -1 {
				break
			}
			startIdx += nextSpace
		}
	}
	return refs
}

// ParseBibleReferenceFromStart attempts to parse a Bible reference at the beginning of the string.
// Returns the normalized reference, the length of text consumed from input, and whether a match was found.
func ParseBibleReferenceFromStart(input string) (string, int, bool) {
	// 1. Skip leading whitespace
	startOffset := 0
	for startOffset < len(input) && (input[startOffset] == ' ' || input[startOffset] == '\t' || input[startOffset] == '\n' || input[startOffset] == '\r') {
		startOffset++
	}
	if startOffset == len(input) {
		return "", 0, false
	}

	currentInput := input[startOffset:]
	lowerInput := strings.ToLower(currentInput)

	var foundBook string
	var bookName string
	var matchLen int // Length in currentInput

	// 1. Try exact match (Greedy)
	for _, key := range sortedBookKeys {
		if strings.HasPrefix(lowerInput, key) {
			mLen := len(key)
			// Ensure whole word match
			if len(lowerInput) > mLen {
				nextChar := lowerInput[mLen]
				if isLetter(nextChar) {
					continue
				}
			}

			foundBook = key
			bookName = BibleBooks[key]
			matchLen = mLen
			break
		}
	}

	// 2. If no exact match, try fuzzy matching
	if foundBook == "" {
		fBook, mLen := findFuzzyMatch(currentInput)
		if fBook != "" {
			bookName = fBook
			foundBook = fBook
			matchLen = mLen
		}
	}

	if foundBook == "" {
		return "", 0, false
	}

	// We found a book. Now parse the numbers (remainder).
	remainderStart := matchLen
	// Skip spaces after book
	for remainderStart < len(currentInput) && (currentInput[remainderStart] == ' ' || currentInput[remainderStart] == '\t') {
		remainderStart++
	}

	remainder := currentInput[remainderStart:]

	// Consume valid reference syntax
	syntax, syntaxLen := consumeReferenceSyntax(remainder)

	totalConsumed := startOffset + remainderStart + syntaxLen

	if syntax == "" {
		if SingleChapterBooks[bookName] {
			return bookName, totalConsumed, true
		}
		// Multi-chapter book defaults to chapter 1
		return bookName + " 1", totalConsumed, true
	}

	if hasDigit(syntax) {
		return bookName + " " + syntax, totalConsumed, true
	}

	if SingleChapterBooks[bookName] {
		return bookName, startOffset + matchLen, true // Don't consume the invalid syntax
	}
	return bookName + " 1", startOffset + matchLen, true
}

func consumeReferenceSyntax(s string) (string, int) {
	lastDigit := -1

	for i, r := range s {
		if r >= '0' && r <= '9' {
			lastDigit = i
		} else if r == ':' || r == '-' || r == '.' || r == ',' || r == ' ' || r == '\t' {
			continue
		} else {
			break
		}
	}

	if lastDigit == -1 {
		return "", 0
	}

	return s[:lastDigit+1], lastDigit + 1
}

func hasDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func levenshteinDistance(s1, s2 string) int {
	r1, r2 := []rune(s1), []rune(s2)
	n, m := len(r1), len(r2)

	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}

	// Use two rows instead of full matrix
	row := make([]int, m+1)
	for j := 0; j <= m; j++ {
		row[j] = j
	}

	for i := 1; i <= n; i++ {
		prevRow := make([]int, m+1)
		copy(prevRow, row)
		row[0] = i
		for j := 1; j <= m; j++ {
			cost := 0
			if r1[i-1] != r2[j-1] {
				cost = 1
			}
			row[j] = min(
				prevRow[j]+1,      // deletion
				row[j-1]+1,        // insertion
				prevRow[j-1]+cost, // substitution
			)
		}
	}
	return row[m]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
