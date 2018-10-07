package processors

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// metalCostAssignmentProcessor is a Processor that can Process metal cost questions
// for example - galactic lines "how many Credits is glob prok Silver ?".
type metalCostQuestionProcessor struct {
	rx       *regexp.Regexp
	toArabic fnArabic
}

// NewMetalCostQuestionProcessor returns a Processor that will process metalCostQuestions
func NewMetalCostQuestionProcessor(rx *regexp.Regexp, toArabic fnArabic) Processor {
	return metalCostQuestionProcessor{rx: rx, toArabic: toArabic}
}

// metalCostQuestionProcessor's implementation of how to Process a galaticLine
func (processor metalCostQuestionProcessor) Process(galacticLine string, w io.Writer) error {
	out := bufio.NewWriter(w)
	matches := processor.rx.FindStringSubmatch(galacticLine)
	if len(matches) != 3 {
		return writeTo(out, fmt.Sprintf("%s is not a metal cost question line", galacticLine))
	}

	// if sufficient info not present for getting a roman representaion of the provided units,
	// (example - not enough constants assigned in input) then write it to the output
	romanNumber, err := getRomanNumber(matches[1])
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	// get the arabic representation to calculate the cost for the metal in the question
	arabic, err := processor.toArabic(romanNumber)
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}
	units := float64(arabic)

	// if the metal in question is not assigned any credits for, then we write it to the output
	metal := strings.TrimSpace(matches[2])
	creditsPerUnit, ok := metalRates[strings.ToLower(metal)]
	if !ok {
		return writeTo(out, fmt.Sprintf("%s , %s", errNoCredits, metal))
	}

	return writeTo(out, fmt.Sprintf("%s %s is %g Credits", matches[1], metal, creditsPerUnit*units))
}
