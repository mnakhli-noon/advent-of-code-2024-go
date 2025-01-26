package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type Grid [][]rune
type WalkDirection int

const (
	UP WalkDirection = iota
	RIGHT
	DOWN
	LEFT
)

type Position struct {
	x int
	y int
}

var walkPosition = map[WalkDirection]Position{
	UP:    {x: -1, y: 0},
	RIGHT: {x: 0, y: 1},
	DOWN:  {x: 1, y: 0},
	LEFT:  {x: 0, y: -1},
}

var nextDircetion = map[WalkDirection]WalkDirection{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}

type State struct {
	x, y int
	dir  WalkDirection
}

func readInput() (Grid, int, int) {
	file, err := os.Open("input.txt")
	defer file.Close()

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

func (g Grid) checkIfWeCanWalk(x int, y int, dir WalkDirection) (bool, bool, int, int) {
	delta := walkPosition[dir]
	newX := x + delta.x
	newY := y + delta.y

	if newX < 0 || newX >= len(g) || newY < 0 || newY >= len(g[newX]) {
		return false, true, -1, -1
	}

	if g[newX][newY] == '#' {
		return false, false, x, y
	}

	return true, false, newX, newY
}

func (g Grid) walkTheMaze(x int, y int, dir WalkDirection) {
	if x < 0 || x >= len(g) || y < 0 || y >= len(g[x]) {
		return
	}
	canWalk, finish, newX, newY := g.checkIfWeCanWalk(x, y, dir)

	if canWalk {
		g[x][y] = 'X'
		g.walkTheMaze(newX, newY, dir)
	} else if finish {
		g[x][y] = 'X'
		return
	} else {
		g.walkTheMaze(x, y, nextDircetion[dir])
	}
}

func solvePartOne() {
	lab, x, y := readInput()
	lab.walkTheMaze(x, y, UP)
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

func (g Grid) copyGrid() Grid {
	newG := make(Grid, len(g))
	for i, line := range g {
		newG[i] = make([]rune, len(g[i]))
		copy(newG[i], line)
	}
	return newG
}

func (g Grid) checkIfPositionInside(x int, y int) bool {
	return x >= 0 && x < len(g) && y >= 0 && y < len(g[x])
}

func (g Grid) hasLoop(x int, y int, dir WalkDirection) bool {
	visitied := make(map[State]bool)

	currentX, currentY := x, y
	for {
		state := State{currentX, currentY, dir}
		if visitied[state] {
			return true
		}
		visitied[state] = true

		nextX := currentX + walkPosition[dir].x
		nextY := currentY + walkPosition[dir].y
		if g.checkIfPositionInside(nextX, nextY) {
			if g[nextX][nextY] == '#' {
				dir = nextDircetion[dir]
			} else {
				currentX = nextX
				currentY = nextY
			}
		} else {
			return false
		}
	}
}

func (g Grid) countTheLoops(x int, y int) int {
	count := 0

	for i := range g {
		for j := range g[i] {
			if (i == x && j == y) || g[i][j] == '#' {
				continue
			}

			modified := make(Grid, len(g))
			modified = g.copyGrid()
			modified[i][j] = '#'

			if modified.hasLoop(x, y, UP) {
				count++
			}
		}
	}
	return count
}

func solvePartTwo() {
	lab, x, y := readInput()

	loops := lab.countTheLoops(x, y)

	fmt.Println(loops)
}

func main() {
	solvePartTwo()
}

