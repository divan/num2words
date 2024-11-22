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
		return "minus " + result + handleDecimalPart(fraction, precision, useAnd)
	}

	result := convert(intPart, useAnd)

	if fraction == 0 {
		if number < 0 {
			return "minus " + result
		}
		return result
	}

	if number < 0 {
		return "minus " + result + handleDecimalPart(fraction, precision, useAnd)
	}
	return result + handleDecimalPart(fraction, precision, useAnd)
}

func handleDecimalPart(fraction float64, precision Precision, useAnd bool) string {
	var decimalDigits int
	if precision == AutoPrecision {
		// Remove trailing zeros automatically
		decimalStr := fmt.Sprintf("%.9f", fraction)
		decimalStr = strings.TrimRight(strings.TrimRight(decimalStr[2:], "0"), ".")
		decimalDigits = len(decimalStr)
		if decimalDigits == 0 {
			return ""
		}
		fraction = fraction * math.Pow10(decimalDigits)
	} else {
		decimalDigits = int(precision)
		fraction = fraction * math.Pow10(decimalDigits)
	}

	// Use Trunc to avoid any rounding
	decimalPart := int(math.Trunc(fraction))
	if decimalPart == 0 {
		return ""
	}

	return " point " + convert(decimalPart, useAnd)
}
