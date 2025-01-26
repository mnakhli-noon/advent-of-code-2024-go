package main

import (
	"bufio"
	"fmt"
	"os"
)

type StateImproved struct {
	x, y, dir int
}

func solvePartOneImproved() {
	grid, startX, startY, initialDir := readInputImproved()
	count := simulatePartOne(grid, startX, startY, initialDir)
	fmt.Println(count)
}

func solvePartTwoImproved() {
	grid, startX, startY, initialDir := readInputImproved()
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if (x == startX && y == startY) || grid[y][x] == '#' {
				continue
			}

			modified := copyGrid(grid)
			modified[y][x] = '#'

			if simulate(modified, startX, startY, initialDir) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func readInputImproved() ([][]rune, int, int, int) {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	startX, startY := -1, -1
	initialDir := -1
	dirMap := map[rune]int{'^': 0, '>': 1, 'v': 2, '<': 3}

	y := 0
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
		for x, c := range line {
			if c == '^' || c == '>' || c == 'v' || c == '<' {
				startX, startY = x, y
				initialDir = dirMap[c]
				grid[y][x] = '.' // Replace starting symbol
			}
		}
		y++
	}
	return grid, startX, startY, initialDir
}

func copyGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func simulate(grid [][]rune, startX, startY, initialDir int) bool {
	dirs := []struct{ dx, dy int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	visited := make(map[StateImproved]bool)
	x, y, currentDir := startX, startY, initialDir

	for {
		state := StateImproved{x, y, currentDir}
		if visited[state] {
			return true
		}
		visited[state] = true

		dx := dirs[currentDir].dx
		dy := dirs[currentDir].dy
		nx := x + dx
		ny := y + dy

		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) && grid[ny][nx] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			x, y = nx, ny
			if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
				return false
			}
		}
	}
}

func simulatePartOne(grid [][]rune, startX, startY, initialDir int) int {
	dirs := []struct{ dx, dy int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	currentDir := initialDir
	x, y := startX, startY
	visited := make(map[[2]int]bool)
	visited[[2]int{x, y}] = true

	for {
		dx := dirs[currentDir].dx
		dy := dirs[currentDir].dy
		nx := x + dx
		ny := y + dy

		isObstruction := false
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) {
			if grid[ny][nx] == '#' {
				isObstruction = true
			}
		}

		if isObstruction {
			currentDir = (currentDir + 1) % 4
		} else {
			x, y = nx, ny
			if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
				break
			}
			visited[[2]int{x, y}] = true
		}
	}

	return len(visited)
}
