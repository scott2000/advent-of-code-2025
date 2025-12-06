package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("inputs/day5.txt")
	if err != nil {
		panic(err)
	}

	splitContents := strings.Split(string(contents), "\n\n")
	rangesStr := splitContents[0]
	queriesStr := splitContents[1]

	var operations []Operation
	for line := range strings.Lines(rangesStr) {
		line = strings.TrimSpace(line)
		splitLine := strings.Split(line, "-")
		start, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(splitLine[1])
		if err != nil {
			panic(err)
		}
		operations = append(operations,
			Operation{position: start, isEnd: false},
			Operation{position: end + 1, isEnd: true})
	}

	slices.SortFunc(operations, func(a, b Operation) int {
		result := cmp.Compare(a.position, b.position)
		if result != 0 {
			return result
		}
		if a.isEnd == b.isEnd {
			return 0
		}
		if a.isEnd {
			return 1
		}
		return -1
	})

	var flattened []Operation
	depth := 0
	for _, operation := range operations {
		if !operation.isEnd {
			if depth == 0 {
				flattened = append(flattened,
					Operation{position: operation.position, isEnd: false})
			}
			depth++
			continue
		}

		depth--
		if depth == 0 {
			flattened = append(flattened,
				Operation{position: operation.position, isEnd: true})
		}
	}

	queriedFreshCount := 0
	for queryStr := range strings.Lines(queriesStr) {
		queryStr = strings.TrimSpace(queryStr)
		query, err := strconv.Atoi(queryStr)
		if err != nil {
			panic(err)
		}

		index, found := slices.BinarySearchFunc(flattened, query, func(operation Operation, query int) int {
			return cmp.Compare(operation.position, query)
		})
		if found {
			if !flattened[index].isEnd {
				queriedFreshCount++
			}
		} else {
			if index > 0 && !flattened[index-1].isEnd {
				queriedFreshCount++
			}
		}
	}

	totalFreshCount := 0
	startIndex := 0
	for _, operation := range flattened {
		if operation.isEnd {
			totalFreshCount += operation.position - startIndex
		} else {
			startIndex = operation.position
		}
	}

	fmt.Println(queriedFreshCount)
	fmt.Println(totalFreshCount)
}

type Operation struct {
	position int
	isEnd    bool
}
