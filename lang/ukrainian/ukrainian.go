// Package ukrainian provides Ukrainian language support for num2words.
package ukrainian

import (
	"fmt"
	"math"
	"strings"
)

// Case represents grammatical case
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
type NounSelector func(count int, gender Gender) string

// Options contains Ukrainian-specific conversion options
type Options struct {
	Gender       Gender       // Gender for number agreement (1-2)
	Case         Case         // Grammatical case (future use)
	UseAnd       bool         // Whether to use "і" between groups
	WithNoun     bool         // Whether to include noun
	NounForms    *NounForms   // Noun forms for agreement
	NounSelector NounSelector // Custom noun selection function
}

// Convert converts a number to Ukrainian words using default options
func Convert(number int) string {
	return ConvertWithOptions(number, Options{
		Gender: Masculine,
		Case:   Nominative,
	})
}

// ConvertAnd converts a number to Ukrainian words with "і" between groups
func ConvertAnd(number int) string {
	return ConvertWithOptions(number, Options{
		Gender: Masculine,
		Case:   Nominative,
		UseAnd: true,
	})
}

// ConvertWithOptions converts a number to Ukrainian words with specified options
func ConvertWithOptions(number int, opts Options) string {
	// Zero rule
	if number == 0 {
		result := smallNumbers[Masculine][0]
		if opts.WithNoun {
			result += " " + selectNoun(0, opts)
		}
		return result
	}

	// Handle negative numbers
	negative := number < 0
	if negative {
		number = -number
	}

	// Divide into three-digit groups
	const groupsNumber = 4
	var groups [groupsNumber]int
	positive := float64(number)

	// Form three-digit groups
	for i := 0; i < groupsNumber; i++ {
		groups[i] = int(math.Mod(positive, 1000))
		positive /= 1000
	}

	var textGroup [groupsNumber]string
	for i := 0; i < groupsNumber; i++ {
		if groups[i] != 0 {
			textGroup[i] = digitGroup2Text(groups[i], opts, i)
		}
	}

	// Combine groups
	combined := textGroup[0]
	useAnd := opts.UseAnd && (groups[0] > 0 && groups[0] < 100)

	for i := 1; i < groupsNumber; i++ {
		if groups[i] != 0 {
			scaleWord := GetScaleWord(groups[i], i)
			prefix := textGroup[i] + " " + scaleWord

			if len(combined) != 0 {
				if useAnd {
					prefix += " і "
				} else {
					prefix += " "
				}
			}

			useAnd = false
			combined = prefix + combined
		}
	}

	if negative {
		combined = "мінус " + combined
	}

	// Add noun if requested
	if opts.WithNoun {
		combined += " " + selectNoun(number, opts)
	}

	return combined
}

// ConvertFloat converts a float number to Ukrainian words
func ConvertFloat(number float64, opts Options) string {
	if number == 0 {
		return smallNumbers[opts.Gender][0]
	}

	// Handle negative numbers
	negative := number < 0
	if negative {
		number = -number
	}

	// Split into integer and fractional parts
	intPart := int(number)
	fracPart := number - float64(intPart)

	result := ConvertWithOptions(intPart, Options{
		Gender: opts.Gender,
		Case:   opts.Case,
		UseAnd: opts.UseAnd,
	})

	// Add fractional part if it exists
	if fracPart > 0 {
		// Convert decimal part intelligently
		fracStr := fmt.Sprintf("%.10f", fracPart)[2:] // Remove "0."
		fracStr = strings.TrimRight(fracStr, "0")      // Remove trailing zeros
		
		if len(fracStr) > 0 {
			// Check if we have leading zeros - if so, read digit by digit
			hasLeadingZero := fracStr[0] == '0'
			
			if hasLeadingZero {
				// Read each digit separately to preserve leading zeros
				var digitWords []string
				for _, digit := range fracStr {
					digitNum := int(digit - '0')
					digitWords = append(digitWords, smallNumbers[opts.Gender][digitNum])
				}
				result += " кома " + strings.Join(digitWords, " ")
			} else {
				// Read as a number
				fracInt := 0
				for _, digit := range fracStr {
					fracInt = fracInt*10 + int(digit-'0')
				}
				result += " кома " + ConvertWithOptions(fracInt, Options{
					Gender: opts.Gender,
					Case:   opts.Case,
				})
			}
		}
	}

	if negative {
		result = "мінус " + result
	}

	return result
}

// selectNoun selects the appropriate noun form based on the number
func selectNoun(count int, opts Options) string {
	if opts.NounSelector != nil {
		return opts.NounSelector(count, opts.Gender)
	}

	if opts.NounForms != nil {
		return selectNounFromForms(count, opts.NounForms)
	}

	return ""
}

// selectNounFromForms selects the appropriate noun form from NounForms
func selectNounFromForms(count int, forms *NounForms) string {
	if count < 0 {
		count = -count
	}

	lastDigit := count % 10
	lastTwoDigits := count % 100

	// Special case for teens (11-14)
	if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
		return forms.GenitivePlural
	}

	switch lastDigit {
	case 1:
		return forms.Singular
	case 2, 3, 4:
		return forms.NominativePlural
	default: // 0, 5, 6, 7, 8, 9
		return forms.GenitivePlural
	}
}

// digitGroup2Text converts a three-digit group to text
func digitGroup2Text(group int, opts Options, groupIndex int) string {
	if group == 0 {
		return ""
	}

	hundreds := group / 100
	tensUnits := group % 100

	var result string

	// Handle hundreds
	if hundreds != 0 {
		if hundreds == 1 {
			// For 1xx, always just say "сто"
			result += GetHundredsWord(hundreds)
		} else {
			// For 2-9 hundreds, Ukrainian uses special forms without the digit
			result += GetHundredsWord(hundreds)
		}
		
		if tensUnits != 0 {
			if opts.UseAnd {
				result += " і "
			} else {
				result += " "
			}
		}
	}

	// Handle tens and units
	if tensUnits != 0 {
		result += convertTensUnits(tensUnits, opts, groupIndex)
	}

	return result
}

// convertTensUnits converts the tens and units part
func convertTensUnits(tensUnits int, opts Options, groupIndex int) string {
	tens := tensUnits / 10
	units := tensUnits % 10

	// Determine gender for this group
	gender := opts.Gender
	// Special handling for thousands (group 1) - they are feminine in Ukrainian
	if groupIndex == 1 {
		gender = Feminine
	}

	if tens >= 2 {
		result := tensWords[tens]
		if units != 0 {
			result += " " + smallNumbers[gender][units]
		}
		return result
	} else if tensUnits != 0 {
		return smallNumbers[gender][tensUnits]
	}

	return ""
}