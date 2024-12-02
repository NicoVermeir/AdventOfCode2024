package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	//read the file
	lines, err := readLines("input.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	numberOfReports := 0
	numberOfSafeReports := 0
	data := parseLines(lines)

	for line := range data {
		numberOfReports++
		if isSafeReport(line) {
			numberOfSafeReports++
		}
	}

	println("Number of safe reports: ", numberOfSafeReports)
	println("Number of reports: ", numberOfReports)
}

func partTwo() {
	//read the file
	lines, err := readLines("input.txt")

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	numberOfReports := 0
	numberOfSafeReports := 0
	data := parseLines(lines)

	for line := range data {
		numberOfReports++
		if isSafeReportWithRecursion(line, false) {
			numberOfSafeReports++
		}
	}

	println("Number of safe reports: ", numberOfSafeReports)
	println("Number of reports: ", numberOfReports)
}

func isSafeReport(line []string) bool {
	isIncreasing := false

	for i := 1; i < len(line); i++ {
		//is the level increasing or decreasing?
		if line[i] == line[i-1] {
			return false
		}

		//parse string to int
		currentReport, err1 := strconv.Atoi(line[i])
		previousReport, err2 := strconv.Atoi(line[i-1])

		//check if report is consistently ascending or descending
		if currentReport > previousReport {
			if i > 1 && !isIncreasing {
				return false
			}
			isIncreasing = true
		} else {
			if i > 1 && isIncreasing {
				return false
			}
			isIncreasing = false
		}

		if err1 != nil {
			log.Fatalf("Error parsing int: %s", err1)
		}

		if err2 != nil {
			log.Fatalf("Error parsing int: %s", err2)
		}

		diff := Abs(currentReport - previousReport)

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isSafeReportWithRecursion(line []string, isSecondRun bool) bool {
	isIncreasing := false

	for i := 1; i < len(line); i++ {
		isSafeReport := true

		//is the level increasing or decreasing?
		if line[i] == line[i-1] {
			isSafeReport = false
		}

		//parse string to int
		currentReport, err1 := strconv.Atoi(line[i])
		previousReport, err2 := strconv.Atoi(line[i-1])

		//check if report is consistently ascending or descending
		if currentReport > previousReport {
			if i > 1 && !isIncreasing {
				isSafeReport = false
			}
			isIncreasing = true
		} else {
			if i > 1 && isIncreasing {
				isSafeReport = false
			}
			isIncreasing = false
		}

		if err1 != nil {
			log.Fatalf("Error parsing int: %s", err1)
		}

		if err2 != nil {
			log.Fatalf("Error parsing int: %s", err2)
		}

		diff := Abs(currentReport - previousReport)

		if diff < 1 || diff > 3 {
			isSafeReport = false
		}

		if !isSafeReport {
			if isSecondRun {
				return false
			}

			for j := 0; j < len(line); j++ {
				temp := append([]string{}, line...)
				fixedLine := slices.Delete(temp, j, j+1)

				if isSafeReportWithRecursion(fixedLine, true) {
					return true
				}
			}

			return false
		}
	}

	return true
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

// Go channel, to mimic yield return in C#
func parseLines(lines []string) <-chan []string {
	channel := make(chan []string)

	go func() {
		defer close(channel)

		for i := range lines {
			channel <- strings.Fields(lines[i])
		}
	}()

	return channel
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
