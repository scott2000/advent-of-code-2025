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
		wasZero := position == 0
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
		if position <= 0 && !wasZero {
			zeros += 1
		}
		zeros += Abs(position) / 100
		position %= 100
		if position < 0 {
			position += 100
		}
	}
	fmt.Println(position)
	fmt.Println(zeros)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
