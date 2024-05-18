package romans

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Number int
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{95, "XCV"},
	{80, "LXXX"},
	{105, "CV"},
	{500, "D"},
	{400, "CD"},
	{490, "CDXC"},
	{1000, "M"},
	{900, "CM"},
	{1984, "MCMLXXXIV"},
}

func TestRomainNumerals(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprintf("convert %d to %q", c.Number, c.Roman), func(t *testing.T) {
			got := ConvertToRoman(c.Number)

			if got != c.Roman {
				t.Errorf("got %q, want %q", got, c.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Number), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Number {
				t.Errorf("got %d, want %d", got, test.Number)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed chekcs", err)
	}
}
