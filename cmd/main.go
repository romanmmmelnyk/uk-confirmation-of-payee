package main

import (
	"fmt"

	cop "github.com/romanmelnyk/uk-confirmation-of-payee"
)

func main() {

pairs := [][2]string{
	// Titles
	{"Dr John Smith", "John Smith"},
	{"Dr John Smith", "J Smith"},
	{"Mr John Smith", "John Smith"},
	{"Mrs John Smith", "John Smith"},
	{"Ms John Smith", "John Smith"},
	{"Miss John Smith", "John Smith"},
	{"Mx John Smith", "John Smith"},
	{"Prof John Smith", "John Smith"},
	{"Professor John Smith", "J Smith"},
	{"Sir John Smith", "John Smith"},
	{"Dame Jane Smith", "Jane Smith"},
	{"Rev John Smith", "John Smith"},
	{"The Honourable John Smith", "John Smith"},

	// Middle names
	{"Dr Jane Elizabeth Smith", "Jane Smith"},
	{"Dr Jane Elizabeth Smith", "J E Smith"},
	{"Mrs Jane Elizabeth Smith", "Jane E Smith"},
	{"Mrs Jane Elizabeth Smith", "J Smith"},
	{"John Michael David Smith", "John David Smith"},
	{"John Michael David Smith", "John Smith"},
	{"John Michael David Smith", "J M D Smith"},
	{"Robert James Alexander Brown", "Robert A Brown"},
	{"Robert James Alexander Brown", "R J A Brown"},
	{"Emily Rose Catherine Wilson", "Emily Catherine Wilson"},
	{"Emily Rose Catherine Wilson", "Emily R C Wilson"},

	// Apostrophes
	{"Dr Michael James O'Connor", "Michael O'Connor"},
	{"Dr Michael James O'Connor", "M O'Connor"},
	{"Mr Michael O'Connor", "Michael O Connor"},
	{"Mr O'Connor", "OConnor"},
	{"Mrs Siobhan O'Brien", "Siobhan Obrien"},
	{"Dr Patrick O'Neill", "Patrick O'Neil"},
	{"Dr Emily O'Neill-Smith", "Emily Smith"},
	{"Dr Emily O'Neill-Smith", "E O'Neill"},

	// Hyphenated first names and surnames
	{"Mrs Mary-Anne Smith-Jones", "Mary Smith-Jones"},
	{"Mrs Mary-Anne Smith-Jones", "M Smith-Jones"},
	{"Mary-Anne Smith-Jones", "Mary Anne Smith Jones"},
	{"Jean-Pierre Dubois", "Jean Dubois"},
	{"Jean-Pierre Dubois", "J P Dubois"},
	{"Anne-Marie Brown", "Anne Marie Brown"},
	{"Anne-Marie Brown", "Anne Brown"},
	{"Elizabeth Taylor-Williams", "Elizabeth Williams"},
	{"Elizabeth Taylor-Williams", "E Taylor-Williams"},
	{"Peter Smith-Jones", "Peter Smith"},
	{"Peter Smith-Jones", "Peter Jones"},

	// Spelling variants
	{"Dr Sarah McDonald", "Sarah MacDonald"},
	{"Dr Sarah McDonald", "S McDonald"},
	{"Mohammed Al-Fayed", "Mohammed Fayed"},
	{"Mohammed Al-Fayed", "M Al-Fayed"},
	{"Muhammad Khan", "Mohammed Khan"},
	{"Katherine Johnson", "Catherine Johnson"},
	{"Steven Wilson", "Stephen Wilson"},
	{"Jon Smith", "John Smith"},
	{"Mike O'Connor", "Michael O'Connor"},
	{"Bob Williams", "Robert Williams"},
	{"Liz Taylor", "Elizabeth Taylor"},
	{"Alex Thompson", "Alexander Thompson"},

	// Accents and Unicode
	{"Ana María García López", "Ana García"},
	{"Ana María García López", "A M García López"},
	{"Jose Garcia", "José García"},
	{"François Martin", "Francois Martin"},
	{"Søren Jensen", "Soren Jensen"},
	{"Renée Dubois", "Renee Dubois"},
	{"Müller Schmidt", "Muller Schmidt"},
	{"Łukasz Kowalski", "Lukasz Kowalski"},
	{"Zoë Williams", "Zoe Williams"},

	// Particles in surnames
	{"Robert van der Berg", "Robert Berg"},
	{"Robert van der Berg", "R van der Berg"},
	{"Maria de la Cruz", "Maria Cruz"},
	{"Maria de la Cruz", "M de la Cruz"},
	{"Alexander von Humboldt", "Alexander Humboldt"},
	{"Ludwig van Beethoven", "Ludwig Beethoven"},
	{"Leonardo da Vinci", "Leonardo Vinci"},
	{"Charles de Gaulle", "Charles Gaulle"},
	{"Vincent van Gogh", "Vincent Gogh"},
	{"Aisha bint Abdullah", "Aisha Abdullah"},
	{"Abdul Rahman Al Saud", "Abdul Rahman Saud"},

	// Initials
	{"Christopher James Williams", "C J Williams"},
	{"Christopher James Williams", "CJ Williams"},
	{"Christopher J. Williams", "Chris Williams"},
	{"R Alexander Thompson", "Robert Thompson"},
	{"R Alexander Thompson", "R A Thompson"},
	{"J Smith", "John Smith"},
	{"J Smith", "Jane Smith"},
	{"J R R Tolkien", "John Tolkien"},
	{"A B Smith", "A Smith"},
	{"A B Smith", "B Smith"},

	// Suffixes
	{"James McArthur Jr.", "James McArthur"},
	{"James McArthur Sr.", "James McArthur"},
	{"John Smith III", "John Smith"},
	{"John Smith II", "John Smith"},
	{"Robert Downey Jr.", "Robert Downey"},
	{"Henry Ford IV", "Henry Ford"},

	// Different name order
	{"Smith John", "John Smith"},
	{"Garcia Lopez Ana Maria", "Ana Maria Garcia Lopez"},
	{"Wang Xiaoming", "Xiaoming Wang"},
	{"Lee Min-ho", "Min-ho Lee"},

	// Missing or abbreviated first names
	{"William Frederick Churchill", "William Churchill"},
	{"William Frederick Churchill", "W F Churchill"},
	{"Patricia Ann O'Brien", "Patricia O'Brien"},
	{"Patricia Ann O'Brien", "P A O'Brien"},
	{"Elizabeth Alexandra Mary Windsor", "Elizabeth Windsor"},
	{"Elizabeth Alexandra Mary Windsor", "E A M Windsor"},

	// Exact and near-exact matches
	{"John Smith", "John Smith"},
	{"JOHN SMITH", "John Smith"},
	{"john smith", "John Smith"},
	{"  John Smith  ", "John Smith"},
	{"John  Smith", "John Smith"},
	{"John-Smith", "John Smith"},
	{"John Smith.", "John Smith"},

	// Similar names that should not be blindly treated as equivalent
	{"John Smith", "Jon Smith"},
	{"John Smith", "Jane Smith"},
	{"John Smith", "John Smyth"},
	{"Peter Jones", "Peter Johnson"},
	{"Michael Brown", "Michelle Brown"},
	{"Ann Williams", "Anne Williams"},
	{"Sara Khan", "Sarah Khan"},
	{"Robert Taylor", "Albert Taylor"},
	{"James Wilson", "Jameson Wilson"},
	{"John Smith", "John Smithson"},
	{"David Miller", "Daniel Miller"},
	{"Maria Garcia", "Mario Garcia"},

	// Same surname, different first name
	{"John Smith", "James Smith"},
	{"Robert Brown", "Richard Brown"},
	{"Sarah Jones", "Sophie Jones"},
	{"Michael Wilson", "Matthew Wilson"},

	// Same first name, different surname
	{"John Smith", "John Jones"},
	{"Jane Brown", "Jane Wilson"},
	{"Mohammed Khan", "Mohammed Ali"},
	{"Emily Taylor", "Emily Turner"},

	// Title differences
	{"Dr John Smith", "Mr John Smith"},
	{"Mr John Smith", "Dr John Smith"},
	{"Dr Jane Smith", "Mrs Jane Smith"},
	{"Mrs Jane Smith", "Dr Jane Smith"},
	{"Professor James Wilson", "Prof James Wilson"},
	{"Sir David Attenborough", "David Attenborough"},
	{"Dame Jane Goodall", "Jane Goodall"},

	// Empty and incomplete names
	{"", "John Smith"},
	{"John Smith", ""},
	{"", ""},
	{"John", "John Smith"},
	{"Smith", "John Smith"},
	{"J", "John Smith"},
	{"John Smith", "Smith"},
	{"Dr", "John Smith"},
	{"John Smith", "John"},
}


	for _, p := range pairs {
		r := cop.Verify(p[0], p[1])
		fmt.Printf("%-35s | %-25s -> %-12s (%.2f)\n",
			p[0], p[1], r.Verdict, r.Score)
	}
}
