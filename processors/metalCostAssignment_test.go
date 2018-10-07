package processors

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/roman"
)

func TestMetalCostAssignmentProcess(t *testing.T) {
	initNumbers()
	initConstantsAssignments()
	processor := NewMetalCostAssignmentProcessor(regexp.MustCompile("^([A-Za-z\\s]+) ([A-Za-z]+) is ([0-9]*[.]?[0-9]+) [c|C]redits$"), roman.ToArabic)
	testCases := []struct {
		lines      []string
		exp        error
		out        string
		metalRates map[string]float64
	}{
		{lines: []string{"glob glob Silver is 34 Credits", "glob prok Gold is 57800 Credits", "pish pish Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"silver": 17,
				"gold":   14450,
				"iron":   195.5,
			},
		},
		{lines: []string{"glob glob Silver is 34 Credits", "glob prok Gold is AA Credits", "pish pish Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"silver": 17,
				"iron":   195.5,
			},
			out: "glob prok Gold is AA Credits is not a valid metal cost assignment line\n",
		},
		{lines: []string{"glob glob Silver is 34 Credits", "glob glob glob glob Gold is 57800 Credits", "pish pish Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"silver": 17,
				"iron":   195.5,
			},
			out: "Invalid roman number, violated roman number rules\n",
		},
		{lines: []string{"glob glob Silver is 34 Credits", "glob blot glob Gold is 57800 Credits", "pish pish Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"silver": 17,
				"iron":   195.5,
			},
			out: "No roman number symbol assigned for this galactic constant for blot\n",
		},
		{lines: []string{"GLOB glob Silver is 34 Credits", "glob PROK Gold is 578.123 Credits", "pish piSH Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"silver": 17,
				"gold":   144.53075,
				"iron":   195.5,
			},
		},
		{lines: []string{"", "glob PROK Gold is 57800 Credits?", "pish piSH Iron is 3910 Credits"},
			metalRates: map[string]float64{
				"iron": 195.5,
			},
			out: " is not a valid metal cost assignment line\nglob PROK Gold is 57800 Credits? is not a valid metal cost assignment line\n",
		},
	}

	for _, tc := range testCases {
		fakeWriter := &bytes.Buffer{}
		for _, line := range tc.lines {
			if act := processor.Process(line, fakeWriter); act != tc.exp {
				t.Fatal("Expected", tc.exp, ",got", act)
			}
		}
		if out := fakeWriter.String(); out != tc.out {
			t.Fatal("Expected", tc.out, ", got", out)
		}
		if !equalMetalRates(metalRates, tc.metalRates) {
			t.Fatal("Expected", tc.metalRates, ",got", metalRates)
		}
		metalRates = map[string]float64{}
	}
}
