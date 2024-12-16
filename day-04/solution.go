package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func isRow(characters [][]string, x int, y int) bool {
	if y > len(characters[0])-4 {
		return false
	}

	if characters[x][y+1] != "M" || characters[x][y+2] != "A" || characters[x][y+3] != "S" {
		return false
	}

	return true
}

func isReverseRow(characters [][]string, x int, y int) bool {
	if y < 3 {
		return false
	}

	if characters[x][y-1] != "M" || characters[x][y-2] != "A" || characters[x][y-3] != "S" {
		return false
	}

	return true
}

func isCol(characters [][]string, x int, y int) bool {
	if x > len(characters)-4 {
		return false
	}

	if characters[x+1][y] != "M" || characters[x+2][y] != "A" || characters[x+3][y] != "S" {
		return false
	}

	return true
}

func isReverseCol(characters [][]string, x int, y int) bool {
	if x < 3 {
		return false
	}

	if characters[x-1][y] != "M" || characters[x-2][y] != "A" || characters[x-3][y] != "S" {
		return false
	}

	return true
}

func isDiagonalLeft(characters [][]string, x int, y int) bool {
	if x > len(characters)-4 || y < 3 {
		return false
	}
	if characters[x+1][y-1] != "M" || characters[x+2][y-2] != "A" || characters[x+3][y-3] != "S" {
		return false
	}

	return true
}

func isReverseDiagonalLeft(characters [][]string, x int, y int) bool {
	if x < 3 || y-3 < 0 {
		return false
	}
	if characters[x-1][y-1] != "M" || characters[x-2][y-2] != "A" || characters[x-3][y-3] != "S" {
		return false
	}

	return true
}

func isDiagonalRight(characters [][]string, x int, y int) bool {
	if x > len(characters)-4 || y > len(characters[0])-4 {
		return false
	}

	if characters[x+1][y+1] != "M" || characters[x+2][y+2] != "A" || characters[x+3][y+3] != "S" {
		return false
	}

	return true
}

func isReverseDiagonalRight(characters [][]string, x int, y int) bool {
	if x < 3 || y > len(characters[0])-4 {
		return false
	}
	if characters[x-1][y+1] != "M" || characters[x-2][y+2] != "A" || characters[x-3][y+3] != "S" {
		return false
	}

	return true
}

func checkIfItsAStart(characters [][]string, x int, y int) int {
	if characters[x][y] != "X" {
		return 0
	}
	count := 0

	if isRow(characters, x, y) {
		count++
	}

	if isReverseRow(characters, x, y) {
		count++
	}

	if isCol(characters, x, y) {
		count++
	}

	if isReverseCol(characters, x, y) {
		count++
	}

	if isDiagonalLeft(characters, x, y) {
		count++
	}

	if isReverseDiagonalLeft(characters, x, y) {
		count++
	}

	if isDiagonalRight(characters, x, y) {
		count++
	}

	if isReverseDiagonalRight(characters, x, y) {
		count++
	}
	return count
}

func solvePartOne() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	rows := len(lines)
	cols := len(lines[0])

	characters := make([][]string, rows)

	for i := 0; i < rows; i++ {
		row := make([]string, cols)
		for j := 0; j < cols; j++ {
			row[j] = string(lines[i][j])
		}
		characters[i] = row
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
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
	rows := len(lines)
	cols := len(lines[0])

	characters := make([][]string, rows)

	for i := 0; i < rows; i++ {
		row := make([]string, cols)
		for j := 0; j < cols; j++ {
			row[j] = string(lines[i][j])
		}
		characters[i] = row
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
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
