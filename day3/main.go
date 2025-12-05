package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	total := 0
	for line := range strings.Lines(string(contents)) {
		line = strings.TrimSpace(line)
		total += findJoltage(line)
	}
	fmt.Println(total)
}

func findJoltage(line string) int {
	highestIndex := 0
	highestValue := '0'
	for i, ch := range line[:len(line)-1] {
		if ch > highestValue {
			highestIndex, highestValue = i, ch
		}
	}

	secondHighestValue := '0'
	for _, ch := range line[highestIndex+1:] {
		if ch > secondHighestValue {
			secondHighestValue = ch
		}
	}
	return int(highestValue-'0')*10 + int(secondHighestValue-'0')
}
