package cmd

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/caseyleask/rem-e/robot"
)

func readFile(file string) []string {
	dat, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func Read(lines []string) []string {
	commands := robot.ConvertLinesToCommands(lines)

	walle := robot.NewRobot()
	var result []string

	for _, command := range commands {
		var output *string

		walle, output = walle.Run(command)
		if output != nil {
			result = append(result, *output)
		}
	}

	return result
}

func Scan(stdin io.Reader) {
	// Create a new scanner that reads from standard input (os.Stdin)
	scanner := bufio.NewScanner(stdin)
	walle := robot.NewRobot()

	for {
		var output *string

		// Scan() blocks until a line is available (user presses Enter)
		scanner.Scan()

		// Get the text that was scanned
		input := scanner.Text()

		command := robot.ConvertLineToCommand(input)
		walle, output = walle.Run(command)

		if output != nil {
			println(output)
		}
	}

}

// Parse external input
func Parse(args []string) {
	switch len(args) {

	case 1:
		Scan(os.Stdin)

	case 2:
		println(Read(readFile(args[1])))

	default:
		panic("Program supplied with too many arguments")

	}
}
