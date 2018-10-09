package processors

import (
	"fmt"
	"testing"
)

func TestGetRomanNumber(t *testing.T) {
	initConstantsAssignments()
	defer ClearAssignments()
	testCases := []struct {
		galactic string
		roman    string
		err      error
	}{
		{"glob", "I", nil},
		{"prok", "V", nil},
		{"pish", "X", nil},
		{"tegj", "L", nil},
		{"glob glob", "II", nil},
		{"glob prok", "IV", nil},
		{"prok glob", "VI", nil},
		{"glob glob", "II", nil},
		{"glob glob", "II", nil},
		{"GLOB GLOB", "II", nil},
		{"GLOB GLOB blot", "", fmt.Errorf("%s for %s", errNoConstant, "blot")},
		{"blot", "", fmt.Errorf("%s for %s", errNoConstant, "blot")},
		{"", "0", nil},
	}

	for _, tc := range testCases {
		act, err := getRomanNumber(tc.galactic)
		if err == nil && tc.err != nil ||
			err != nil && tc.err == nil ||
			err != nil && tc.err != nil && err.Error() != tc.err.Error() ||
			act != tc.roman {
			t.Fatal("Expected", tc.roman, tc.err, ", got", act, err, "for", tc.galactic)
		}
	}
}

// helpers
func equalMaps(map1, map2 map[string]string) bool {
	keys := []string{}
	for key := range map1 {
		keys = append(keys, key)
	}
	for _, key := range keys {
		v2, ok2 := map2[key]
		if !ok2 || map1[key] != v2 {
			return false
		}
	}
	return true
}

func equalMetalRates(map1, map2 map[string]float64) bool {
	if len(map1) != len(map2) {
		return false
	}
	keys := []string{}
	for key := range map1 {
		keys = append(keys, key)
	}
	for _, key := range keys {
		v2, ok2 := map2[key]
		if !ok2 || map1[key] != v2 {
			return false
		}
	}
	return true
}

func initNumbers() {
	numbers = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

func initConstantsAssignments() {
	constantsAssignments = map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}
}

func initMetalRates() {
	metalRates = map[string]float64{
		"silver": 17,
		"gold":   14450,
		"iron":   195.5,
	}
}
