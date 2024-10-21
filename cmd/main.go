package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/caseyleask/rem-e/consoleinteraction"
	"github.com/caseyleask/rem-e/robot"
)

func readFile(file string) []string {
	dat, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(dat), "\n")
}

func runFromFile(file string) {
	println("Read from file:", file)

	lines := readFile(file)
	commands := robot.ConvertLinesToCommands(lines)

	for _, command := range commands {
		println(fmt.Sprintf("%#v", command))
	}
}

// Parse external input
func Parse(args []string) {
	switch len(args) {

	case 1:
		consoleinteraction.Begin()

	case 2:
		runFromFile(args[1])

	default:
		panic("Program supplied with too many arguments")

	}
}
