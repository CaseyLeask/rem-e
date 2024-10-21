package robot

import (
	"strconv"
	"strings"
)

type Command interface {
	isCommand()
}

type Place struct {
	X         int
	Y         int
	Direction Direction
}

func (Place) isCommand() {}

type Move struct{}

func (Move) isCommand() {}

type Left struct{}

func (Left) isCommand() {}

type Right struct{}

func (Right) isCommand() {}

type Report struct{}

func (Report) isCommand() {}

func ConvertLinesToCommands(lines []string) []Command {
	commands := []Command{}

	for _, line := range lines {
		command := ConvertLineToCommand(line)
		commands = append(commands, command)
	}

	return commands
}

func ConvertLineToCommand(line string) Command {
	words := strings.Split(line, " ")
	switch words[0] {
	case "PLACE":
		placeArguments := strings.Split(words[1], ",")

		if len(placeArguments) != 3 {
			return nil
		}

		x, err_x := strconv.Atoi(placeArguments[0])
		y, err_y := strconv.Atoi(placeArguments[1])
		if err_x != nil || err_y != nil {
			return nil
		}

		var direction Direction
		switch placeArguments[2] {
		case "NORTH":
			direction = North{}
		case "EAST":
			direction = East{}
		case "SOUTH":
			direction = South{}
		case "WEST":
			direction = West{}
		default:
			return nil
		}
		return Place{
			X:         x,
			Y:         y,
			Direction: direction,
		}
	case "MOVE":
		return Move{}
	case "LEFT":
		return Left{}
	case "RIGHT":
		return Right{}
	case "REPORT":
		return Report{}
	default:
		return nil
	}
}

func InitialState() {

}
