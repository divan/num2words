package num2words_test

import (
	"fmt"
	"testing"

	"github.com/divan/num2words"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	Convert := num2words.Convert
	ConvertAnd := num2words.ConvertAnd
	Convey("Should convert correctly", t, func() {
		Convey("Small numbers should convert correctly", func() {
			So(Convert(0), ShouldEqual, "zero")
			So(Convert(1), ShouldEqual, "one")
			So(Convert(5), ShouldEqual, "five")
			So(Convert(10), ShouldEqual, "ten")
			So(Convert(11), ShouldEqual, "eleven")
			So(Convert(12), ShouldEqual, "twelve")
			So(Convert(17), ShouldEqual, "seventeen")
		})
		Convey("Tens should convert correctly", func() {
			So(Convert(20), ShouldEqual, "twenty")
			So(Convert(30), ShouldEqual, "thirty")
			So(Convert(40), ShouldEqual, "forty")
			So(Convert(50), ShouldEqual, "fifty")
			So(Convert(60), ShouldEqual, "sixty")
			So(Convert(90), ShouldEqual, "ninety")
		})
		Convey("Combined numbers should convert correctly", func() {
			So(Convert(21), ShouldEqual, "twenty-one")
			So(Convert(34), ShouldEqual, "thirty-four")
			So(Convert(49), ShouldEqual, "forty-nine")
			So(Convert(53), ShouldEqual, "fifty-three")
			So(Convert(68), ShouldEqual, "sixty-eight")
			So(Convert(99), ShouldEqual, "ninety-nine")
		})
		Convey("Big numbers should convert correctly", func() {
			So(Convert(100), ShouldEqual, "one hundred")
			So(Convert(200), ShouldEqual, "two hundred")
			So(Convert(500), ShouldEqual, "five hundred")
			So(Convert(123), ShouldEqual, "one hundred twenty-three")
			So(Convert(666), ShouldEqual, "six hundred sixty-six")
			So(Convert(1024), ShouldEqual, "one thousand twenty-four")
		})
		Convey("Negative numbers should convert correclty", func() {
			So(Convert(-123), ShouldEqual, "minus one hundred twenty-three")
		})
		Convey("Convert with 'and' should convert correclty", func() {
			So(ConvertAnd(123), ShouldEqual, "one hundred and twenty-three")
			So(ConvertAnd(514), ShouldEqual, "five hundred and fourteen")
			So(ConvertAnd(1111), ShouldEqual, "one thousand one hundred and eleven")
		})
	})
}

func ExampleConvert() {
	fmt.Println(num2words.Convert(11))
	fmt.Println(num2words.Convert(123))
	fmt.Println(num2words.Convert(-99))
	// Output: eleven
	// one hundred twenty-three
	// minus ninety-nine
}

func ExampleConvertAnd() {
	fmt.Println(num2words.ConvertAnd(123))
	fmt.Println(num2words.ConvertAnd(514))
	// Output: one hundred and twenty-three
	// five hundred and fourteen
}

func TestConvertFloat(t *testing.T) {
	ConvertFloat := num2words.ConvertFloat
	ConvertFloatAnd := num2words.ConvertFloatAnd

	Convey("Should convert floats correctly", t, func() {
		Convey("Basic float numbers should convert correctly", func() {
			So(ConvertFloat(123.45, 2), ShouldEqual, "one hundred twenty-three point forty-five")
			So(ConvertFloat(123.45, 1), ShouldEqual, "one hundred twenty-three point four")
			So(ConvertFloat(-123.45, 2), ShouldEqual, "minus one hundred twenty-three point forty-five")
			So(ConvertFloat(0.42, 2), ShouldEqual, "zero point forty-two")
		})

		Convey("Whole numbers should convert without decimal part", func() {
			So(ConvertFloat(1.0, 2), ShouldEqual, "one")
			So(ConvertFloat(100.0, 2), ShouldEqual, "one hundred")
			So(ConvertFloat(-50.0, 2), ShouldEqual, "minus fifty")
		})

		Convey("Auto precision should handle trailing zeros correctly", func() {
			So(ConvertFloat(1.50, num2words.AutoPrecision), ShouldEqual, "one point five")
			So(ConvertFloat(1.500, num2words.AutoPrecision), ShouldEqual, "one point five")
			So(ConvertFloat(100.203, num2words.AutoPrecision), ShouldEqual, "one hundred point two hundred three")
		})

		Convey("Zero values should convert correctly", func() {
			So(ConvertFloat(0.0, 2), ShouldEqual, "zero")
			So(ConvertFloat(0.42, 2), ShouldEqual, "zero point forty-two")
			So(ConvertFloat(-0.42, 2), ShouldEqual, "minus zero point forty-two")
		})
	})

	Convey("Should convert floats with 'and' correctly", t, func() {
		Convey("Basic float numbers should convert correctly with 'and'", func() {
			So(ConvertFloatAnd(123.45, 2), ShouldEqual, "one hundred and twenty-three point forty-five")
			So(ConvertFloatAnd(123.45, 1), ShouldEqual, "one hundred and twenty-three point four")
			So(ConvertFloatAnd(-123.45, 2), ShouldEqual, "minus one hundred and twenty-three point forty-five")
		})

		Convey("Auto precision should work with 'and'", func() {
			So(ConvertFloatAnd(1.50, num2words.AutoPrecision), ShouldEqual, "one point five")
			So(ConvertFloatAnd(100.203, num2words.AutoPrecision), ShouldEqual, "one hundred point two hundred and three")
			So(ConvertFloatAnd(1234.5678, num2words.AutoPrecision), ShouldEqual, "one thousand two hundred and thirty-four point five thousand six hundred and seventy-eight")
		})
	})
}

func ExampleConvertFloat() {
	fmt.Println(num2words.ConvertFloat(123.45, 2))
	fmt.Println(num2words.ConvertFloat(-0.42, 2))
	fmt.Println(num2words.ConvertFloat(1.50, num2words.AutoPrecision))
	// Output:
	// one hundred twenty-three point forty-five
	// minus zero point forty-two
	// one point five
}

func ExampleConvertFloatAnd() {
	fmt.Println(num2words.ConvertFloatAnd(123.45, 2))
	fmt.Println(num2words.ConvertFloatAnd(1234.5678, 4))
	// Output:
	// one hundred and twenty-three point forty-five
	// one thousand two hundred and thirty-four point five thousand six hundred and seventy-eight
}
