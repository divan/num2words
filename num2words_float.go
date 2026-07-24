package num2words

import (
	"fmt"
	"math"
	"strings"
)

// Precision represents the number of decimal places to convert
// -1 means automatic precision (removes trailing zeros)
type Precision int

const (
	AutoPrecision Precision = -1
)

// ConvertFloat converts float number into words representation.
// precision controls decimal places (-1 for automatic precision)
func ConvertFloat(number float64, precision Precision) string {
	return convertFloat(number, precision, false)
}

// ConvertFloatAnd converts float number into words representation
// with " and " added between number groups.
// precision controls decimal places (-1 for automatic precision)
func ConvertFloatAnd(number float64, precision Precision) string {
	return convertFloat(number, precision, true)
}

func convertFloat(number float64, precision Precision, useAnd bool) string {
	// Handle integer part - use Trunc instead of Floor to avoid rounding
	intPart := int(math.Trunc(math.Abs(number)))
	fraction := math.Abs(number) - math.Trunc(math.Abs(number))

	// Special handling for -0.x numbers
	if intPart == 0 && number < 0 {
		result := "zero"
		if fraction == 0 {
			return "zero"
		}
		return "minus " + result + handleDecimalPart(number, precision, useAnd)
	}

	result := convert(intPart, useAnd)

	if fraction == 0 {
		if number < 0 {
			return "minus " + result
		}
		return result
	}

	if number < 0 {
		return "minus " + result + handleDecimalPart(number, precision, useAnd)
	}
	return result + handleDecimalPart(number, precision, useAnd)
}

func handleDecimalPart(number float64, precision Precision, useAnd bool) string {
	// Derive the fractional digits from the canonical decimal STRING of the
	// number, NOT from fraction*math.Pow10(n)+math.Trunc. The float remainder
	// (math.Abs(number) - math.Trunc(math.Abs(number))) is frequently slightly
	// BELOW the true decimal (e.g. 2.3 -> 0.29999999999999982), and crucially it
	// is a DIFFERENT, more-degraded float64 than the nearest float64 to the
	// true decimal fraction -- so even its shortest round-trip is
	// "0.2999999999999998", and multiplying by 10^n + math.Trunc lands one
	// digit low (ConvertFloat(2.3, ...) -> "two point two" instead of
	// "two point three"). Formatting the FULL |number| instead yields the
	// correct shortest round-trip ("2.3"), whose fractional digits are right.
	abs := math.Abs(number)
	var fracStr string
	if precision == AutoPrecision {
		// Preserve the library's 9-digit auto-precision shape: format the
		// fraction with %.9f and trim trailing zeros, but read the digits
		// straight from the string instead of discarding it and recomputing
		// via fraction*math.Pow10(decimalDigits)+math.Trunc.
		fraction := abs - math.Trunc(abs)
		decStr := fmt.Sprintf("%.9f", fraction)
		// decStr starts with "0." for a fraction in [0, 1).
		fracStr = strings.TrimRight(decStr[2:], "0")
	} else {
		// Format |number| with extra precision (n+10) so the IEEE-754
		// representation noise -- which lives well beyond the requested
		// precision -- is rounded away (for below-true-decimal values the
		// trailing 9s carry up to the true digit; for above-true-decimal values
		// the trailing 0s round down). Then TRUNCATE (not round) the result to
		// the requested precision, matching the library's "Use Trunc to avoid
		// any rounding" design. This recovers the correct fractional digits for
		// inputs like 2.3 (-> "30" at prec2) while preserving the truncation
		// behavior the existing tests encode (e.g. 123.45 at prec1 -> "four").
		n := int(precision)
		decStr := fmt.Sprintf("%.*f", n+10, abs)
		if dot := strings.IndexByte(decStr, '.'); dot >= 0 {
			fracStr = decStr[dot+1:]
		} else {
			fracStr = ""
		}
		if len(fracStr) > n {
			fracStr = fracStr[:n]
		}
	}

	if fracStr == "" {
		return ""
	}

	var decimalPart int
	if _, err := fmt.Sscanf(fracStr, "%d", &decimalPart); err != nil {
		return ""
	}
	if decimalPart == 0 {
		return ""
	}

	return " point " + convert(decimalPart, useAnd)
}
