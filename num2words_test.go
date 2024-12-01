package num2words

import (
	"testing"
)

func TestConvertFloat(t *testing.T) {
	testCases := []struct {
		number     float64
		precision  int
		want       string
	}{
		{123.45, 2, "one hundred twenty-three point forty-five"},
		{123.45, 1, "one hundred twenty-three point four"},
		{-123.45, 2, "minus one hundred twenty-three point forty-five"},
		{0.42, 2, "zero point forty-two"},
		{3.14, 2, "three point one four"}, // Added test case for "3.14"
	}

	for _, tc := range testCases {
		got := ConvertFloat(tc.number, tc.precision)
		if got != tc.want {
			t.Errorf("ConvertFloat(%f, %d) = %s; want %s", tc.number, tc.precision, got, tc.want)
		}
	}
}

func TestConvertFloatAnd(t *testing.T) {
	testCases := []struct {
		number     float64
		precision  int
		want       string
	}{
		{123.45, 2, "one hundred and twenty-three point forty-five"},
		{123.45, 1, "one hundred and twenty-three point four"},
		{-123.45, 2, "minus one hundred and twenty-three point forty-five"},
		{0.42, 2, "zero point forty-two"},
		{3.14, 2, "three point one four"}, // Added test case for "3.14"
	}

	for _, tc := range testCases {
		got := ConvertFloatAnd(tc.number, tc.precision)
		if got != tc.want {
			t.Errorf("ConvertFloatAnd(%f, %d) = %s; want %s", tc.number, tc.precision, got, tc.want)
		}
	}
}
