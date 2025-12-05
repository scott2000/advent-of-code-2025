package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	position := 50
	zeros := 0
	for line := range strings.Lines(string(contents)) {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parsed, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch line[0] {
		case 'L':
			position += parsed
		case 'R':
			position -= parsed
		default:
			panic("invalid line: " + line)
		}
		position %= 100
		if position < 0 {
			position += 100
		}
		if position == 0 {
			zeros++
		}
	}
	fmt.Println(zeros)
}
