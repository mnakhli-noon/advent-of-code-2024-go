package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func readInput() ([][]rune, int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	x, y := 0, 0

	scanner := bufio.NewScanner(file)
	lab := [][]rune{}

	for scanner.Scan() {
		line := []rune(scanner.Text())
		if index := slices.Index(line, '^'); index >= 0 {
			y = index
			x = len(lab)
		}
		lab = append(lab, line)
	}

	return lab, x, y
}
func checkIfWeCanWalk(maze [][]rune, x int, y int) (bool, bool) {
	if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[x]) {
		return false, true
	}
	if maze[x][y] == '#' {
		return false, false
	}
	return true, false
}

func walkTheMaze(maze [][]rune, x int, y int) {
	if x < 0 || x >= len(maze) || y < 0 || y >= len(maze[x]) {
		return
	}
	switch maze[x][y] {
	case '^':
		if canWalk, finish := checkIfWeCanWalk(maze, x-1, y); canWalk {
			maze[x][y] = 'X'
			maze[x-1][y] = '^'
			walkTheMaze(maze, x-1, y)
			return
		} else {
			if finish {
				return
			} else {
				maze[x][y] = '>'
				walkTheMaze(maze, x, y)
				return
			}
		}
	case '>':
		if canWalk, finish := checkIfWeCanWalk(maze, x, y+1); canWalk {
			maze[x][y] = 'X'
			maze[x][y+1] = '>'
			walkTheMaze(maze, x, y+1)
			return
		} else {
			if finish {
				maze[x][y] = 'X'
				return
			} else {
				maze[x][y] = 'v'
				walkTheMaze(maze, x, y)
				return
			}
		}
	case 'v':
		if canWalk, finish := checkIfWeCanWalk(maze, x+1, y); canWalk {
			maze[x][y] = 'X'
			maze[x+1][y] = 'v'
			walkTheMaze(maze, x+1, y)
			return
		} else {
			if finish {
				maze[x][y] = 'X'
				return
			} else {
				maze[x][y] = '<'
				walkTheMaze(maze, x, y)
				return
			}
		}
	case '<':
		if canWalk, finish := checkIfWeCanWalk(maze, x, y-1); canWalk {
			maze[x][y] = 'X'
			maze[x][y-1] = '<'
			walkTheMaze(maze, x, y-1)
			return
		} else {
			if finish {
				maze[x][y] = 'X'
				return
			} else {
				maze[x][y] = '^'
				walkTheMaze(maze, x, y)
				return
			}
		}
	default:
		return
	}
}
func solvePartOne() {
	lab, x, y := readInput()
	walkTheMaze(lab, x, y)
	count := 0
	for _, line := range lab {
		for _, character := range line {
			if character == 'X' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	solvePartOne()
}
