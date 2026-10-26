package main

import (
	"fmt"

	cop "github.com/romanmelnyk/uk-confirmation-of-payee"
)

func main() {
	pairs := [][2]string{
		{"Dr John Smith", "John Smith"},
		{"Dr John Smith", "J Smith"},
		{"Mr John Smith", "John Smith"},
		{"Mrs John Smith", "John Smith"},
		{"Ms John Smith", "John Smith"},
		{"Miss John Smith", "John Smith"},
		{"Prof John Smith", "John Smith"},
		{"Professor John Smith", "J Smith"},
		{"Dr Jane Elizabeth Smith", "Jane Smith"},
		{"Dr Jane Elizabeth Smith", "J E Smith"},
		{"Mrs Jane Elizabeth Smith", "Jane E Smith"},
		{"Mrs Jane Elizabeth Smith", "J Smith"},
		{"Dr Michael James O'Connor", "Michael O'Connor"},
		{"Dr Michael James O'Connor", "M O'Connor"},
		{"Mr Michael James O'Connor", "Mike O'Connor"},
		{"Mrs Mary-Anne Smith-Jones", "Mary Smith-Jones"},
		{"Mrs Mary-Anne Smith-Jones", "M Smith-Jones"},
		{"Dr Sarah McDonald", "Sarah MacDonald"},
		{"Dr Sarah McDonald", "S McDonald"},
		{"Mr Jean-Pierre Dubois", "Jean Dubois"},
		{"Dr Jean-Pierre Dubois", "J P Dubois"},
		{"Mrs Ana María García López", "Ana García"},
		{"Dr Ana María García López", "A M García López"},
		{"Mr Mohammed Al-Fayed", "Mohammed Fayed"},
		{"Dr Mohammed Al-Fayed", "M Al-Fayed"},
		{"Mrs Aisha bint Abdullah", "Aisha Abdullah"},
		{"Dr Robert van der Berg", "Robert Berg"},
		{"Mr Robert van der Berg", "R van der Berg"},
		{"Dr Maria de la Cruz", "Maria Cruz"},
		{"Mrs Maria de la Cruz", "M de la Cruz"},
		{"Dr Alexander von Humboldt", "Alexander Humboldt"},
		{"Sir David Attenborough", "David Attenborough"},
		{"Dame Jane Goodall", "Jane Goodall"},
		{"Dr Emily O'Neill-Smith", "Emily Smith"},
		{"Dr Emily O'Neill-Smith", "E O'Neill"},
		{"Mr William Frederick Churchill", "William Churchill"},
		{"Mr William Frederick Churchill", "W F Churchill"},
		{"Dr Christopher J. Williams", "Chris Williams"},
		{"Dr Christopher J. Williams", "C J Williams"},
		{"Mrs Elizabeth Taylor-Williams", "Elizabeth Williams"},
		{"Mrs Elizabeth Taylor-Williams", "E Taylor-Williams"},
		{"Dr R. Alexander Thompson", "Robert Thompson"},
		{"Dr R. Alexander Thompson", "R A Thompson"},
		{"Professor James McArthur Jr.", "James McArthur"},
		{"Dr James McArthur Jr.", "J McArthur"},
		{"Mrs Patricia Ann O'Brien", "Patricia O'Brien"},
		{"Mrs Patricia Ann O'Brien", "P A O'Brien"},
		{"Dr John Smith", "Mr John Smith"},
		{"Mr John Smith", "Dr John Smith"},
		{"Dr Jane Smith", "Mrs Jane Smith"},
		{"Mrs Jane Smith", "Dr Jane Smith"},
	}

	for _, p := range pairs {
		r := cop.Verify(p[0], p[1])
		fmt.Printf("%-35s | %-25s -> %-12s (%.2f)\n",
			p[0], p[1], r.Verdict, r.Score)
	}
}
