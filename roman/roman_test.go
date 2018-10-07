package roman

import (
	"testing"
)

func TestValidateSymbol(t *testing.T) {
	testCases := []struct {
		symbol string
		exp    error
	}{
		{"I", nil},
		{"V", nil},
		{"X", nil},
		{"L", nil},
		{"C", nil},
		{"D", nil},
		{"M", nil},
		{"i", errInvalidRomanSymbol},
		{"", errInvalidRomanSymbol},
		{"A", errInvalidRomanSymbol},
		{"4", errInvalidRomanSymbol},
		{"II", errInvalidRomanSymbol},
	}

	for _, tc := range testCases {
		if act := ValidateSymbol(tc.symbol); act != tc.exp {
			t.Fatal("Expected", tc.exp, ", got", act)
		}
	}
}

func TestValidateNumber(t *testing.T) {
	testCases := []struct {
		number string
		exp    bool
	}{
		{"II", true},
		{"I", true},
		{"XXX", true},
		{"XXXX", false},
		{"XXIX", true},
		{"DD", false},
		{"LL", false},
		{"VV", false},
		{"IM", false},
		{"XM", false},
		{"CM", true},
		{"VM", false},
		{"XM", false},
		{"XXXD", false},
		{"XDM", false},
		{"XXXD", false},
		{"LC", false},
		{"", true},
		{"M", true},
	}

	for _, tc := range testCases {
		if act := isValidNumber(tc.number); act != tc.exp {
			t.Fatal("Expected", tc.exp, ", got", act)
		}
	}
}

func TestToArabic(t *testing.T) {
	testCases := []struct {
		roman  string
		arabic int
		err    error
	}{
		{"I", 1, nil},
		{"II", 2, nil},
		{"III", 3, nil},
		{"IV", 4, nil},
		{"V", 5, nil},
		{"VI", 6, nil},
		{"VII", 7, nil},
		{"VIII", 8, nil},
		{"IX", 9, nil},
		{"X", 10, nil},
		{"XI", 11, nil},
		{"L", 50, nil},
		{"C", 100, nil},
		{"D", 500, nil},
		{"CMXCIX", 999, nil},
		{"M", 1000, nil},
		{"MDCCCLXXXII", 1882, nil},
		{"MDCCCLXXXIII", 1883, nil},
		{"MDCCCLXXXIV", 1884, nil},
		{"MDCCCLXXXV", 1885, nil},
		{"MDCCCLXXXVI", 1886, nil},
		{"MDCCCLXXXVII", 1887, nil},
		{"MDCCCLXXXVIII", 1888, nil},
		{"MDCCCLXXXIX", 1889, nil},
		{"MDCCCXC", 1890, nil},
		{"MCMXLIV", 1944, nil},
		{"MCMXCIX", 1999, nil},
		{"MMM", 3000, nil},
		{"", 0, nil},

		{"i", 0, errInvalidRomanNumber},
		{"A", 0, errInvalidRomanNumber},
		{"4", 0, errInvalidRomanNumber},
		{"IIII", 0, errInvalidRomanNumber},
		{"MMMM", 0, errInvalidRomanNumber},
		{"XXXXX", 0, errInvalidRomanNumber},
		{"VM", 0, errInvalidRomanNumber},
	}

	for _, tc := range testCases {
		if act, err := ToArabic(tc.roman); err != tc.err || act != tc.arabic {
			t.Fatal("Expected", tc.arabic, tc.err, ", got", act, err, "for", tc.roman)
		}
	}
}
