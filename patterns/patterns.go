package patterns

import (
	"regexp"
)

// Different kind of input lines supported
const (
	ConstantAssignment     PatternType = iota // assignment of roman value to galatic constants
	MetalCostAssignment                       // assignment of cost to different metals
	GalacticAmountQuestion                    // question about galactic amounts
	MetalCostQuestion                         // question about metals' costs
	InvalidLine                               // None of the valid ones
)

// PatternType is used to identify a pattern regex.
type PatternType int

// Pattern holds the information mapping a PatternType to it's regex.
type Pattern struct {
	Type      PatternType
	PatternRx *regexp.Regexp
}

// Patterns basically holds a slice of patterns
type Patterns []Pattern

// GetTypeOf would return the PatternType of line supplied by the client
// ie. one of the above PatternType constants
func (pats Patterns) GetTypeOf(line string) PatternType {
	for _, pat := range pats {
		if pat.PatternRx.MatchString(line) {
			return pat.Type
		}
	}
	return InvalidLine
}
