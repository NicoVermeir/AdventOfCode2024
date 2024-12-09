package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partOne()
}

func partOne() {
	antiNodeCount := calculateVectors(readLines("input.txt"))

	println("Number of antinodes: ", antiNodeCount)
}

func calculateVectors(grid []string) int {
	nodes := make(map[rune][][2]int)
	antinodes := make(map[string]struct{})

	// Collect positions of all nodes
	for i, row := range grid {
		for j, char := range row {
			if char != '.' {
				nodes[char] = append(nodes[char], [2]int{i, j})
			}
		}
	}

	// Calculate vectors for each value
	for _, positions := range nodes {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				x1, y1 := positions[i][0], positions[i][1]
				x2, y2 := positions[j][0], positions[j][1]
				vector := [2]int{x2 - x1, y2 - y1}

				//add the antennae as antinode
				antinodes[fmt.Sprintf("%d,%d", x1, y1)] = struct{}{}
				antinodes[fmt.Sprintf("%d,%d", x2, y2)] = struct{}{}

				// Apply vector backward from the first node
				bx, by := x1, y1
				for {
					bx, by = bx-vector[0], by-vector[1]
					if bx < 0 || bx >= len(grid) || by < 0 || by >= len(grid[0]) {
						break
					}
					antinodes[fmt.Sprintf("%d,%d", bx, by)] = struct{}{}
					fmt.Printf("Backward from (%d, %d): (%d, %d)\n", x1, y1, bx, by)
				}

				// Apply vector forward from the second node
				fx, fy := x2, y2
				for {
					fx, fy = fx+vector[0], fy+vector[1]
					if fx < 0 || fx >= len(grid) || fy < 0 || fy >= len(grid[0]) {
						break
					}
					antinodes[fmt.Sprintf("%d,%d", fx, fy)] = struct{}{}
					fmt.Printf("Forward from (%d, %d): (%d, %d)\n", x2, y2, fx, fy)
				}
			}
		}
	}

	return len(antinodes)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
