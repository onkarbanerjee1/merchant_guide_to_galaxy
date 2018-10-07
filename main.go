package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/patterns"
	procs "github.com/onkarbanerjee1/merchant_guide_to_galaxy/processors"
	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/roman"
)

// pats holds the different patterns for diff kind of statements in input file
var pats = patterns.Patterns{
	patterns.Pattern{Type: patterns.ConstantAssignment, PatternRx: regexp.MustCompile("^([A-Za-z]+) is ([I|V|X|L|C|D|M])$")},
	patterns.Pattern{Type: patterns.MetalCostAssignment, PatternRx: regexp.MustCompile("^([A-Za-z\\s]+) ([A-Za-z]+) is ([0-9]*[.]?[0-9]+) [c|C]redits$")},
	patterns.Pattern{Type: patterns.GalacticAmountQuestion, PatternRx: regexp.MustCompile("^how much is ([A-Za-z\\s]+)\\?$")},
	patterns.Pattern{Type: patterns.MetalCostQuestion, PatternRx: regexp.MustCompile("^how many [c|C]redits is ([A-Za-z\\s]+) ([A-Za-z\\s]+)\\?$")},
}

// when we have no idea about the input we receive
var errNoIdea = errors.New("I have no idea what you are talking about")

// exit the program after printing out the error causing it
func abort(err error) {
	fmt.Println(err)
	os.Exit(1)

}

// init all types of processors
func initProcessors(processors *map[patterns.PatternType]procs.Processor) {
	(*processors)[patterns.ConstantAssignment] = procs.NewConstantsAssignmentProcessor(pats[patterns.ConstantAssignment].PatternRx, roman.ValidateSymbol)
	(*processors)[patterns.MetalCostAssignment] = procs.NewMetalCostAssignmentProcessor(pats[patterns.MetalCostAssignment].PatternRx, roman.ToArabic)
	(*processors)[patterns.GalacticAmountQuestion] = procs.NewGalacticAmountQuestionProcessor(pats[patterns.GalacticAmountQuestion].PatternRx, roman.ToArabic)
	(*processors)[patterns.MetalCostQuestion] = procs.NewMetalCostQuestionProcessor(pats[patterns.MetalCostQuestion].PatternRx, roman.ToArabic)
}

func main() {
	// read the input from in (which could be anything like a file, the stdin or any stream of data)
	var in io.Reader
	// write the output to out (which could be anything like a file, the stdout or any stream of data)
	var out io.Writer

	if len(os.Args) != 2 {
		abort(fmt.Errorf("Please provide the input file path as a single argument only"))
	}

	// In this client(main program) we choose to use a file to read the inputs from and writes output to stdout
	// we can update in and out to any io.Reader and io.Writer if source and dest changes
	inputFilePath := os.Args[1]
	file, err := os.Open(inputFilePath)
	if err != nil {
		abort(fmt.Errorf("Could not open %s, got %s", inputFilePath, err))
	}
	defer file.Close()
	in, out = file, os.Stdout

	// get an instance of all the diff processors we have
	processors := &map[patterns.PatternType]procs.Processor{}
	initProcessors(processors)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		typ := pats.GetTypeOf(line)
		if typ == patterns.InvalidLine {
			fmt.Println(errNoIdea)
			continue
		}
		// process the line using a processor based on the line's typ
		(*processors)[typ].Process(line, out)
	}

	if err := scanner.Err(); err != nil {
		abort(fmt.Errorf("Could not read from %s, got %s", inputFilePath, err))
	}
}
