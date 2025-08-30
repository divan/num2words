package ukrainian

import (
	"testing"
)

func TestBasicNumbers(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{0, "нуль"},
		{1, "один"},
		{2, "два"},
		{3, "три"},
		{4, "чотири"},
		{5, "п'ять"},
		{10, "десять"},
		{11, "одинадцять"},
		{12, "дванадцять"},
		{19, "дев'ятнадцять"},
		{20, "двадцять"},
		{21, "двадцять один"},
		{30, "тридцять"},
		{42, "сорок два"},
		{99, "дев'яносто дев'ять"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestGenderAgreement(t *testing.T) {
	tests := []struct {
		number   int
		gender   Gender
		expected string
	}{
		// Masculine
		{1, Masculine, "один"},
		{2, Masculine, "два"},
		{21, Masculine, "двадцять один"},
		{22, Masculine, "двадцять два"},

		// Feminine
		{1, Feminine, "одна"},
		{2, Feminine, "дві"},
		{21, Feminine, "двадцять одна"},
		{22, Feminine, "двадцять дві"},

		// Neuter
		{1, Neuter, "одне"},
		{2, Neuter, "два"},
		{21, Neuter, "двадцять одне"},
		{22, Neuter, "двадцять два"},
	}

	for _, test := range tests {
		result := ConvertWithOptions(test.number, Options{Gender: test.gender})
		if result != test.expected {
			t.Errorf("ConvertWithOptions(%d, gender=%v) = %s, expected %s",
				test.number, test.gender, result, test.expected)
		}
	}
}

func TestHundreds(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{100, "сто"},
		{200, "двісті"},
		{300, "триста"},
		{400, "чотириста"},
		{500, "п'ятсот"},
		{600, "шістсот"},
		{700, "сімсот"},
		{800, "вісімсот"},
		{900, "дев'ятсот"},
		{101, "сто один"},
		{111, "сто одинадцять"},
		{123, "сто двадцять три"},
		{999, "дев'ятсот дев'яносто дев'ять"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestThousands(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{1000, "одна тисяча"},
		{2000, "дві тисячі"},
		{3000, "три тисячі"},
		{4000, "чотири тисячі"},
		{5000, "п'ять тисяч"},
		{10000, "десять тисяч"},
		{11000, "одинадцять тисяч"},
		{12000, "дванадцять тисяч"},
		{21000, "двадцять одна тисяча"},
		{22000, "двадцять дві тисячі"},
		{25000, "двадцять п'ять тисяч"},
		{1001, "одна тисяча один"},
		{1234, "одна тисяча двісті тридцять чотири"},
		{2345, "дві тисячі триста сорок п'ять"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestMillions(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{1000000, "один мільйон"},
		{2000000, "два мільйони"},
		{3000000, "три мільйони"},
		{4000000, "чотири мільйони"},
		{5000000, "п'ять мільйонів"},
		{10000000, "десять мільйонів"},
		{11000000, "одинадцять мільйонів"},
		{21000000, "двадцять один мільйон"},
		{22000000, "двадцять два мільйони"},
		{1234567, "один мільйон двісті тридцять чотири тисячі п'ятсот шістдесят сім"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestTeensSpecialCase(t *testing.T) {
	// Test that 11-14 always use genitive plural forms
	tests := []struct {
		number   int
		expected string
	}{
		{11, "одинадцять"},
		{12, "дванадцять"},
		{13, "тринадцять"},
		{14, "чотирнадцять"},
		{111, "сто одинадцять"},
		{1011, "одна тисяча одинадцять"},
		{11000, "одинадцять тисяч"}, // teens use genitive plural
		{12000, "дванадцять тисяч"},
		{13000, "тринадцять тисяч"},
		{14000, "чотирнадцять тисяч"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestNegativeNumbers(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{-1, "мінус один"},
		{-42, "мінус сорок два"},
		{-100, "мінус сто"},
		{-1234, "мінус одна тисяча двісті тридцять чотири"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestWithAnd(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{101, "сто і один"},
		{1001, "одна тисяча і один"},
		{1234, "одна тисяча двісті і тридцять чотири"},
		{123456, "сто і двадцять три тисячі чотириста і п'ятдесят шість"},
	}

	for _, test := range tests {
		result := ConvertAnd(test.number)
		if result != test.expected {
			t.Errorf("ConvertAnd(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestNounForms(t *testing.T) {
	hryvnaForms := &NounForms{
		Singular:         "гривня",
		NominativePlural: "гривні",
		GenitivePlural:   "гривень",
	}

	tests := []struct {
		number   int
		expected string
	}{
		{0, "нуль гривень"},
		{1, "одна гривня"},
		{2, "дві гривні"},
		{3, "три гривні"},
		{4, "чотири гривні"},
		{5, "п'ять гривень"},
		{10, "десять гривень"},
		{11, "одинадцять гривень"}, // teens special case
		{12, "дванадцять гривень"},
		{13, "тринадцять гривень"},
		{14, "чотирнадцять гривень"},
		{15, "п'ятнадцять гривень"},
		{21, "двадцять одна гривня"},
		{22, "двадцять дві гривні"},
		{25, "двадцять п'ять гривень"},
		{100, "сто гривень"}, // 100 uses genitive plural
		{101, "сто одна гривня"},
		{102, "сто дві гривні"},
		{105, "сто п'ять гривень"},
		{11000, "одинадцять тисяч гривень"},
	}

	for _, test := range tests {
		result := ConvertWithOptions(test.number, Options{
			Gender:    Feminine,
			WithNoun:  true,
			NounForms: hryvnaForms,
		})
		if result != test.expected {
			t.Errorf("ConvertWithOptions(%d, with гривня) = %s, expected %s",
				test.number, result, test.expected)
		}
	}
}

func TestNounSelector(t *testing.T) {
	daySelector := func(count int, gender Gender) string {
		if count < 0 {
			count = -count
		}

		lastDigit := count % 10
		lastTwoDigits := count % 100

		// Special case for teens (11-14)
		if lastTwoDigits >= 11 && lastTwoDigits <= 14 {
			return "днів"
		}

		switch lastDigit {
		case 1:
			return "день"
		case 2, 3, 4:
			return "дні"
		default:
			return "днів"
		}
	}

	tests := []struct {
		number   int
		expected string
	}{
		{1, "один день"},
		{2, "два дні"},
		{5, "п'ять днів"},
		{11, "одинадцять днів"},
		{21, "двадцять один день"},
		{22, "двадцять два дні"},
	}

	for _, test := range tests {
		result := ConvertWithOptions(test.number, Options{
			Gender:       Masculine,
			WithNoun:     true,
			NounSelector: daySelector,
		})
		if result != test.expected {
			t.Errorf("ConvertWithOptions(%d, with день) = %s, expected %s",
				test.number, result, test.expected)
		}
	}
}

func TestFloatNumbers(t *testing.T) {
	tests := []struct {
		number   float64
		expected string
	}{
		{0.0, "нуль"},
		{1.0, "один"},
		{3.14, "три кома чотирнадцять"},
		{42.5, "сорок два кома п'ять"},
		{100.01, "сто кома нуль один"},
		{-3.14, "мінус три кома чотирнадцять"},
	}

	for _, test := range tests {
		result := ConvertFloat(test.number, Options{Gender: Masculine})
		if result != test.expected {
			t.Errorf("ConvertFloat(%f) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

func TestComplexNumbers(t *testing.T) {
	tests := []struct {
		number   int
		expected string
	}{
		{1234567890, "один мільярд двісті тридцять чотири мільйони п'ятсот шістдесят сім тисяч вісімсот дев'яносто"},
		{999999999, "дев'ятсот дев'яносто дев'ять мільйонів дев'ятсот дев'яносто дев'ять тисяч дев'ятсот дев'яносто дев'ять"},
	}

	for _, test := range tests {
		result := Convert(test.number)
		if result != test.expected {
			t.Errorf("Convert(%d) = %s, expected %s", test.number, result, test.expected)
		}
	}
}

// Benchmark tests
func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(12345)
	}
}

func BenchmarkConvertWithGender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertWithOptions(12345, Options{Gender: Feminine})
	}
}

func BenchmarkConvertWithNoun(b *testing.B) {
	forms := &NounForms{
		Singular:         "гривня",
		NominativePlural: "гривні",
		GenitivePlural:   "гривень",
	}

	for i := 0; i < b.N; i++ {
		ConvertWithOptions(12345, Options{
			Gender:    Feminine,
			WithNoun:  true,
			NounForms: forms,
		})
	}
}
