package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/onkarbanerjee1/merchant_guide_to_galaxy/processors"
)

func TestIntegrationMain(t *testing.T) {
	testCases := []struct {
		inputFolder string
	}{
		{"input_ok"},
		{"input_floating_credits"},
		{"input_invalid_const_in_metal_cost_assignment"},
		{"input_invalid_line_in_beginning"},
		{"input_invalid_line_in_middle"},
		{"input_invalid_metal_in_question"},
		{"input_invalid_roman_number"},
		{"input_invalid_roman_number_in_question"},
		{"input_mixed_case"},
		{"input_no_blank"},
		{"input_question_before_assignment"},
		{"invalid_roman_symbol"},
	}

	for _, tc := range testCases {
		act := getMainOutput(filepath.Join("test", tc.inputFolder, "input.txt"))
		b, err := ioutil.ReadFile(filepath.Join("test", tc.inputFolder, "output.txt"))
		if err != nil {
			t.Fatal("Expected no error in reading output contents, got", err)
		}
		if act != string(b) {
			t.Fatal("got", act, "for", tc.inputFolder)
		}
		processors.ClearAssignments()
	}
}

func getMainOutput(inputFilePath string) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"", inputFilePath}

	main()

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		out <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	return <-out
}
