package cop_test

import (
	"testing"

	cop "github.com/romanmelnyk/uk-confirmation-of-payee"
)

func TestVerify(t *testing.T) {
	tests := []struct {
		registered string
		candidate  string
		want       cop.Verdict
	}{
		// Exact
		{"John Smith", "John Smith", cop.Match},
		{"John Smith", "john smith", cop.Match},

		// Honorifics stripped
		{"John Smith", "Mr John Smith", cop.Match},
		{"John Smith", "Dr John Smith", cop.Match},

		// Initials
		{"John Smith", "J Smith", cop.CloseMatch},
		{"John Smith", "J. Smith", cop.CloseMatch},

		// Company suffixes
		{"Acme Limited", "Acme Ltd", cop.Match},
		{"Acme PLC", "Acme Public Limited Company", cop.Match},
		{"Widgets LLP", "Widgets Limited Liability Partnership", cop.Match},

		// Token order
		{"Smith John", "John Smith", cop.Match},

		// Phonetic (Smyth ≈ Smith)
		{"John Smith", "John Smyth", cop.CloseMatch},

		// Clear mismatch
		{"John Smith", "Jane Doe", cop.NoMatch},
		{"Acme Ltd", "Beta Corp", cop.NoMatch},
	}

	for _, tt := range tests {
		t.Run(tt.registered+"_vs_"+tt.candidate, func(t *testing.T) {
			got := cop.Verify(tt.registered, tt.candidate)
			if got.Verdict != tt.want {
				t.Errorf("Verify(%q, %q) = %s (score %.2f), want %s\n  evidence: %v",
					tt.registered, tt.candidate, got.Verdict, got.Score, tt.want, got.Evidence)
			}
		})
	}
}
