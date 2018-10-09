package processors

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// maps to hold the assignments
var (
	constantsAssignments = map[string]string{}
	metalRates           = map[string]float64{}
)

// different errors returned by this package
var (
	errNoConstant = errors.New("No roman number symbol assigned for this galactic constant")
	errNoCredits  = errors.New("Found no credits assigned for this metal")
)

// fnValidate is a function that can be used to validate input if reqd.
type fnValidate func(string) error

// fnArabic is a function that can be used to convert a string to arabic number when reqd.
type fnArabic func(string) (int, error)

// Processor guarantees that any implementation of it can Process galacticLine and write to
// the supplied io.Writer
type Processor interface {
	Process(galacticLine string, w io.Writer) error
}

// getRomanNumber takes a galactic number and returns roman representation of it.
// It doesn't validate the returned roman representation
func getRomanNumber(line string) (string, error) {
	if line == "" {
		return "0", nil
	}
	var romanNumberSb strings.Builder
	keys := strings.Fields(line)
	for _, each := range keys {
		roman, ok := constantsAssignments[strings.ToLower(each)]
		if !ok {
			return "", fmt.Errorf("%s for %s", errNoConstant, each)
		}
		romanNumberSb.WriteString(roman)
	}

	return romanNumberSb.String(), nil
}

// just a convenience function to write output to the supplied *bufio.Writer
func writeTo(out *bufio.Writer, response string) error {
	if _, err := out.Write(append([]byte(response), byte('\n'))); err != nil {
		return err
	}
	return out.Flush()
}

// ClearAssignments clears all assignments made
func ClearAssignments() {
	constantsAssignments = map[string]string{}
	metalRates = map[string]float64{}
}
