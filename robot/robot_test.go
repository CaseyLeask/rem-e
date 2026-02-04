package robot

import (
	"reflect"
	"testing"
)

type OptionalString string

func (s OptionalString) toPointer() *string {
	tmp := string(s)
	return &tmp
}

func TestRobot(t *testing.T) {
	type robotInput struct {
		robot   Robot
		command Command
	}
	type robotResult struct {
		robot  Robot
		output *string
	}
	type test struct {
		input    robotInput
		expected robotResult
	}

	lines := []test{
		{
			input: robotInput{
				robot:   NewRobot(),
				command: Place{X: 0, Y: 0, Direction: North{}},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Place{X: 0, Y: 0, Direction: nil},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Place{X: 5, Y: 0, Direction: North{}},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Place{X: 0, Y: 5, Direction: North{}},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Place{X: -1, Y: 0, Direction: North{}},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Place{X: 0, Y: -1, Direction: North{}},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 1, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 4, facing: North{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 4, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: East{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 1, Y: 0, facing: East{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 4, Y: 0, facing: East{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 4, Y: 0, facing: East{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: South{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: South{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 4, facing: South{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 3, facing: South{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: West{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: West{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 4, Y: 0, facing: West{}},
				command: Move{},
			},
			expected: robotResult{
				robot:  Robot{X: 3, Y: 0, facing: West{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Left{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: West{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Right{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: East{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: East{}},
				command: Left{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: East{}},
				command: Right{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: South{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: South{}},
				command: Left{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: East{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: South{}},
				command: Right{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: West{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: West{}},
				command: Left{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: South{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: West{}},
				command: Right{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   NewRobot(),
				command: Report{},
			},
			expected: robotResult{
				robot:  NewRobot(),
				output: nil,
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 0, facing: North{}},
				command: Report{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 0, facing: North{}},
				output: OptionalString("Output: 0,0,NORTH").toPointer(),
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 0, Y: 1, facing: East{}},
				command: Report{},
			},
			expected: robotResult{
				robot:  Robot{X: 0, Y: 1, facing: East{}},
				output: OptionalString("Output: 0,1,EAST").toPointer(),
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 1, Y: 0, facing: South{}},
				command: Report{},
			},
			expected: robotResult{
				robot:  Robot{X: 1, Y: 0, facing: South{}},
				output: OptionalString("Output: 1,0,SOUTH").toPointer(),
			},
		},
		{
			input: robotInput{
				robot:   Robot{X: 4, Y: 4, facing: West{}},
				command: Report{},
			},
			expected: robotResult{
				robot:  Robot{X: 4, Y: 4, facing: West{}},
				output: OptionalString("Output: 4,4,WEST").toPointer(),
			},
		},
	}

	for _, line := range lines {
		newRobot, output := line.input.robot.Run(line.input.command)
		if !reflect.DeepEqual(line.expected.robot, newRobot) {
			t.Fatalf("expected: %#v, got: %#v", line.expected.robot, newRobot)
		}

		if line.expected.output == nil {
			if output != nil {
				t.Fatalf("expected nil, got: %#v", output)
			}
		} else {
			if !reflect.DeepEqual(*line.expected.output, *output) {
				t.Fatalf("expected: %#v, got: %#v", *line.expected.output, *output)
			}
		}
	}
}
