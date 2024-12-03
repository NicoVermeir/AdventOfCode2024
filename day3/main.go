package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	//read the file
	content, err := os.ReadFile("input.txt")
	text := string(content)
	if err != nil {
		log.Fatal(err)
	}

	mulOperations := ExtractMulOperations(text)
	println("Total of all Mul calculations: ", calculateTotal(mulOperations))
}

func partTwo() {
	//read the file
	content, err := os.ReadFile("input.txt")
	text := string(content)
	if err != nil {
		log.Fatal(err)
	}

	mulOperations := extractOperations(text)
	println("Total of all Mul calculations: ", calculateTotalWithDisableOption(mulOperations))
}

func calculateTotalWithDisableOption(operations []string) int {

	isDisabled := false
	var total int

	for _, operation := range operations {
		if strings.HasPrefix(operation, "mul") && !isDisabled {
			value1, value2 := MatchAndParse(operation)
			total += value1 * value2
		}

		if operation == "don't" {
			isDisabled = true
		} else if operation == "do" {
			isDisabled = false
		}
	}

	return total
}

func extractOperations(text string) []string {
	regexPattern := regexp.MustCompile("don\\'t|do|mul[(][0-9]*,[0-9]*[)]")
	return regexPattern.FindAllString(text, -1)
}

func ExtractMulOperations(text string) []string {
	regexPattern := regexp.MustCompile("mul[(][0-9]*,[0-9]*[)]")

	return regexPattern.FindAllString(text, -1)
}

func MatchAndParse(text string) (int, int) {
	// Define a regular expression to extract numbers from the string "mul(x,y)"
	numRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Find submatches of the numbers
	matches := numRegex.FindStringSubmatch(text)
	if matches != nil && len(matches) == 3 {
		// Convert the extracted strings to integers
		num1, err1 := strconv.Atoi(matches[1])
		num2, err2 := strconv.Atoi(matches[2])

		if err1 == nil && err2 == nil {
			// Now you have both numbers as integers
			log.Printf("Extracted numbers: %d, %d", num1, num2)
		} else {
			log.Printf("Error converting strings to integers")
		}

		return num1, num2
	}

	return 0, 0
}

func calculateTotal(operations []string) int {
	var total int

	for _, operation := range operations {
		value1, value2 := MatchAndParse(operation)
		total += value1 * value2
	}

	return total
}
