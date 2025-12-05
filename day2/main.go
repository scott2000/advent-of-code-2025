package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("inputs/day2.txt")
	if err != nil {
		panic(err)
	}

	invalidSum := 0
	for r := range strings.SplitSeq(strings.TrimSpace(string(contents)), ",") {
		sides := strings.Split(r, "-")
		start, err := strconv.Atoi(sides[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(sides[1])
		if err != nil {
			panic(err)
		}
		for i := nextInvalidId(start - 1); i <= end; i = nextInvalidId(i) {
			invalidSum += i
		}
	}
	fmt.Println(invalidSum)
}

func nextInvalidId(id int) int {
	if id < 10 {
		return 11
	}
	numDigits := numDigits(id)
	currentPowerOfTen := pow10(numDigits / 2)
	multiplier := currentPowerOfTen + 1
	maxFactor := multiplier - 2

	factor := id / multiplier
	var nextId int
	if factor < maxFactor {
		nextId = (factor + 1) * multiplier
	} else {
		nextPowerOfTen := currentPowerOfTen * 10
		nextMultiplier := nextPowerOfTen + 1
		nextId = currentPowerOfTen * nextMultiplier
	}

	if nextId <= id {
		panic(fmt.Errorf("invalid: nextInvalidId(%v) = %v", id, nextId))
	}
	return nextId
}

func numDigits(id int) int {
	digits := 1
	for id >= 10 {
		id /= 10
		digits++
	}
	return digits
}

func pow10(p int) int {
	result := 1
	for range p {
		result *= 10
	}
	return result
}
