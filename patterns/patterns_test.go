package patterns

import (
	"regexp"
	"testing"
)

// here we test the GetTypeOf(). So scope is only to assert that we get the right type based on statement.
// we dont validate the number here or the metal in the statement.
func TestGetTypeOf(t *testing.T) {
	pats := Patterns{
		Pattern{Type: ConstantAssignment, PatternRx: regexp.MustCompile("^([A-Za-z]+) is ([I|V|X|L|C|D|M])$")},
		Pattern{Type: MetalCostAssignment, PatternRx: regexp.MustCompile("^([A-Za-z\\s]+) ([A-Za-z]+) is ([0-9]*[.]?[0-9]+) [c|C]redits$")},
		Pattern{Type: GalacticAmountQuestion, PatternRx: regexp.MustCompile("^how much is ([A-Za-z\\s]+)\\?$")},
		Pattern{Type: MetalCostQuestion, PatternRx: regexp.MustCompile("^how many [c|C]redits is ([A-Za-z\\s]+) ([A-Za-z\\s]+)\\?$")},
	}

	testCases := []struct {
		line string
		exp  PatternType
	}{
		{"glob is I", ConstantAssignment},
		{"prok is V", ConstantAssignment},
		{"glob glob Silver is 34 Credits", MetalCostAssignment},
		{"glob glob Silver is 34 credits", MetalCostAssignment},
		{"glob gLOb Silver is 34.8 credits", MetalCostAssignment},
		{"glob gLOb Silver is 0 credits", MetalCostAssignment},
		{"glob gLOb Silver is 0.5 credits", MetalCostAssignment},
		{"glob gLOb Silver is .5 credits", MetalCostAssignment},
		{"how much is pish tegj glob glob ?", GalacticAmountQuestion},
		{"how much is pish pish pish tegj glob glob ?", GalacticAmountQuestion},
		{"how much is pish pish pish tegj glob glob?", GalacticAmountQuestion},
		{"how much is glob ?", GalacticAmountQuestion},
		{"how much is glob glob glob glob?", GalacticAmountQuestion},
		{"how many Credits is glob prok Silver ?", MetalCostQuestion},
		{"how many Credits is Glob Prok Silver ?", MetalCostQuestion},
		{"how many Credits is tegj Silver ?", MetalCostQuestion},
		{"how many Credits is glob Silver ?", MetalCostQuestion},
		{"how many Credits is glob prok pish tegj Copper ?", MetalCostQuestion},
		{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", InvalidLine},
		{" prok is V", InvalidLine},                   // starts with space
		{"prok is V ", InvalidLine},                   //ends with space
		{"glob glob Silver is 1 Credit", InvalidLine}, // missing s i credit even if it is singular amount
		{"prok is A", InvalidLine},                    // A cannot be matched against roman symbols
		{"", InvalidLine},                             // empty line matches nothing
	}

	for _, tc := range testCases {
		if act := pats.GetTypeOf(tc.line); act != tc.exp {
			t.Fatal("Expected", tc.exp, ",got", act)
		}
	}
}
