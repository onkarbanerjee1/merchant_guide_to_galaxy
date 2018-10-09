package processors

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/roman"
)

func TestMetalCostQuestion(t *testing.T) {
	initConstantsAssignments()
	initMetalRates()
	defer ClearAssignments()
	processor := NewMetalCostQuestionProcessor(regexp.MustCompile("^how many [c|C]redits is ([A-Za-z\\s]+) ([A-Za-z\\s]+)\\?$"), roman.ToArabic)

	testCases := []struct {
		galacticlines []string
		exp           string
	}{
		{galacticlines: []string{"how many Credits is glob prok Silver ?", "how many Credits is glob prok Gold ?", "how many Credits is glob prok Iron ?"},
			exp: `glob prok Silver is 68 Credits
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
`,
		},
		{galacticlines: []string{"how many credits is glob prok Silver?", "how many credits is glob prok Gold ?", "how many Credits is glob prok Iron ?"},
			exp: `glob prok Silver is 68 Credits
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
`,
		},
		{galacticlines: []string{"how many Credits is glob prok Silver ", "how many Credits is glob prok Gold ?", "how many Credits is glob prok Iron ?"},
			exp: `how many Credits is glob prok Silver  is not a metal cost question line
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
`,
		},
		{galacticlines: []string{"how many Credits is glob prok XXXX Silver ?", "how many Credits is glob prok Gold ?", "how many Credits is glob prok Iron ?"},
			exp: `No roman number symbol assigned for this galactic constant for XXXX
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
`,
		},
		{galacticlines: []string{"how many Credits is glob prok Copper ?", "how many Credits is glob prok Gold ?", "how many Credits is glob prok Iron ?"},
			exp: `Found no credits assigned for this metal , Copper
glob prok Gold is 57800 Credits
glob prok Iron is 782 Credits
`,
		},
	}

	for _, tc := range testCases {
		fakeWriter := &bytes.Buffer{}
		for _, line := range tc.galacticlines {
			if err := processor.Process(line, fakeWriter); err != nil {
				t.Fatal("Expected no error, got", err)
			}
		}
		if act := fakeWriter.String(); act != tc.exp {
			t.Fatal("Expected", tc.exp, ", got", act)
		}
	}
}
