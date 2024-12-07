package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	lines, _ := readLines("input.txt")

	var grid [][]string
	var startposition string
	var currentDirection string

	for i, line := range lines {
		var slicedString []string

		if strings.Contains(line, "v") {
			index := strings.IndexAny(line, "v")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Down
		}

		if strings.Contains(line, "^") {
			index := strings.IndexAny(line, "^")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Up
		}

		if strings.Contains(line, "<") {
			index := strings.IndexAny(line, "<")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Left
		}

		if strings.Contains(line, ">") {
			index := strings.IndexAny(line, ">")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Right
		}

		slicedString = append(slicedString, strings.Split(line, "")...)

		grid = append(grid, slicedString)
	}

	println("StartPosition:", startposition)
	println("Startdirection:", currentDirection)
	for _, row := range grid {
		for column := range row {
			print(row[column])
		}
		println()
	}
	numberOfVisitedNodes := TraverseMap(grid, startposition, currentDirection)

	println("Number of visited nodes: ", numberOfVisitedNodes)
}

func partTwo() {
	lines, _ := readLines("input.txt")

	var grid [][]string
	var startposition string
	var currentDirection string

	for i, line := range lines {
		var slicedString []string

		if strings.Contains(line, "v") {
			index := strings.IndexAny(line, "v")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Down
		}

		if strings.Contains(line, "^") {
			index := strings.IndexAny(line, "^")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Up
		}

		if strings.Contains(line, "<") {
			index := strings.IndexAny(line, "<")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Left
		}

		if strings.Contains(line, ">") {
			index := strings.IndexAny(line, ">")
			startposition = strconv.Itoa(i) + "," + strconv.Itoa(index)
			currentDirection = Right
		}

		slicedString = append(slicedString, strings.Split(line, "")...)

		grid = append(grid, slicedString)
	}

	count := 0
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == "." {
				grid[i][j] = "0"

				if hasLoop(grid, startposition, currentDirection) {
					count++
				}

				grid[i][j] = "."
			}
		}
	}

	println("Number of possible loop points: ", count)
}

func hasLoop(levelMap [][]string, currentposition string, currentdirection string) bool {
	numberOfTurns := 0

	row, _ := strconv.Atoi(strings.Split(currentposition, ",")[0])
	column, _ := strconv.Atoi(strings.Split(currentposition, ",")[1])

	for ok := true; ok; {
		if numberOfTurns > 200 {
			return true
		}

		switch currentdirection {
		case Up:
			row--
		case Down:
			row++
		case Left:
			column--
		case Right:
			column++
		}

		if row < 0 || row >= len(levelMap[0]) || column < 0 || column >= len(levelMap) {
			break
		}

		if levelMap[row][column] == "#" || levelMap[row][column] == "0" {
			switch currentdirection {
			case Up:
				row++
			case Down:
				row--
			case Left:
				column++
			case Right:
				column--
			}

			currentdirection = rotateRight90Degrees(currentdirection)
			numberOfTurns++
		}
	}

	return false
}

func TraverseMap(levelMap [][]string, currentposition string, currentdirection string) int {
	row, _ := strconv.Atoi(strings.Split(currentposition, ",")[0])
	column, _ := strconv.Atoi(strings.Split(currentposition, ",")[1])

	//start at 1, the start node counts as well
	visitCount := 1

	for ok := true; ok; {
		switch currentdirection {
		case Up:
			row--
		case Down:
			row++
		case Left:
			column--
		case Right:
			column++
		}

		if row < 0 || row >= len(levelMap[0]) || column < 0 || column >= len(levelMap) {
			break
		}

		if levelMap[row][column] == "." {
			levelMap[row][column] = "X"
			visitCount++
		}

		if levelMap[row][column] == "#" {
			switch currentdirection {
			case Up:
				row++
			case Down:
				row--
			case Left:
				column++
			case Right:
				column--
			}

			currentdirection = rotateRight90Degrees(currentdirection)
		}
	}

	println()
	println()
	for _, row := range levelMap {
		for column := range row {
			print(row[column])
		}
		println()
	}

	return visitCount
}

func rotateRight90Degrees(currentdirection string) string {
	switch currentdirection {
	case Up:
		return Right
	case Down:
		return Left
	case Left:
		return Up
	case Right:
		return Down
	}

	return currentdirection
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

// Declaring constants for representing directions
const (
	Up    = "^"
	Down  = "v"
	Left  = "<"
	Right = ">"
)
