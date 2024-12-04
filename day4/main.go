package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	numberOfXmasses := 0

	numberOfXmasses += countHorizontal(lines)
	numberOfXmasses += countVertical(lines)
	numberOfXmasses += countDiagonally(lines, "XMAS")
	numberOfXmasses += countDiagonally(lines, "SAMX")

	println(numberOfXmasses)
}

func partTwo() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	numberOfXmasses := 0

	letterAindexes := countCross(lines, "MAS")
	reverseIndexes := countCross(lines, "SAM")

	letterAindexes = append(letterAindexes, reverseIndexes...)

	sort.Strings(letterAindexes)
	copySlice := make([]string, len(letterAindexes))
	copy(copySlice, letterAindexes)

	uniqueIndexes := slices.Compact(copySlice)

	count := 0
	for _, uniqueIndex := range uniqueIndexes {
		count += countValues(letterAindexes, uniqueIndex)

		if count == 2 {
			numberOfXmasses++
		}
		count = 0
	}

	println(numberOfXmasses)
}

func countValues(slice []string, value string) int {
	var count int = 0

	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			count++
		}
	}

	return count
}

func countHorizontal(lines []string) int {
	numberOfXmasses := 0
	for _, line := range lines {
		numberOfXmasses += strings.Count(line, "XMAS")
		numberOfXmasses += strings.Count(line, "SAMX")
	}

	return numberOfXmasses
}

func countVertical(lines []string) int {
	numberOfXmasses := 0
	lineLength := len(lines[0])

	for i := 0; i < lineLength; i++ {
		var verticalLine string
		for _, line := range lines {
			verticalLine += string(line[i])
		}

		numberOfXmasses += strings.Count(verticalLine, "XMAS")
		numberOfXmasses += strings.Count(verticalLine, "SAMX")
	}

	return numberOfXmasses
}

func countDiagonally(lines []string, word string) int {
	numberOfXmasses := 0

	var grid [][]string

	for _, line := range lines {
		var slicedString []string
		for j, _ := range line {
			slicedString = append(slicedString, string(line[j]))
		}

		grid = append(grid, slicedString)
	}

	rows := len(grid)
	cols := len(grid[0])
	wordLen := 4

	// Check diagonals from top-left to bottom-right
	for r := 0; r <= rows-wordLen; r++ {
		for c := 0; c <= cols-wordLen; c++ {
			match := true
			for i := 0; i < wordLen; i++ {
				if grid[r+i][c+i] != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				numberOfXmasses++
			}
		}
	}

	// Check diagonals from top-right to bottom-left
	for r := 0; r <= rows-wordLen; r++ {
		for c := wordLen - 1; c < cols; c++ {
			match := true
			for i := 0; i < wordLen; i++ {
				if grid[r+i][c-i] != string(word[i]) {
					match = false
					break
				}
			}
			if match {
				numberOfXmasses++
			}
		}
	}

	return numberOfXmasses
}

func countCross(lines []string, word string) []string {
	var letterAindexes []string
	var grid [][]string

	for _, line := range lines {
		var slicedString []string
		for j, _ := range line {
			slicedString = append(slicedString, string(line[j]))
		}

		grid = append(grid, slicedString)
	}

	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)

	// Check diagonals from top-left to bottom-right
	var aIndex string
	for r := 0; r <= rows-wordLen; r++ {
		for c := 0; c <= cols-wordLen; c++ {
			match := true
			for i := 0; i < wordLen; i++ {
				if grid[r+i][c+i] != string(word[i]) {
					match = false
					break
				}

				if grid[r+i][c+i] == "A" {
					aIndex = strconv.Itoa(r+i) + "|" + strconv.Itoa(c+i)
				}
			}
			if match {
				letterAindexes = append(letterAindexes, aIndex)
			}
		}
	}

	// Check diagonals from top-right to bottom-left
	for r := 0; r <= rows-wordLen; r++ {
		for c := wordLen - 1; c < cols; c++ {
			match := true
			for i := 0; i < wordLen; i++ {
				if grid[r+i][c-i] != string(word[i]) {
					match = false
					break
				}

				if grid[r+i][c-i] == "A" {
					aIndex = strconv.Itoa(r+i) + "|" + strconv.Itoa(c-i)
				}
			}
			if match {
				letterAindexes = append(letterAindexes, aIndex)
			}
		}
	}

	return letterAindexes
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
