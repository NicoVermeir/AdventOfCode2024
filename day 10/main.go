package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	grid := readGrid("input.txt")
	trailheads := findTrailheads(grid)

	count := 0
	for _, coord := range trailheads {
		visited := make(map[[2]int]bool)
		count += countTrails(grid, coord[0], coord[1], 0, visited)
	}

	fmt.Println("Number of trails:", count)
}

func partTwo() {
	grid := readGrid("input.txt")
	trailheads := findTrailheads(grid)

	count := 0
	for _, coord := range trailheads {
		count += countDistinctTrails(grid, coord[0], coord[1], 0)
	}

	fmt.Println("Number distinct trails:", count)
}

func findTrailheads(grid [][]int) [][2]int {
	var coordinates [][2]int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				coordinates = append(coordinates, [2]int{i, j})
			}
		}
	}

	return coordinates
}

func countTrails(grid [][]int, x, y, current int, visited map[[2]int]bool) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != current {
		return 0
	}

	if visited[[2]int{x, y}] {
		return 0
	}

	if current == 9 {
		visited[[2]int{x, y}] = true
		return 1
	}

	visited[[2]int{x, y}] = true
	next := current + 1
	return countTrails(grid, x+1, y, next, visited) +
		countTrails(grid, x-1, y, next, visited) +
		countTrails(grid, x, y+1, next, visited) +
		countTrails(grid, x, y-1, next, visited)
}

func countDistinctTrails(grid [][]int, x, y, current int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != current {
		return 0
	}

	//if visited[[2]int{x, y}] {
	//	return 0
	//}

	if current == 9 {
		//visited[[2]int{x, y}] = true
		return 1
	}

	//visited[[2]int{x, y}] = true
	next := current + 1
	return countDistinctTrails(grid, x+1, y, next) +
		countDistinctTrails(grid, x-1, y, next) +
		countDistinctTrails(grid, x, y+1, next) +
		countDistinctTrails(grid, x, y-1, next)
}

func readGrid(path string) [][]int {
	file, _ := os.Open(path)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var grid [][]int

	for _, line := range lines {
		var slice []int
		for j, _ := range line {
			value, _ := strconv.Atoi(string(line[j]))
			slice = append(slice, value)
		}

		grid = append(grid, slice)
	}

	return grid
}
