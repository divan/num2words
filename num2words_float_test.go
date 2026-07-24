package num2words

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// TestConvertFloatIEEE754Truncation covers the IEEE-754 truncation regression
// in handleDecimalPart: the float remainder (math.Abs(number) -
// math.Trunc(math.Abs(number))) is slightly BELOW the true decimal for many
// inputs (e.g. 2.3 -> 0.29999999999999982), so the previous
// fraction*math.Pow10(n)+math.Trunc extraction landed one digit low. The fix
// reads the fractional digits from the canonical decimal string instead.
func TestConvertFloatIEEE754Truncation(t *testing.T) {
	testCases := []struct {
		name      string
		number    float64
		precision Precision
		want      string
	}{
		// RED cases fixed by the string-based decimal extraction.
		{"2.3 auto", 2.3, AutoPrecision, "two point three"},
		{"2.3 prec2", 2.3, 2, "two point thirty"},
		{"3.3 auto", 3.3, AutoPrecision, "three point three"},
		{"0.29 prec2", 0.29, 2, "zero point twenty-nine"},
		{"0.58 prec2", 0.58, 2, "zero point fifty-eight"},
		{"5.27 prec2", 5.27, 2, "five point twenty-seven"},
		{"0.57 prec2", 0.57, 2, "zero point fifty-seven"},
		{"7.3 prec2", 7.3, 2, "seven point thirty"},

		// Controls (behavior unchanged, both before and after the fix).
		{"3.14 prec2", 3.14, 2, "three point fourteen"},
		{"123.45 prec2", 123.45, 2, "one hundred twenty-three point forty-five"},
		{"0.42 prec2", 0.42, 2, "zero point forty-two"},
		{"1.5 prec1", 1.5, 1, "one point five"},

		// Negative variants of the fixed cases.
		{"-2.3 auto", -2.3, AutoPrecision, "minus two point three"},
		{"-0.29 prec2", -0.29, 2, "minus zero point twenty-nine"},
	}

	for _, tc := range testCases {
		got := ConvertFloat(tc.number, tc.precision)
		if got != tc.want {
			t.Errorf("ConvertFloat(%v, %v) [%s] = %q; want %q",
				tc.number, tc.precision, tc.name, got, tc.want)
		}
	}
}

// TestConvertFloatAndIEEE754Truncation mirrors the ConvertFloat regression
// suite for the "and"-joined variant.
func TestConvertFloatAndIEEE754Truncation(t *testing.T) {
	testCases := []struct {
		name      string
		number    float64
		precision Precision
		want      string
	}{
		{"2.3 auto", 2.3, AutoPrecision, "two point three"},
		{"2.3 prec2", 2.3, 2, "two point thirty"},
		{"3.3 auto", 3.3, AutoPrecision, "three point three"},
		{"0.29 prec2", 0.29, 2, "zero point twenty-nine"},
		{"0.58 prec2", 0.58, 2, "zero point fifty-eight"},
		{"5.27 prec2", 5.27, 2, "five point twenty-seven"},
		{"0.57 prec2", 0.57, 2, "zero point fifty-seven"},
		{"7.3 prec2", 7.3, 2, "seven point thirty"},

		// Controls.
		{"3.14 prec2", 3.14, 2, "three point fourteen"},
		{"123.45 prec2", 123.45, 2, "one hundred and twenty-three point forty-five"},
		{"0.42 prec2", 0.42, 2, "zero point forty-two"},
		{"1.5 prec1", 1.5, 1, "one point five"},
	}

	for _, tc := range testCases {
		got := ConvertFloatAnd(tc.number, tc.precision)
		if got != tc.want {
			t.Errorf("ConvertFloatAnd(%v, %v) [%s] = %q; want %q",
				tc.number, tc.precision, tc.name, got, tc.want)
		}
	}
}

// TestConvertFloatSweepIEEE754Truncation characterizes the full input space
// of two-decimal fractions (0.01..0.99 at prec2) and one-decimal integer+decimal
// combos (0.1..0.9 at prec1 and auto) against a string-based oracle. Before the
// fix ~36% of these inputs produced wrong output; after the fix the wrong-rate
// must be zero.
func TestConvertFloatSweepIEEE754Truncation(t *testing.T) {
	stringOracle := func(number float64, precision int) string {
		// Build the canonical decimal words form expected from the fix.
		neg := number < 0
		abs := number
		if neg {
			abs = -abs
		}
		intPart := int(abs)
		intWords := Convert(intPart)
		if intPart == 0 {
			intWords = "zero"
		}
		var fracStr string
		if precision < 0 {
			s := strconv.FormatFloat(abs, 'f', -1, 64)
			if dot := strings.IndexByte(s, '.'); dot >= 0 {
				fracStr = strings.TrimRight(s[dot+1:], "0")
			}
		} else {
			s := fmt.Sprintf("%.*f", precision, abs)
			if dot := strings.IndexByte(s, '.'); dot >= 0 {
				fracStr = s[dot+1:]
			}
		}
		out := intWords
		if len(fracStr) > 0 {
			var fracInt int
			fmt.Sscanf(fracStr, "%d", &fracInt)
			if fracInt != 0 {
				out += " point " + Convert(fracInt)
			}
		}
		if neg {
			out = "minus " + out
		}
		return out
	}

	wrong := 0
	total := 0
	// prec2: 0.01..0.99
	for i := 1; i <= 99; i++ {
		f := float64(i) / 100.0
		got := ConvertFloat(f, 2)
		want := stringOracle(f, 2)
		total++
		if got != want {
			wrong++
			if wrong <= 5 {
				t.Errorf("ConvertFloat(%.2f, 2) = %q; want %q", f, got, want)
			}
		}
	}
	// prec1 and auto: int 0..9 + 0.1..0.9
	for n := 0; n <= 9; n++ {
		for i := 1; i <= 9; i++ {
			f := float64(n) + float64(i)/10.0
			for _, prec := range []Precision{1, AutoPrecision} {
				got := ConvertFloat(f, prec)
				p := int(prec)
				want := stringOracle(f, p)
				total++
				if got != want {
					wrong++
					if wrong <= 5 {
						t.Errorf("ConvertFloat(%v, %v) = %q; want %q", f, prec, got, want)
					}
				}
			}
		}
	}
	if wrong != 0 {
		t.Errorf("sweep wrong-rate = %d/%d (%.2f%%); want 0", wrong, total,
			100.0*float64(wrong)/float64(total))
	}
}
