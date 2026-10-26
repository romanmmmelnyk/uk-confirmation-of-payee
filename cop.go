package cop

import "math"

// Verdict is the outcome of a Confirmation of Payee name check.
type Verdict string

const (
	Match      Verdict = "MATCH"
	CloseMatch Verdict = "CLOSE_MATCH"
	NoMatch    Verdict = "NO_MATCH"
)

// Evidence records a single matcher's contribution to the final score.
type Evidence struct {
	Matcher    string
	Confidence float64
	Reason     string
}

// Result is the full output of a CoP name verification.
type Result struct {
	Verdict  Verdict
	Score    float64
	Evidence []Evidence
}

const (
	thresholdMatch      = 0.90
	thresholdCloseMatch = 0.65
)

// Verify checks whether candidate matches the registered account name.
// It returns MATCH, CLOSE_MATCH, or NO_MATCH with a full evidence trail.
func Verify(registered, candidate string) Result {
	a := normalize(registered)
	b := normalize(candidate)

	evidence := make([]Evidence, len(pipeline))
	for i, m := range pipeline {
		conf, reason := m.run(a, b)
		evidence[i] = Evidence{Matcher: m.name, Confidence: conf, Reason: reason}
	}

	score := aggregate(evidence)
	return Result{
		Verdict:  decide(score),
		Score:    round2(score),
		Evidence: evidence,
	}
}

// aggregate returns the strongest evidence: max(confidence × weight).
// A weighted average would let silent matchers dilute one strong signal;
// instead we ask "what is the best case made for a match?"
func aggregate(evidence []Evidence) float64 {
	var best float64
	for i, e := range evidence {
		if s := e.Confidence * pipeline[i].weight; s > best {
			best = s
		}
	}
	return best
}

func decide(score float64) Verdict {
	switch {
	case score >= thresholdMatch:
		return Match
	case score >= thresholdCloseMatch:
		return CloseMatch
	default:
		return NoMatch
	}
}

func round2(f float64) float64 { return math.Round(f*100) / 100 }
