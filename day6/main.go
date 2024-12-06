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

func printGrid(grid [][]string) {
	for _, row := range grid {
		for column := range row {
			print(row[column])
		}
		println()
	}
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

func TraverseMapToFindLoopPoints(levelMap [][]string, currentposition string, currentdirection string) int {
	row, _ := strconv.Atoi(strings.Split(currentposition, ",")[0])
	column, _ := strconv.Atoi(strings.Split(currentposition, ",")[1])

	firstPassMap := loop(levelMap, row, column, currentdirection)

	row, _ = strconv.Atoi(strings.Split(currentposition, ",")[0])
	column, _ = strconv.Atoi(strings.Split(currentposition, ",")[1])

	//do a second pass to see how many we've missed
	secondPassMap := loop(firstPassMap, row, column, currentdirection)

	loopPositionCount := 0
	for _, row := range secondPassMap {
		for _, column := range row {
			if column == "0" {
				loopPositionCount++
			}
		}
	}

	return loopPositionCount
}

func loop(levelMap [][]string, row int, column int, currentdirection string) [][]string {
	for ok := true; ok; {
		//println("Position: ", strconv.Itoa(row), ",", strconv.Itoa(column))
		//printGrid(levelMap)
		//println()
		//println()

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

		if levelMap[row][column] != "#" || levelMap[row][column] == "^" || levelMap[row][column] == "v" || levelMap[row][column] == "<" || levelMap[row][column] == ">" || levelMap[row][column] == "X" || levelMap[row][column] == "0" {
			if levelMap[row][column] != "0" {
				levelMap[row][column] = "X"
			}

			var neighbourNodeRow int
			var neighbourNodeColumn int

			switch currentdirection {
			case Up:
				if column > 0 {
					neighbourNodeRow = row
					neighbourNodeColumn = column + 1
				}
			case Down:
				if column < len(levelMap[0]) {
					neighbourNodeRow = row
					neighbourNodeColumn = column - 1
				}
			case Left:
				if row > 0 {
					neighbourNodeRow = row - 1
					neighbourNodeColumn = column
				}
			case Right:
				if row < len(levelMap) {
					neighbourNodeRow = row + 1
					neighbourNodeColumn = column
				}
			}

			if levelMap[neighbourNodeRow][neighbourNodeColumn] == "X" {
				//if levelMap[row][column] != "0" {
				//	switch currentdirection {
				//	case Up:
				//		levelMap[row][column] = "^"
				//	case Down:
				//		levelMap[row][column] = "v"
				//	case Left:
				//		levelMap[row][column] = "<"
				//	case Right:
				//		levelMap[row][column] = ">"
				//	}
				//}
				//println()
				//println()
				//for _, row := range levelMap {
				//	for column := range row {
				//		print(row[column])
				//	}
				//	println()
				//}
				//if levelMap[row][column] != "0" {
				//	levelMap[row][column] = "."
				//}

				//neighbour node has been visited before, we can loop here
				switch currentdirection {
				case Up:
					if column > 0 && levelMap[row-1][column] != "#" {
						canGoOutOfMap := true
						for i := column; i < len(levelMap[0]); i++ {
							//can we go out of the map?
							if levelMap[i][column] == "#" {
								canGoOutOfMap = false
								break
							}
						}

						if !canGoOutOfMap {
							levelMap[row-1][column] = "0"
						}
					}
				case Down:
					if row > 0 && levelMap[row+1][column] != "#" {
						canGoOutOfMap := true
						for i := column; i >= 0; i-- {
							//can we go out of the map?
							if levelMap[row][i] == "#" {
								canGoOutOfMap = false
								break
							}
						}

						if !canGoOutOfMap {
							levelMap[row+1][column] = "0"
						}
					}
				case Left:
					if row > 0 && levelMap[row][column-1] != "#" {
						canGoOutOfMap := true
						for i := row; i >= 0; i-- {
							//can we go out of the map?
							if levelMap[i][column] == "#" {
								canGoOutOfMap = false
								break
							}
						}

						if !canGoOutOfMap {
							levelMap[row][column-1] = "0"
						}
					}
				case Right:
					if row < len(levelMap) && levelMap[row][column+1] != "#" {
						canGoOutOfMap := true
						for i := row; i < len(levelMap); i++ {
							//can we go out of the map?
							if levelMap[i][column] == "#" {
								canGoOutOfMap = false
								break
							}
						}

						if !canGoOutOfMap {
							levelMap[row][column+1] = "0"
						}
					}
				}
			}
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
	return levelMap
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
