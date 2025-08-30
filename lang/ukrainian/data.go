// Package ukrainian contains Ukrainian language data for number conversion.
package ukrainian

// Gender represents grammatical gender
type Gender int

const (
	Masculine Gender = iota
	Feminine
	Neuter
)

// smallNumbers contains basic numbers 0-19 for each gender
var smallNumbers = map[Gender][]string{
	Masculine: {
		"нуль", "один", "два", "три", "чотири", "п'ять", "шість", "сім", "вісім", "дев'ять",
		"десять", "одинадцять", "дванадцять", "тринадцять", "чотирнадцять", "п'ятнадцять",
		"шістнадцять", "сімнадцять", "вісімнадцять", "дев'ятнадцять",
	},
	Feminine: {
		"нуль", "одна", "дві", "три", "чотири", "п'ять", "шість", "сім", "вісім", "дев'ять",
		"десять", "одинадцять", "дванадцять", "тринадцять", "чотирнадцять", "п'ятнадцять",
		"шістнадцять", "сімнадцять", "вісімнадцять", "дев'ятнадцять",
	},
	Neuter: {
		"нуль", "одне", "два", "три", "чотири", "п'ять", "шість", "сім", "вісім", "дев'ять",
		"десять", "одинадцять", "дванадцять", "тринадцять", "чотирнадцять", "п'ятнадцять",
		"шістнадцять", "сімнадцять", "вісімнадцять", "дев'ятнадцять",
	},
}

// tensWords contains tens from 20-90
var tensWords = []string{
	"", "", "двадцять", "тридцять", "сорок", "п'ятдесят",
	"шістдесят", "сімдесят", "вісімдесят", "дев'яносто",
}

// scaleNumbers contains scale words (thousand, million, etc.)
var scaleNumbers = []string{
	"", "тисяч", "мільйонів", "мільярдів",
}

// Special scale forms for proper agreement
var scaleNumbersForms = map[int]map[string]string{
	1: { // thousands - feminine
		"1":   "тисяча",      // одна тисяча
		"2-4": "тисячі",      // дві тисячі, три тисячі, чотири тисячі
		"5+":  "тисяч",       // п'ять тисяч
	},
	2: { // millions - masculine
		"1":   "мільйон",     // один мільйон
		"2-4": "мільйони",    // два мільйони
		"5+":  "мільйонів",   // п'ять мільйонів
	},
	3: { // billions - masculine
		"1":   "мільярд",     // один мільярд
		"2-4": "мільярди",    // два мільярди
		"5+":  "мільярдів",   // п'ять мільярдів
	},
}

// hundredsWord is the word for "hundred"
const hundredsWord = "сто"

// Special hundreds forms for agreement
var hundredsForms = map[int]string{
	1: "сто",         // один сто
	2: "двісті",      // два двісті -> just "двісті"
	3: "триста",      // три триста -> just "триста"
	4: "чотириста",   // чотири чотириста -> just "чотириста"
	5: "п'ятсот",     // п'ять п'ятсот -> just "п'ятсот"
	6: "шістсот",     // шість шістсот -> just "шістсот"
	7: "сімсот",      // сім сімсот -> just "сімсот"
	8: "вісімсот",    // вісім вісімсот -> just "вісімсот"
	9: "дев'ятсот",   // дев'ять дев'ятсот -> just "дев'ятсот"
}

// GetScaleWord returns the appropriate scale word based on the number
func GetScaleWord(number int, scaleIndex int) string {
	if scaleIndex == 0 {
		return ""
	}

	forms, exists := scaleNumbersForms[scaleIndex]
	if !exists {
		return scaleNumbers[scaleIndex]
	}

	lastDigit := number % 10
	lastTwoDigits := number % 100

	// Special case for teens (11-14)
	if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
		return forms["5+"]
	}

	switch lastDigit {
	case 1:
		return forms["1"]
	case 2, 3, 4:
		return forms["2-4"]
	default:
		return forms["5+"]
	}
}

// GetHundredsWord returns the appropriate hundreds word
func GetHundredsWord(hundreds int) string {
	if form, exists := hundredsForms[hundreds]; exists {
		return form
	}
	return hundredsWord
}