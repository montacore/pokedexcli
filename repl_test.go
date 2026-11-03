package main

import (
	"testing"
)


func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input  string
		expected  []string
	}{
		"empty string": {
			input: "",
			expected: []string{},
		},
		"single word": {
			input:	"Hello",
			expected: []string{"hello"},
		},
		"multiple words with extra spaces": {
			input:	"   Hello   world     ",
			expected: []string{"hello", "world"},
		},
		"mixed case words": {
			input: "Golang Is fun",
			expected: []string{"golang","is", "fun"},
		},
		"tabs and newlines": {
			input: "line1\tline2\nline3",
			expected: []string{"line1", "line2", "line3"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		//TODO: Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case.
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Errorf("Expected length %d, but got %d for input: %q", len(c.expected), len(actual), c.input)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			//TODO: Check each word in the slice
			// if they don't match, use t.Errorf to print an error message and
			// fail the test
			if word != expectedWord {
				t.Errorf("Expected word %q, but got %q for input: %q", expectedWord, word, c.input)
			}
		}
	}
}
