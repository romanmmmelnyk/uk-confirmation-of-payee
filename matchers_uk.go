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
