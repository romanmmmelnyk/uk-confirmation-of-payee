package cop

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type matchFn func(a, b string) (confidence float64, reason string)

type matcher struct {
	name   string
	weight float64
	run    matchFn
}

// pipeline is the ordered set of evidence matchers.
// Each weight is chosen so that when a matcher fires (confidence=1),
// the resulting score lands in the correct verdict tier:
//   exact / suffix / tokenSort → ≥ 0.90 → MATCH
//   initials / phonetic        → ≥ 0.65 → CLOSE_MATCH
var pipeline = []matcher{
	{"exact",     1.00, matchExact},
	{"suffix",    0.95, matchSuffix},
	{"tokenSort", 0.92, matchTokenSort},
	{"initials",  0.75, matchInitials},
	{"phonetic",  0.68, matchPhonetic},
}

// ── Normalizer ───────────────────────────────────────────────────────────────

var honorifics = map[string]bool{
	"mr": true, "mrs": true, "ms": true, "miss": true,
	"dr": true, "prof": true, "rev": true, "sir": true,
	"lord": true, "lady": true, "the": true,
}

// normalize strips honorifics, lowercases, removes punctuation,
// and collapses whitespace. All matchers operate on normalised strings.
func normalize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))

	var buf strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || r == ' ' {
			buf.WriteRune(r)
		} else if r == '&' {
			buf.WriteString(" and ")
		}
	}

	words := strings.Fields(buf.String())
	out := words[:0]
	for _, w := range words {
		if !honorifics[w] {
			out = append(out, w)
		}
	}
	return strings.Join(out, " ")
}

// ── Exact ────────────────────────────────────────────────────────────────────

func matchExact(a, b string) (float64, string) {
	if a == b {
		return 1, "exact match"
	}
	return 0, "no exact match"
}

// ── Company suffix ───────────────────────────────────────────────────────────

// suffixRules maps legal forms to canonical abbreviations, longest first
// to prevent "limited" substituting before "limited partnership".
var suffixRules = []struct{ long, short string }{
	{"charitable incorporated organisation", "cio"},
	{"charitable incorporated organization", "cio"},
	{"public limited company",               "plc"},
	{"limited liability partnership",        "llp"},
	{"community interest company",           "cic"},
	{"limited partnership",                  "lp"},
	{"limited",                              "ltd"},
	{"trading as",                           "ta"},
}

func canonicaliseSuffix(s string) string {
	for _, r := range suffixRules {
		s = strings.ReplaceAll(s, " "+r.long, " "+r.short)
	}
	return strings.TrimSpace(s)
}

func matchSuffix(a, b string) (float64, string) {
	ca, cb := canonicaliseSuffix(a), canonicaliseSuffix(b)
	if ca == a && cb == b {
		return 0, "not applicable: no suffix present"
	}
	if ca == cb {
		return 1, "match after suffix normalisation"
	}
	return 0, "suffix mismatch"
}

// ── Token sort ───────────────────────────────────────────────────────────────

func matchTokenSort(a, b string) (float64, string) {
	if tokenSorted(a) == tokenSorted(b) {
		return 1, "match after token sort"
	}
	return 0, "no match after token sort"
}

func tokenSorted(s string) string {
	words := strings.Fields(s)
	sort.Strings(words)
	return strings.Join(words, " ")
}

// ── Initials expansion ───────────────────────────────────────────────────────

func matchInitials(a, b string) (float64, string) {
	if expandsTo(a, b) {
		return 1, fmt.Sprintf("%q is an initials form of %q", a, b)
	}
	if expandsTo(b, a) {
		return 1, fmt.Sprintf("%q is an initials form of %q", b, a)
	}
	return 0, "no initials match"
}

// expandsTo returns true when abbreviated looks like an initialled form of full,
// e.g. "j smith" → "john smith". Requires at least one actual initial so
// identical strings do not falsely trigger this matcher.
func expandsTo(abbreviated, full string) bool {
	wa := strings.Fields(abbreviated)
	wf := strings.Fields(full)
	if len(wa) == 0 || len(wa) != len(wf) {
		return false
	}
	hasInitial := false
	for i := range wa {
		if wa[i] == wf[i] {
			continue
		}
		if len(wa[i]) == 1 && strings.HasPrefix(wf[i], wa[i]) {
			hasInitial = true
			continue
		}
		return false
	}
	return hasInitial
}

// ── Phonetic (Soundex) ───────────────────────────────────────────────────────

func matchPhonetic(a, b string) (float64, string) {
	wa, wb := strings.Fields(a), strings.Fields(b)
	if len(wa) == 0 || len(wa) != len(wb) {
		return 0, "not applicable: token count mismatch"
	}
	for i := range wa {
		// single-char tokens are initials, not full words — skip phonetic comparison
		if len(wa[i]) == 1 || len(wb[i]) == 1 {
			continue
		}
		sa, sb := soundex(wa[i]), soundex(wb[i])
		if sa != sb {
			return 0, fmt.Sprintf("soundex mismatch: %s(%s) ≠ %s(%s)", sa, wa[i], sb, wb[i])
		}
	}
	return 1, "phonetic match via soundex"
}

var soundexTable = map[rune]byte{
	'B': '1', 'F': '1', 'P': '1', 'V': '1',
	'C': '2', 'G': '2', 'J': '2', 'K': '2', 'Q': '2', 'S': '2', 'X': '2', 'Z': '2',
	'D': '3', 'T': '3',
	'L': '4',
	'M': '5', 'N': '5',
	'R': '6',
}

func soundex(s string) string {
	s = strings.ToUpper(s)
	if s == "" {
		return "0000"
	}
	code := []byte{s[0]}
	prev := soundexTable[rune(s[0])]
	for _, r := range s[1:] {
		d, ok := soundexTable[r]
		if ok && d != prev {
			code = append(code, d)
			if len(code) == 4 {
				break
			}
		}
		if ok {
			prev = d
		} else {
			prev = 0
		}
	}
	for len(code) < 4 {
		code = append(code, '0')
	}
	return string(code)
}
