package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	stoneString := readLines("input.txt")
	stones := strings.Split(stoneString, " ")
	var stoneNumbers []int

	for _, stone := range stones {
		number, _ := strconv.Atoi(stone)
		stoneNumbers = append(stoneNumbers, number)
	}

	totalCount := blink(stoneNumbers, 75)

	println("Number of stones:", totalCount)
}

func blink(stones []int, amountOfTimes int) int {

	tempStones := stones
	stoneMap := make(map[int]int)

	for _, stone := range tempStones {
		stoneMap[stone] = 1
	}

	for i := 0; i < amountOfTimes; i++ {
		tempMap := make(map[int]int)

		for key, value := range stoneMap {

			if key == 0 {
				tempMap[1] += value
				continue
			}

			digitCount := countDigits(key)
			if digitCount%2 == 0 {
				firstPart, lastPart := splitDigits(key, digitCount)
				tempMap[firstPart] += value
				tempMap[lastPart] += value
				continue
			}

			tempMap[key*2024] += value
		}

		stoneMap = tempMap
		println("Loop ", i)
	}

	count := 0

	for _, value := range stoneMap {
		count += value
	}
	return count
}

func findIndices(slice []int, value int) []int {
	var indices []int
	for i, v := range slice {
		if v == value {
			indices = append(indices, i)
		}
	}
	return indices
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}

	return int(math.Log10(float64(n))) + 1
}

func splitDigits(stone int, numberOfDigits int) (int, int) {
	divisor := int(math.Pow(10, float64(numberOfDigits/2)))
	firstPart := stone / divisor // Get the first two digits
	lastPart := stone % divisor  // Get the last two digits

	return firstPart, lastPart
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) string {
	content, err := os.ReadFile(path)
	text := string(content)
	if err != nil {
		log.Fatal(err)
	}

	return text
}
