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

	totalWithTwo := 0
	totalWithTwelve := 0
	for line := range strings.Lines(string(contents)) {
		line = strings.TrimSpace(line)
		totalWithTwo += findJoltage(line, 2)
		totalWithTwelve += findJoltage(line, 12)
	}
	fmt.Println(totalWithTwo)
	fmt.Println(totalWithTwelve)
}

func findJoltage(line string, numDigits int) int {
	firstIndexOfEachDigit := make([]int, 10)
	for i := range firstIndexOfEachDigit {
		firstIndexOfEachDigit[i] = -1
	}
	for i, ch := range line {
		digit := int(ch - '0')
		if firstIndexOfEachDigit[digit] == -1 {
			firstIndexOfEachDigit[digit] = i
		}
	}

	findIndexInRange := func(digit int, startIndex int) int {
		cachedIndex := firstIndexOfEachDigit[digit]
		if cachedIndex == -1 || cachedIndex >= startIndex {
			return cachedIndex
		}

		for i := startIndex; i < len(line); i++ {
			newDigit := int(line[i] - '0')
			if newDigit == digit {
				firstIndexOfEachDigit[digit] = i
				return i
			}
		}
		firstIndexOfEachDigit[digit] = -1
		return -1
	}

	accumulator := 0
	startIndex := 0
outer:
	for i := range numDigits {
		for digit := 9; digit >= 0; digit-- {
			index := findIndexInRange(digit, startIndex)
			if index == -1 || index >= len(line)+i+1-numDigits {
				continue
			}

			accumulator *= 10
			accumulator += digit
			startIndex = index + 1

			continue outer
		}
		panic("No digits found")
	}

	return accumulator
}
