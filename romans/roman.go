package romans

import "strings"

type RomanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(i int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for i >= numeral.Value {
			result.WriteString(numeral.Symbol)
			i -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (total int) {
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(roman, numeral.Symbol) {
			total += numeral.Value
			roman = strings.TrimPrefix(roman, numeral.Symbol)
		}
	}

	return
}
