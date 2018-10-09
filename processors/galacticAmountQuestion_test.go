package processors

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/roman"
)

func TestGalacticAmountQuestionProcess(t *testing.T) {
	initNumbers()
	initConstantsAssignments()
	defer ClearAssignments()
	processor := NewGalacticAmountQuestionProcessor(regexp.MustCompile("^how much is ([A-Za-z\\s]+)\\?$"), roman.ToArabic)
	testCases := []struct {
		galacticlines []string
		exp           string
	}{
		{galacticlines: []string{"how much is prok ?", "how much is pish tegj glob glob ?", "how much is pish tegj glob ?"},
			exp: `prok is 5
pish tegj glob glob is 42
pish tegj glob is 41
`,
		},
		{galacticlines: []string{"how much is prok ?", "how much is pish tegj glob XXXX ?", "how much is pish tegj glob ?"},
			exp: `prok is 5
No roman number symbol assigned for this galactic constant for XXXX
pish tegj glob is 41
`,
		},
		{galacticlines: []string{"how much is prok ?", "how much is pish pish pish pish ?", "how much is pish tegj glob ?"},
			exp: `prok is 5
Invalid roman number, violated roman number rules
pish tegj glob is 41
`,
		},
		{galacticlines: []string{"how much is prok ?", "abcdefgh", "how much is pish tegj glob ?"},
			exp: `prok is 5
abcdefgh is not a valid galactic amount question line
pish tegj glob is 41
`,
		},
		{galacticlines: []string{"how much is prok "},
			exp: `how much is prok  is not a valid galactic amount question line
`,
		},
		{galacticlines: []string{"how much is prok?"},
			exp: `prok is 5
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
