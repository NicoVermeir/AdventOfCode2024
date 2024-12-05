package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var rules []rule

func main() {
	partOne()
	partTwo()
}

func partOne() {
	parseRules()

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var grid [][]string

	for _, line := range lines {
		split := strings.Split(line, ",")
		var slicedString []string

		for _, value := range split {
			slicedString = append(slicedString, value)
		}

		grid = append(grid, slicedString)
	}

	result := validateRules(grid, rules)
	println("Sum is: ", result)
}

func partTwo() {
	parseRules()

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var grid [][]string

	for _, line := range lines {
		split := strings.Split(line, ",")
		var slicedString []string

		for _, value := range split {
			slicedString = append(slicedString, value)
		}

		grid = append(grid, slicedString)
	}

	invalidupdates := findInvalidUpdates(grid, rules)
	result := fixUpdates(invalidupdates, rules)

	println("Sum of fixed updates is: ", result)
}

func fixUpdates(invalidupdates [][]string, rules []rule) int {
	var fixedUpdates [][]string

	for i := 0; i < len(invalidupdates); i++ {
		for j := 0; j < len(invalidupdates[i])-1; j++ {
			firstnumber, _ := strconv.Atoi(invalidupdates[i][j])
			secondnumber, _ := strconv.Atoi(invalidupdates[i][j+1])

			for _, rule := range rules {
				if rule.mustComeBefore == firstnumber && rule.pagenumber == secondnumber {
					//invalid, fix order
					temp := invalidupdates[i][j]
					invalidupdates[i][j] = invalidupdates[i][j+1]
					invalidupdates[i][j+1] = temp

					//go back in case we broke the previous order
					j = -1
					break
				}
			}
		}

		fixedUpdates = append(fixedUpdates, invalidupdates[i])
	}

	sum := 0
	for i := 0; i < len(fixedUpdates); i++ {
		index := (len(fixedUpdates[i]) - 1) / 2

		middlevalue, _ := strconv.Atoi(fixedUpdates[i][index])
		sum += middlevalue
	}

	return sum
}

func findInvalidUpdates(grid [][]string, r []rule) [][]string {
	var invalidUpdates [][]string

	for i := 0; i < len(grid); i++ {
		isvalid := true

		for j := 0; j < len(grid[i])-1; j++ {
			firstnumber, _ := strconv.Atoi(grid[i][j])
			secondnumber, _ := strconv.Atoi(grid[i][j+1])

			if !isvalid {
				break
			}

			for _, rule := range rules {
				if rule.mustComeBefore == firstnumber && rule.pagenumber == secondnumber {
					//invalid
					isvalid = false
					break
				}
			}
		}

		if !isvalid {
			invalidUpdates = append(invalidUpdates, grid[i])
			isvalid = true
			continue
		}
	}

	return invalidUpdates
}

func validateRules(grid [][]string, r []rule) int {
	var validUpdates [][]string

	for i := 0; i < len(grid); i++ {
		isvalid := true

		for j := 0; j < len(grid[i])-1; j++ {
			firstnumber, _ := strconv.Atoi(grid[i][j])
			secondnumber, _ := strconv.Atoi(grid[i][j+1])

			if !isvalid {
				break
			}

			for _, rule := range rules {
				if rule.mustComeBefore == firstnumber && rule.pagenumber == secondnumber {
					//invalid
					isvalid = false
					break
				}
			}
		}

		if !isvalid {
			isvalid = true
			continue
		}

		validUpdates = append(validUpdates, grid[i])
	}

	sum := 0
	for i := 0; i < len(validUpdates); i++ {
		index := (len(validUpdates[i]) - 1) / 2

		middlevalue, _ := strconv.Atoi(validUpdates[i][index])
		sum += middlevalue
	}

	return sum
}

func parseRules() {
	ruleLines, err := readLines("rules.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for _, ruleString := range ruleLines {
		split := strings.Split(ruleString, "|")
		pagenumber, _ := strconv.Atoi(split[0])
		mustComeBefore, _ := strconv.Atoi(split[1])
		parsedRule := rule{pagenumber: pagenumber, mustComeBefore: mustComeBefore}
		rules = append(rules, parsedRule)
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

type rule struct {
	mustComeBefore int
	pagenumber     int
}
