package main

import (
	"bytes"
	"fmt"
	c "github.com/nasccped/godvent-of-code-2025/internal/color"
	"github.com/nasccped/godvent-of-code-2025/solves/day01"
	"github.com/nasccped/godvent-of-code-2025/solves/day02"
	"github.com/nasccped/godvent-of-code-2025/solves/day03"
	"github.com/nasccped/godvent-of-code-2025/solves/day04"
	"github.com/nasccped/godvent-of-code-2025/solves/day05"
	"github.com/nasccped/godvent-of-code-2025/solves/day06"
	"os"
	"strconv"
	"strings"
)

// Mapping to store solves.
var mapping [][]func(string) int = [][]func(string) int{
	{day01.Part1, day01.Part2},
	{day02.Part1, day02.Part2},
	{day03.Part1, day03.Part2},
	{day04.Part1, day04.Part2},
	{day05.Part1, day05.Part2},
	{day06.Part1, day06.Part2},
}

// Remove leading whitespaces from a string input.
func removeLeading(input string) string {
	rows := strings.Split(input, "\n")
	if rows[0] == "" {
		rows = rows[1:]
	}
	if rows[len(rows)-1] == "" {
		rows = rows[:len(rows)-1]
	}
	return strings.Join(rows, "\n")
}

// Returns a `bytes.Buffer` with a error tag.
func bufErrorTag() bytes.Buffer {
	var buf bytes.Buffer
	buf.WriteString(c.Colored("error: ").Red())
	return buf
}

// Returns an error based on the provided args.
func argErrorMessage(args []string) error {
	buf := bufErrorTag()
	buf.WriteString(
		fmt.Sprintf(
			"the provided args aren't valid: [%s]\n\n",
			strings.Join(args, ", "),
		),
	)
	buf.WriteString("The expected program call is something like:\n")
	buf.WriteString(
		fmt.Sprintf(
			"  %s %s",
			c.Colored("go run ./main.go").Green(),
			c.Colored("<DAY NUMBER>").Cyan(),
		),
	)
	return fmt.Errorf("%s", buf.String())
}

// Error message to return when solve still not implemented.
func dayErrorMessage(num int) error {
	buf := bufErrorTag()
	buf.WriteString("no solve for the provided day (")
	buf.WriteString(strconv.Itoa(num))
	buf.WriteString(")\n\n")
	buf.WriteString("Not yet implemented maybe?")
	return fmt.Errorf("%s", buf.String())
}

// Returns an error based on the provided file path.
func fileErrorMessage(filePath string) error {
	buf := bufErrorTag()
	buf.WriteString(
		fmt.Sprintf(
			"couldn't read the `%s` file.\n\n",
			c.Colored(filePath).Cyan(),
		),
	)
	buf.WriteString("This happens due to:\n")
	buf.WriteString("  - file doesn't exists\n")
	buf.WriteString("  - not enough privileges\n")
	buf.WriteString("  - input file not added")
	return fmt.Errorf("%s", buf.String())
}

// Gets the solve function by a given index.
func getSolve(ind int, part1 bool) func(string) int {
	var f func(string) int
	if part1 {
		f = mapping[ind][0]
	} else {
		f = mapping[ind][1]
	}
	return f
}

func main() {
	var (
		num      int
		err      error
		filePath string
		asByte   []byte
		input    string
		f        func(string) int
	)
	args := os.Args[1:]
	errMessage := argErrorMessage(args)
	if len(args) != 1 {
		fmt.Println(errMessage)
		os.Exit(1)
	} else if num, err = strconv.Atoi(args[0]); err != nil {
		fmt.Println(errMessage)
		os.Exit(1)
	}
	filePath = fmt.Sprintf("inputs/day%02d.txt", num)
	errMessage = fileErrorMessage(filePath)
	if asByte, err = os.ReadFile(filePath); err != nil {
		fmt.Println(errMessage)
		os.Exit(1)
	}
	errMessage = dayErrorMessage(num)
	if num > len(mapping) {
		fmt.Println(errMessage)
		os.Exit(1)
	}
	input = removeLeading(string(asByte))
	f = getSolve(num-1, true)
	fmt.Printf(
		"Solve for %s: %d\n",
		c.Colored(fmt.Sprintf("day %02d (part 1)", num)).Green(),
		f(input),
	)
	f = getSolve(num-1, false)
	fmt.Printf(
		"Solve for %s: %d\n",
		c.Colored(fmt.Sprintf("day %02d (part 2)", num)).Green(),
		f(input),
	)
}
