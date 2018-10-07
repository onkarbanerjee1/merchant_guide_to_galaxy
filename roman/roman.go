package roman

import (
	"errors"
	"regexp"
)

var numbers = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// different errors returned by this package
var (
	errInvalidRomanSymbol = errors.New("Invalid roman symbol")
	errInvalidRomanNumber = errors.New("Invalid roman number, violated roman number rules")
)

// ValidateSymbol checks if the given symbol is a valid roman symbol.
func ValidateSymbol(symbol string) error {
	_, found := numbers[symbol]
	if !found {
		return errInvalidRomanSymbol
	}
	return nil
}

// isValidNumber checks if given roman number is a valid roman number.
func isValidNumber(roman string) bool {
	validRomanRx := regexp.MustCompile("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$")
	return validRomanRx.MatchString(roman)
}

// ToArabic returns an arabic number for the given roman number string if it is valid one
func ToArabic(roman string) (int, error) {
	if !isValidNumber(roman) {
		return 0, errInvalidRomanNumber
	}
	lastDigit, arabic := 1000, 0
	for _, c := range []byte(roman) {
		digit := numbers[string(c)]
		if lastDigit < digit {
			arabic -= 2 * lastDigit
		}
		lastDigit = digit
		arabic += lastDigit
	}

	return arabic, nil
}
