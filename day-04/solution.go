package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var M = [][]int{
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
	{0, -1},
	{1, -1},
}
var A = [][]int{
	{2, 0},
	{2, 2},
	{0, 2},
	{-2, 2},
	{-2, 0},
	{-2, -2},
	{0, -2},
	{2, -2},
}

var S = [][]int{
	{3, 0},
	{3, 3},
	{0, 3},
	{-3, 3},
	{-3, 0},
	{-3, -3},
	{0, -3},
	{3, -3},
}

func isInside(x int, y int, maxX int, maxY int) bool {
	if x >= 0 && x <= maxX && y >= 0 && y <= maxY {
		return true
	}
	return false
}

func checkIfItsAStart(characters [][]string, x int, y int) int {
	if characters[x][y] != "X" {
		return 0
	}
	count := 0

	for i := 0; i < 8; i++ {
		if isInside(x+S[i][0], y+S[i][1], len(characters)-1, len(characters[0])-1) &&
			characters[x+M[i][0]][y+M[i][1]] == "M" &&
			characters[x+A[i][0]][y+A[i][1]] == "A" &&
			characters[x+S[i][0]][y+S[i][1]] == "S" {
			count++
		}

	}
	return count
}

func solvePartOne() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	characters := make([][]string, len(lines))

	for i := 0; i < len(lines); i++ {
		row := make([]string, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			row[j] = string(lines[i][j])
		}
		characters[i] = row
	}

	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			count += checkIfItsAStart(characters, i, j)
		}
	}

	fmt.Println(count)
}

func checkLeftToRight(characters [][]string, x int, y int) bool {
	if characters[x-1][y-1] == "M" && characters[x+1][y+1] == "S" {
		return true
	}

	if characters[x-1][y-1] == "S" && characters[x+1][y+1] == "M" {
		return true
	}
	return false

}

func checkRightToLeft(characters [][]string, x int, y int) bool {
	if characters[x-1][y+1] == "M" && characters[x+1][y-1] == "S" {
		return true
	}

	if characters[x-1][y+1] == "S" && characters[x+1][y-1] == "M" {
		return true
	}
	return false

}

func checkIfItsAnX(characters [][]string, x int, y int) int {
	if characters[x][y] != "A" {
		return 0
	}
	if (x < 1) || y < 1 || x > len(characters)-2 || y > len(characters[0])-2 {
		return 0
	}

	if checkLeftToRight(characters, x, y) && checkRightToLeft(characters, x, y) {
		return 1
	}
	return 0
}

func solvePartTwo() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	characters := make([][]string, len(lines))

	for i := 0; i < len(lines); i++ {
		row := make([]string, len(lines[0]))
		for j := 0; j < len(lines[0]); j++ {
			row[j] = string(lines[i][j])
		}
		characters[i] = row
	}

	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			count += checkIfItsAnX(characters, i, j)
		}
	}

	fmt.Println(count)
}

func main() {
	isPartTwo := flag.Bool("partTwo", false, "Solve Part Two")
	flag.Parse()

	if *isPartTwo {
		solvePartTwo()
	} else {
		solvePartOne()
	}
}
