package cmd

import (
	"reflect"
	"testing"
)

func TestParsing(t *testing.T) {
	type test struct {
		input    string
		expected []string
	}

	tests := []test{
		{
			input: "../testdata/test1.dat",
			expected: []string{
				"Output: 1,1,NORTH",
				"Output: 1,2,NORTH",
			},
		},
		{
			input: "../testdata/test2.dat",
			expected: []string{
				"Output: 0,0,NORTH",
				"Output: 0,4,WEST",
				"Output: 4,4,EAST",
				"Output: 4,0,SOUTH",
				"Output: 0,0,NORTH",
			},
		},
	}

	for _, test := range tests {
		actual := Read(readFile(test.input))
		if !reflect.DeepEqual(test.expected, actual) {
			t.Fatalf("input: %#v, expected: %#v, got: %#v", test.input, test.expected, actual)
		}
	}
}
