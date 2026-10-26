package cop

import (
	"fmt"
	"strings"
)

func init() {
	for _, h := range []string{
		"professor", "dame",
		"jr", "jnr", "sr", "snr", "junior", "senior",
	} {
		honorifics[h] = true
	}

	pipeline = append(pipeline,
		matcher{"subsequence", 0.78, matchSubsequence},
	)
}

// matchSubsequence fires when one name's tokens are a proper subsequence of the
// other's — covering middle-name drops, name particles (van/de/bint/al) and
// hyphenated-component drops in a single unified check.
func matchSubsequence(a, b string) (float64, string) {
	wa, wb := strings.Fields(a), strings.Fields(b)
	if isSubsequence(wa, wb) {
		return 1, fmt.Sprintf("%v ⊂ %v", wa, wb)
	}
	if isSubsequence(wb, wa) {
		return 1, fmt.Sprintf("%v ⊂ %v", wb, wa)
	}
	return 0, "no subsequence match"
}

// isSubsequence returns true when short is a proper subsequence of long:
// every token of short appears in long in order, at least one token is skipped,
// coverage ≥ 40%, and token matching is initial-aware.
func isSubsequence(short, long []string) bool {
	if len(short) < 2 || len(short) >= len(long) {
		return false
	}
	if float64(len(short))/float64(len(long)) < 0.4 {
		return false
	}
	j := 0
	for i := 0; i < len(long) && j < len(short); i++ {
		if tokenMatches(short[j], long[i]) {
			j++
		}
	}
	return j == len(short)
}

// tokenMatches returns true when a equals b, or one is a single-char initial
// of the other (e.g. "j" matches "john").
func tokenMatches(a, b string) bool {
	if a == b {
		return true
	}
	if len(a) == 1 && strings.HasPrefix(b, a) {
		return true
	}
	if len(b) == 1 && strings.HasPrefix(a, b) {
		return true
	}
	return false
}
