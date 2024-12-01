package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var leftList []int
var rightList []int

func main() {
	partOne()
	partTwo()
}

func partOne() {
	//read the file
	lines, err := readLines("input.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	parseLists(lines)
	sortLists()

	//calculate the difference
	var diff int = 0

	for i := range leftList {
		if leftList[i] > rightList[i] {
			diff += leftList[i] - rightList[i]
		} else {
			diff += rightList[i] - leftList[i]

		}
	}

	//print final result
	fmt.Println("diff: ", diff)
}

func partTwo() {
	//calculate similarity score
	var score int = 0

	for i := range leftList {
		var similarity int = count(leftList[i])

		score += leftList[i] * similarity
	}

	println("similarity score: ", score)
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

func parseLists(lines []string) {
	//parse all lines into two lists
	for i := range lines {
		values := strings.Fields(lines[i])

		if len(values) == 2 {
			if leftValue, err := strconv.Atoi(values[0]); err == nil {
				leftList = append(leftList, leftValue)
			}

			if rightValue, err := strconv.Atoi(values[1]); err == nil {
				rightList = append(rightList, rightValue)
			}
		}
	}
}

func sortLists() {
	sort.Ints(leftList)
	sort.Ints(rightList)
}

func count(leftValue int) int {
	var count int = 0

	for _, rightValue := range rightList {
		if rightValue == leftValue {
			count++
		}
	}
	return count
}
