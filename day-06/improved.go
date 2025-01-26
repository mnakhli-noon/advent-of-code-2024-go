package main

import (
	"bufio"
	"fmt"
	"os"
)

type State2 struct {
	x, y, dir int
}

func main2() {
	grid, startX, startY, initialDir := readInput2()
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			// Skip starting position and existing obstructions
			if (x == startX && y == startY) || grid[y][x] == '#' {
				continue
			}

			// Create modified grid with new obstruction
			modified := copyGrid(grid)
			modified[y][x] = '#'

			if simulate(modified, startX, startY, initialDir) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func readInput2() ([][]rune, int, int, int) {
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
	dirs := []struct{ dx, dy int }{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // Up, Right, Down, Left
	visited := make(map[State2]bool)
	x, y, currentDir := startX, startY, initialDir

	for {
		state := State2{x, y, currentDir}
		if visited[state] {
			return true // Loop detected
		}
		visited[state] = true

		// Check next position
		dx := dirs[currentDir].dx
		dy := dirs[currentDir].dy
		nx := x + dx
		ny := y + dy

		// Check if next cell is blocked or out of bounds
		if ny >= 0 && ny < len(grid) && nx >= 0 && nx < len(grid[ny]) && grid[ny][nx] == '#' {
			currentDir = (currentDir + 1) % 4 // Turn right
		} else {
			// Move forward
			x, y = nx, ny
			// Check if out of grid
			if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
				return false // Escaped, no loop
			}
		}
	}
}
