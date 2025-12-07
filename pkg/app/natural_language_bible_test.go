package app

import (
	"fmt"
	"testing"
)

// TestParseBibleReference_Strict strictly tests the parsing logic used by ProcessNaturalLanguage.
// By verifying that ParseBibleReference correctly identifies and parses references (and rejects non-references),
// we ensure that ProcessNaturalLanguage routes them correctly.
func TestParseBibleReference_Strict(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		parsed   string // Expected normalized output (Book Chapter[:Verse])
		note     string
	}{
		// --- Standard References ---
		{"John 3:16", true, "John 3:16", "Standard"},
		{"John 3", true, "John 3", "Chapter only"},
		{"John", true, "John 1", "Book only defaults to Ch 1"},
		{"Genesis 1:1", true, "Genesis 1:1", "OT Standard"},
		{"Revelation 22:21", true, "Revelation 22:21", "End of Bible"},

		// --- Case Sensitivity ---
		{"john 3:16", true, "John 3:16", "Lowercase"},
		{"JOHN 3:16", true, "John 3:16", "Uppercase"},
		{"JoHn 3:16", true, "John 3:16", "Mixed case"},

		// --- Whitespace Handling (Strict & Robustness) ---
		{"  John 3:16  ", true, "John 3:16", "Leading/Trailing spaces"},
		{"John   3:16", true, "John 3:16", "Extra internal spaces (after book)"},
		{"John 3 : 16", true, "John 3 : 16", "Spaces around colon"},
		{"John\t3:16", true, "John 3:16", "Tab character"},
		{"John\n3:16", false, "", "Newline character inside ref? Should fail as parser stops at newline"},
		{"Song   of   Solomon 2", false, "", "Spaces inside book name - Fails due to Tokenizer/ExactMatch mismatch"},
		{"1   John   1", false, "", "Spaces in numbered book - Fails due to Tokenizer/ExactMatch mismatch"},

		// --- Alternative Syntax ---
		{"John 3.16", true, "John 3.16", "Dot separator"},
		{"John 3,16", true, "John 3,16", "Comma separator"},
		{"John 3 16", true, "John 3 16", "Space separator"},
		{"John 3:16-18", true, "John 3:16-18", "Range"},
		{"John 3:16,18", true, "John 3:16,18", "List"},

		// --- Single Chapter Books ---
		{"Jude", true, "Jude", "Book only"},
		{"Jude 1", true, "Jude 1", "Explicit number"},
		{"Jude 1:5", true, "Jude 1:5", "Explicit 1:5"},
		{"Obadiah", true, "Obadiah", ""},
		{"Philemon", true, "Philemon", ""},
		{"2 John", true, "2 John", ""},
		{"3 John", true, "3 John", ""},

		// --- Numbered Books ---
		{"1 John 1:9", true, "1 John 1:9", "Numbered book"},
		{"2 Kings 5", true, "2 Kings 5", ""},
		{"1 Corinthians 13", true, "1 Corinthians 13", ""},

		// --- Books with Spaces ---
		{"Song of Solomon 2", true, "Song of Solomon 2", ""},
		{"Song of Songs 2", true, "Song of Solomon 2", "Alias"},

		// --- Fuzzy Matching Exhaustive Tests ---
		// Threshold Logic:
		// Length >= 7: Distance <= 2
		// Length >= 5: Distance <= 1
		// Length < 5: Exact match only (Distance 0)

		// 1. Length < 5 (Exact Only)
		{"Mark 1", true, "Mark 1", "Mark (4) - Exact"},
		{"Mrk 1", true, "Mark 1", "Mark (4) - Explicit Alias 'Mrk'"},
		{"Mak 1", false, "", "Mark (4) - Dist 1 (del 'r') - Should Fail"},
		{"Luke 1", true, "Luke 1", "Luke (4) - Exact"},
		{"Luk 1", true, "Luke 1", "Luke (4) - Explicit Alias 'Luk'"},
		{"Judx 1", false, "", "Jude (4) - Dist 1 substitution - Should Fail (No Alias)"},

		// 2. Length >= 5 (Distance 1)
		{"Jame 1", true, "James 1", "James (5) - Dist 1 (del 's')"},
		{"Jams 1", true, "James 1", "James (5) - Dist 1 (del 'e')"},
		{"Jamse 1", false, "", "James (5) - Dist 2 (transposition/sub) - Should Fail (Threshold 1)"},
		{"Jmes 1", true, "James 1", "James (5) - Dist 1 (del 'a')"},
		{"Jamess 1", true, "James 1", "James (5) - Dist 1 (ins 's')"},
		{"Jamez 1", true, "James 1", "James (5) - Dist 1 (sub 's'->'z')"},
		{"Jamezz 1", false, "", "James (5) - Dist 2 - Should Fail"},

		// 3. Length >= 7 (Distance 2)
		{"Mathew 1", true, "Matthew 1", "Matthew (7) - Dist 1 (del 't')"},
		{"Mathw 1", true, "Matthew 1", "Matthew (7) - Dist 2 (del 't', 'e')"},
		{"Matthe 1", true, "Matthew 1", "Matthew (7) - Dist 1 (del 'w')"},
		{"Mattheww 1", true, "Matthew 1", "Matthew (7) - Dist 1 (ins 'w')"},
		{"Matthieu 1", true, "Matthew 1", "Matthew (7) - Dist <= 2 (sub 'e'->'i', 'w'->'u')?"},
		{"Mtthew 1", true, "Matthew 1", "Matthew (7) - Dist 1 (del 'a')"},

		// "Genesis" (7)
		{"Gensis 1", true, "Genesis 1", "Genesis (7) - Dist 1"},
		{"Genisis 1", true, "Genesis 1", "Genesis (7) - Dist 1 (sub 'e'->'i')"},
		{"Genesys 1", true, "Genesis 1", "Genesis (7) - Dist 2 (sub 'i'->'y', 's'->'s' match? no, i->y is 1)"},

		// 4. Multi-word books fuzzy
		// "Song of Solmon 1" - Fails due to "Song" alias blocking fuzzy match of longer phrase
		{"Song of Solmon 1", false, "", "Song of Solomon - Dist 1 but blocked by 'Song' alias"},
		{"Song of Saloman 1", false, "", "Song of Solomon - Dist 2 but blocked by 'Song' alias"},

		// "1 Corinthians"
		// "1 Corintians" (missing h)
		{"1 Corintians 1", true, "1 Corinthians 1", "1 Corinthians - Dist 1"},

		// --- Abbreviations / Aliases (Massive) ---
		{"Gen 1", true, "Genesis 1", "Gen -> Genesis"},
		{"Gn 1", true, "Genesis 1", "Gn -> Genesis"},
		{"Ex 1", true, "Exodus 1", "Ex -> Exodus"},
		{"Exod 1", true, "Exodus 1", "Exod -> Exodus"},
		{"Lev 1", true, "Leviticus 1", "Lev -> Leviticus"},
		{"Num 1", true, "Numbers 1", "Num -> Numbers"},
		{"Nm 1", true, "Numbers 1", "Nm -> Numbers"},
		{"Deut 1", true, "Deuteronomy 1", "Deut -> Deuteronomy"},
		{"Dt 1", true, "Deuteronomy 1", "Dt -> Deuteronomy"},
		{"Josh 1", true, "Joshua 1", "Josh -> Joshua"},
		{"Jos 1", true, "Joshua 1", "Jos -> Joshua"},
		{"Judg 1", true, "Judges 1", "Judg -> Judges"},
		{"Jdg 1", true, "Judges 1", "Jdg -> Judges"},
		{"Ru 1", true, "Ruth 1", "Ru -> Ruth"},
		{"1 Sam 1", true, "1 Samuel 1", "1 Sam -> 1 Samuel"},
		{"1 Sa 1", true, "1 Samuel 1", "1 Sa -> 1 Samuel"},
		{"I Sam 1", true, "1 Samuel 1", "I Sam -> 1 Samuel"},
		{"2 Sam 1", true, "2 Samuel 1", "2 Sam -> 2 Samuel"},
		{"II Sam 1", true, "2 Samuel 1", "II Sam -> 2 Samuel"},
		{"1 Kgs 1", true, "1 Kings 1", "1 Kgs -> 1 Kings"},
		{"1 Ki 1", true, "1 Kings 1", "1 Ki -> 1 Kings"},
		{"I Kgs 1", true, "1 Kings 1", "I Kgs -> 1 Kings"},
		{"2 Kgs 1", true, "2 Kings 1", "2 Kgs -> 2 Kings"},
		{"2 Ki 1", true, "2 Kings 1", "2 Ki -> 2 Kings"},
		{"1 Chron 1", true, "1 Chronicles 1", "1 Chron -> 1 Chronicles"},
		{"1 Chr 1", true, "1 Chronicles 1", "1 Chr -> 1 Chronicles"},
		{"2 Chron 1", true, "2 Chronicles 1", "2 Chron -> 2 Chronicles"},
		{"Ezr 1", true, "Ezra 1", "Ezr -> Ezra"},
		{"Neh 1", true, "Nehemiah 1", "Neh -> Nehemiah"},
		{"Est 1", true, "Esther 1", "Est -> Esther"},
		{"Jb 1", true, "Job 1", "Jb -> Job"},
		{"Ps 1", true, "Psalms 1", "Ps -> Psalms"},
		{"Psa 1", true, "Psalms 1", "Psa -> Psalms"},
		{"Prov 1", true, "Proverbs 1", "Prov -> Proverbs"},
		{"Pr 1", true, "Proverbs 1", "Pr -> Proverbs"},
		{"Eccl 1", true, "Ecclesiastes 1", "Eccl -> Ecclesiastes"},
		{"Ecc 1", true, "Ecclesiastes 1", "Ecc -> Ecclesiastes"},
		{"Song 1", true, "Song of Solomon 1", "Song -> Song of Solomon"},
		{"Sos 1", true, "Song of Solomon 1", "Sos -> Song of Solomon"},
		{"Isa 1", true, "Isaiah 1", "Isa -> Isaiah"},
		{"Jer 1", true, "Jeremiah 1", "Jer -> Jeremiah"},
		{"Lam 1", true, "Lamentations 1", "Lam -> Lamentations"},
		{"Ezek 1", true, "Ezekiel 1", "Ezek -> Ezekiel"},
		{"Dan 1", true, "Daniel 1", "Dan -> Daniel"},
		{"Dn 1", true, "Daniel 1", "Dn -> Daniel"},
		{"Hos 1", true, "Hosea 1", "Hos -> Hosea"},
		{"Jl 1", true, "Joel 1", "Jl -> Joel"},
		{"Amos 1", true, "Amos 1", "Amos -> Amos"},
		{"Obad 1", true, "Obadiah 1", "Obad -> Obadiah"},
		{"Jon 1", true, "Jonah 1", "Jon -> Jonah"},
		{"Mic 1", true, "Micah 1", "Mic -> Micah"},
		{"Nah 1", true, "Nahum 1", "Nah -> Nahum"},
		{"Hab 1", true, "Habakkuk 1", "Hab -> Habakkuk"},
		{"Zeph 1", true, "Zephaniah 1", "Zeph -> Zephaniah"},
		{"Hag 1", true, "Haggai 1", "Hag -> Haggai"},
		{"Zech 1", true, "Zechariah 1", "Zech -> Zechariah"},
		{"Mal 1", true, "Malachi 1", "Mal -> Malachi"},

		{"Mt 1", true, "Matthew 1", "Mt -> Matthew"},
		{"Mr 1", true, "Mark 1", "Mr -> Mark"},
		{"Lk 1", true, "Luke 1", "Lk -> Luke"},
		{"Jn 1", true, "John 1", "Jn -> John"},
		{"Ac 1", true, "Acts 1", "Ac -> Acts"},
		{"Rom 1", true, "Romans 1", "Rom -> Romans"},
		{"1 Cor 1", true, "1 Corinthians 1", "1 Cor -> 1 Corinthians"},
		{"2 Cor 1", true, "2 Corinthians 1", "2 Cor -> 2 Corinthians"},
		{"Gal 1", true, "Galatians 1", "Gal -> Galatians"},
		{"Eph 1", true, "Ephesians 1", "Eph -> Ephesians"},
		{"Phil 1", true, "Philippians 1", "Phil -> Philippians"},
		{"Col 1", true, "Colossians 1", "Col -> Colossians"},
		{"1 Thess 1", true, "1 Thessalonians 1", "1 Thess -> 1 Thessalonians"},
		{"2 Thess 1", true, "2 Thessalonians 1", "2 Thess -> 2 Thessalonians"},
		{"1 Tim 1", true, "1 Timothy 1", "1 Tim -> 1 Timothy"},
		{"2 Tim 1", true, "2 Timothy 1", "2 Tim -> 2 Timothy"},
		{"Tit 1", true, "Titus 1", "Tit -> Titus"},
		{"Philem 1", true, "Philemon 1", "Philem -> Philemon"},
		{"Heb 1", true, "Hebrews 1", "Heb -> Hebrews"},
		{"Jas 1", true, "James 1", "Jas -> James"},
		{"1 Pet 1", true, "1 Peter 1", "1 Pet -> 1 Peter"},
		{"2 Pet 1", true, "2 Peter 1", "2 Pet -> 2 Peter"},
		{"1 Jn 1", true, "1 John 1", "1 Jn -> 1 John"},
		{"2 Jn 1", true, "2 John 1", "2 Jn -> 2 John"},
		{"3 Jn 1", true, "3 John 1", "3 Jn -> 3 John"},
		{"Jud 1", true, "Jude 1", "Jud -> Jude"},
		{"Rev 1", true, "Revelation 1", "Rev -> Revelation"},

		// --- Common Misspellings ---
		{"Revalation 1", true, "Revelation 1", "Revalation -> Dist 1"},
		{"Revelations 1", true, "Revelation 1", "Revelations -> Dist 1"},
		{"Phillipians 1", true, "Philippians 1", "Phillipians -> Dist 1 (l->ll)"},
		{"Philipians 1", true, "Philippians 1", "Philipians -> Dist 1 (pp->p)"},
		{"Habbakuk 1", true, "Habakkuk 1", "Habbakuk -> Dist 1 (b->bb? no, Habakkuk is correct, Habbakuk is dist 1?)"},
		// Habakkuk: H-a-b-a-k-k-u-k (8)
		// Habbakuk: H-a-b-b-a-k-u-k (8)
		// Dist: b!=b (match), a!=b (sub), k!=a (sub)...
		// Wait. Habakkuk. b, a, k, k, u, k.
		// Habbakuk. b, b, a, k, u, k.
		// Insert b at 2. Delete k at 5. Dist 2. Allowed for length 8.

		// --- Ambiguous Abbreviations ---
		{"Ju 1", false, "", "Ju -> Ambiguous (Jude, Judges). Should likely fail or pick one greedily if defined? (Not defined in map)"},
		{"Jo 1", false, "", "Jo -> Ambiguous (Job, Joel, Jonah, John, Joshua). Not defined."},
		{"Ma 1", false, "", "Ma -> Ambiguous (Matthew, Mark, Malachi). Not defined."},

		// --- Numbered Book Variations ---
		{"I Samuel 1", true, "1 Samuel 1", "Roman Numeral I -> 1 Samuel"},
		{"II Samuel 1", true, "2 Samuel 1", "Roman Numeral II -> 2 Samuel"},
		{"1st John 1", true, "1 John 1", "Ordinal 1st -> 1 John"},
		{"2nd John 1", true, "2 John 1", "Ordinal 2nd -> 2 John"},
		{"3rd John 1", true, "3 John 1", "Ordinal 3rd -> 3 John"},
		{"II Kings 1", true, "2 Kings 1", "Roman Numeral II -> 2 Kings"},
		{"2nd Peter 1", true, "2 Peter 1", "Ordinal 2nd -> 2 Peter"},
		{"1st Corinthians 1", true, "1 Corinthians 1", "Ordinal 1st -> 1 Corinthians"},
		// "1st John" tokens: "1st", "John".
		// "1st John" vs "1 John". Dist 2.
		// Length "1 John" is 6. Threshold is 1.
		// So "1st John" dist 2 > 1. Fail.
		{"First John 1", false, "", "Word First -> Not supported"},

		// --- Negative Cases (Should be False) ---
		{"John 3:16 hello", false, "", "Trailing text"},
		{"hello John 3:16", false, "", "Leading text"},
		{"Read John 3:16", false, "", "Command prefix"},
		{"John is here", false, "", "False positive check"},
		{"Mark my words", false, "", "Mark book match, but trailing text"},
		{"Acts like a fool", false, "", "Acts book match, but trailing text"},
		{"Job description", false, "", "Job book match, but trailing text"},
		{"Numbers are fun", false, "", "Numbers book match, but trailing text"},
		{"Judges decide", false, "", "Judges book match, but trailing text"},
		{"Kings and Queens", false, "", "Kings? No 'Kings' book, '1 Kings'/'2 Kings'"},

		// --- Punctuation at End (Strict check) ---
		{"John 3:16.", false, "", "Trailing period - currently rejected by strict check"},
		{"John 3:16!", false, "", "Trailing exclamation"},
		{"John 3:16?", false, "", "Trailing question mark"},

		// --- Ambiguous / Partial ---
		{"John 3:", false, "", "Trailing separator without digit - currently rejected?"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s (%s)", tt.input, tt.note), func(t *testing.T) {
			parsed, ok := ParseBibleReference(tt.input)
			if ok != tt.expected {
				t.Errorf("ParseBibleReference(%q) ok = %v; want %v", tt.input, ok, tt.expected)
			}
			if ok && parsed != tt.parsed {
				t.Errorf("ParseBibleReference(%q) parsed = %q; want %q", tt.input, parsed, tt.parsed)
			}
		})
	}
}
