package processors

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// metalCostAssignmentProcessor is a Processor that can Process metal cost assignments galactic
// lines , for example - lines like "glob glob Silver is 34 Credits".
type metalCostAssignmentProcessor struct {
	rx       *regexp.Regexp // to match the substrings in a metalCostAssignment line
	toArabic fnArabic       // We will need a function to convert roman number to arabic
}

// NewMetalCostAssignmentProcessor returns a Processor that will process metalCostAssignments
func NewMetalCostAssignmentProcessor(rx *regexp.Regexp, toArabic fnArabic) Processor {
	return metalCostAssignmentProcessor{rx: rx, toArabic: toArabic}
}

// metalCostAssignmentProcessor's implementation of how to Process a galaticLine
func (processor metalCostAssignmentProcessor) Process(galacticLine string, w io.Writer) error {
	out := bufio.NewWriter(w)
	matches := processor.rx.FindStringSubmatch(galacticLine)
	if len(matches) != 4 {
		return writeTo(out, fmt.Sprintf("%s is not a valid metal cost assignment line", galacticLine))
	}

	// units denotes the amount of metal whose cost(ie.cumulativeValue) is assigned in the galacticLine,
	units, metal := matches[1], matches[2]
	cumulativeValue, err := strconv.ParseFloat(matches[3], 64)
	if err != nil {
		return fmt.Errorf("Could not parse %s to float, got %s", matches[3], err)
	}

	// if sufficient info not present for getting a roman representaion of the provided units,
	// (example - not enough constants assigned in input) then write it to the output
	romanNumber, err := getRomanNumber(units)
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	// get the arabic representation to calculate the cost for the metal
	arabic, err := processor.toArabic(romanNumber)
	if err != nil {
		return writeTo(out, fmt.Sprint(err))
	}

	credits := cumulativeValue / float64(arabic)
	metalRates[strings.ToLower(metal)] = credits
	return nil
}
