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

	invalidSumWithTwo := 0
	invalidSumWithAny := 0
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
		for n := start; n <= end; n++ {
			invalid := getInvalidIdCount(n)
			if invalid == 2 {
				invalidSumWithTwo += n
			}
			if invalid != 0 {
				invalidSumWithAny += n
			}
		}
	}
	fmt.Println(invalidSumWithTwo)
	fmt.Println(invalidSumWithAny)
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getInvalidIdCount(id int) int {
	numDigits := numDigits(id)
outer:
	for factor := 2; factor <= numDigits; factor++ {
		if numDigits%factor != 0 {
			continue
		}

		chunkSize := numDigits / factor
		for offset := range chunkSize {
			digit := (id / pow10(offset)) % 10
			for i := 1; i < factor; i++ {
				otherDigit := (id / pow10(i*chunkSize+offset)) % 10
				if otherDigit != digit {
					continue outer
				}
			}
		}

		return factor
	}

	return 0
}
