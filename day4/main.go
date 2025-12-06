package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	contents, err := os.ReadFile("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	var grid Grid
	for line := range strings.Lines(string(contents)) {
		line = strings.TrimSpace(line)
		row := make([]bool, len(line))
		for i, ch := range line {
			row[i] = ch == '@'
		}
		grid = append(grid, row)
	}

	accessible := 0
	for row, r := range grid {
		for col, c := range r {
			if c && grid.countAdjacent(row, col) < 4 {
				accessible++
			}
		}
	}

	removedTotal := 0
	for {
		removed := 0
		for row, r := range grid {
			for col, c := range r {
				if c && grid.countAdjacent(row, col) < 4 {
					removed++
					r[col] = false
				}
			}
		}

		if removed == 0 {
			break
		}
		removedTotal += removed
	}

	fmt.Println(accessible)
	fmt.Println(removedTotal)
}

type Grid [][]bool

func (grid Grid) get(row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}
	r := grid[row]
	if col < 0 || col >= len(r) {
		return false
	}
	return r[col]
}

func (grid Grid) countAdjacent(row, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r == row && c == col {
				continue
			}
			if grid.get(r, c) {
				count++
			}
		}
	}
	return count
}
