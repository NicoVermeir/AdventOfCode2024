package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	evaluateEquations(lines)

	println("Sum of valid equations: ", testResult)
}

var testResult int = 0

func findCombinations(numbers []int, target int, current int, index int, path []string) bool {
	if index == len(numbers) {
		if current == target {
			fmt.Println(path)
			testResult += target
		}
		return current == target
	}

	// Add the current number
	isSuccess := findCombinations(numbers, target, current+numbers[index], index+1, append(path, fmt.Sprintf("+%d", numbers[index])))

	// Multiply the current number
	if !isSuccess {
		isSuccess = findCombinations(numbers, target, current*numbers[index], index+1, append(path, fmt.Sprintf("*%d", numbers[index])))
	}

	// Concatenate the current number
	if !isSuccess {
		concatenated, _ := strconv.Atoi(fmt.Sprintf("%d%d", current, numbers[index]))
		isSuccess = findCombinations(numbers, target, concatenated, index+1, append(path, fmt.Sprintf("||%d", numbers[index])))
	}

	return isSuccess
}

func evaluateEquations(lines []string) {
	for _, line := range lines {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		var numbers []int

		for _, field := range strings.Fields(split[1]) {
			parsedNumber, _ := strconv.Atoi(field)
			numbers = append(numbers, parsedNumber)
		}

		findCombinations(numbers, result, 0, 0, []string{})
	}
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
