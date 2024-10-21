package robot

import (
	"reflect"
	"testing"
)

func TestParsing(t *testing.T) {
	type test struct {
		input    string
		expected Command
	}

	lines := []test{
		{input: "PLACE 0,0,NORTH", expected: Place{X: 0, Y: 0, Direction: North{}}},
		{input: "PLACE 0,0, NORTH", expected: nil},
		{input: "PLACE 1,1,EAST", expected: Place{X: 1, Y: 1, Direction: East{}}},
		{input: "PLACE -1,1,SOUTH", expected: Place{X: -1, Y: 1, Direction: South{}}},
		{input: "PLACE -1,1,WEST", expected: Place{X: -1, Y: 1, Direction: West{}}},
		{input: "PLACE 1.1,1", expected: nil},
		{input: "PLACE ,1", expected: nil},
		{input: "PLACE 1, 1", expected: nil},
		{input: "MOVE", expected: Move{}},
		{input: "LEFT", expected: Left{}},
		{input: "RIGHT", expected: Right{}},
		{input: "REPORT", expected: Report{}},
		{input: "FAKE", expected: nil},
	}

	var inputs []string
	var commands []Command
	for _, line := range lines {
		actual := ConvertLineToCommand(line.input)
		if !reflect.DeepEqual(line.expected, actual) {
			t.Fatalf("expected: %#v, got: %#v", line.expected, actual)
		}
		inputs = append(inputs, line.input)
		commands = append(commands, actual)
	}

	if !reflect.DeepEqual(commands, ConvertLinesToCommands(inputs)) {
		t.Fatalf("expected: %#v, got: %#v", commands, ConvertLinesToCommands(inputs))
	}
}
