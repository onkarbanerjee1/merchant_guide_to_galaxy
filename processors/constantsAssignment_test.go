package processors

import (
	"bytes"
	"regexp"
	"testing"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/roman"
)

var numbers map[string]int

func TestConstantsAssignmentProcess(t *testing.T) {
	initNumbers()
	processor := NewConstantsAssignmentProcessor(regexp.MustCompile("^([A-Za-z]+) is ([I|V|X|L|C|D|M])$"), roman.ValidateSymbol)
	testCases := []struct {
		galacticlines        []string
		exp                  string
		constantsAssignments map[string]string
	}{
		{galacticlines: []string{"glob is I", "prok is V", "pish is X", "tegj is L"},
			exp: "",
			constantsAssignments: map[string]string{
				"glob": "I",
				"prok": "V",
				"pish": "X",
				"tegj": "L",
			},
		},
		{galacticlines: []string{"bogus line"},
			exp:                  "bogus line is not a valid constants assignment line\n",
			constantsAssignments: map[string]string{},
		},
		{galacticlines: []string{"glob is I", "bogus line", "pish is X", "tegj is L"},
			exp: "bogus line is not a valid constants assignment line\n",
			constantsAssignments: map[string]string{
				"glob": "I",
				"pish": "X",
				"tegj": "L",
			},
		},
		{galacticlines: []string{"GLOB is I", "PROK is V", "PISH is X", "TEGJ is L"},
			exp: "",
			constantsAssignments: map[string]string{
				"glob": "I",
				"prok": "V",
				"pish": "X",
				"tegj": "L",
			},
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
		if !equalMaps(constantsAssignments, tc.constantsAssignments) {

			t.Fatal("Expected", tc.constantsAssignments, ",got", constantsAssignments)
		}
		constantsAssignments = map[string]string{}
	}
}
