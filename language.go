// Package num2words provides language definitions for number-to-words conversion.
package num2words

// Gender represents grammatical gender for languages that require it
type Gender int

const (
	Masculine Gender = iota
	Feminine
	Neuter
)

// Case represents grammatical case for languages that require it
type Case int

const (
	Nominative Case = iota
	Genitive
	Dative
	Accusative
	Instrumental
	Locative
	Vocative
)

// NounForms contains different forms of a noun for proper agreement with numbers
type NounForms struct {
	Singular        string // Used with 1: "один день"
	NominativePlural string // Used with 2-4: "два дні" 
	GenitivePlural  string // Used with 5+: "п'ять днів"
}

// NounSelector is a function that selects the appropriate noun form
// based on the number and gender context
type NounSelector func(count int, gender Gender) string