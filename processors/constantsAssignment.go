package processors

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// constantsAssignmentProcessor is a Processor that can Process constants assignment galactic
// lines like "glob is I", "prok is V" etc.
type constantsAssignmentProcessor struct {
	rx       *regexp.Regexp // to match the subsstrings in a constanstsAssigment line
	validate fnValidate     // We will need a function to validate a given constant
}

// NewConstantsAssignmentProcessor returns a Processor that will process constantAssignments
func NewConstantsAssignmentProcessor(rx *regexp.Regexp, validateSymbol fnValidate) Processor {
	return constantsAssignmentProcessor{rx: rx, validate: validateSymbol}
}

// constantsAssignmentProcessor's implementation of how to Process a galaticLine
func (processor constantsAssignmentProcessor) Process(galacticLine string, w io.Writer) error {
	out := bufio.NewWriter(w)
	matches := processor.rx.FindStringSubmatch(galacticLine)
	if len(matches) != 3 {
		return writeTo(out, fmt.Sprintf("%s is not a valid constants assignment line", galacticLine))
	}
	// if input validation fails, then we should write it to the output
	if err := processor.validate(matches[2]); err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	constantsAssignments[strings.ToLower(matches[1])] = matches[2]
	return nil
}
