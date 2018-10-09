package processors

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// galacticAmountQuestionProcessor is a Processor that can Process question statements about galactic
// amounts , for example - lines like "how much is pish tegj glob glob ?"
type galacticAmountQuestionProcessor struct {
	rx       *regexp.Regexp // to match the subsstrings in a galaticAmountQuestion line
	toArabic fnArabic       // We will need a function to convert roman number to arabic
}

// NewGalacticAmountQuestionProcessor returns a Processor that will process galacticAmountQuestions
func NewGalacticAmountQuestionProcessor(rx *regexp.Regexp, toArabic fnArabic) Processor {
	return galacticAmountQuestionProcessor{rx: rx, toArabic: toArabic}
}

// galacticAmountQuestionProcessor's implementation of a Processor
func (processor galacticAmountQuestionProcessor) Process(galacticLine string, w io.Writer) error {
	out := bufio.NewWriter(w)
	matches := processor.rx.FindStringSubmatch(galacticLine)
	if len(matches) != 2 {
		return writeTo(out, fmt.Sprintf("%s is not a valid galactic amount question line", galacticLine))
	}

	// if sufficient info not present for getting a roman representaion of the provided units,
	// (example - not enough constants assigned in input) then write it to the output
	romanNumber, err := getRomanNumber(matches[1])
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	arabic, err := processor.toArabic(romanNumber)
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	return writeTo(out, fmt.Sprintf("%s is %d", strings.TrimSpace(matches[1]), arabic))
}
