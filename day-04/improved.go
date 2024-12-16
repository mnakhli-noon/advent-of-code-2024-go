// this improved version was inspired by https://www.youtube.com/@josh_ackland
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func solvePartOneImproved() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Couldn't read file: ", err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	characters := make([][]rune, len(lines))
	for i, line := range lines {
		characters[i] = []rune(line)
	}

	directions := [8][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1}}
	word := "XMAS"
	count := 0

	for row := 0; row < len(characters); row++ {
		for col := 0; col < len(characters[row]); col++ {
			if characters[row][col] != 'X' {
				continue
			}
			for _, dir := range directions {
				x := dir[0]
				y := dir[1]
				isXmas := true

				for charIndex := 1; charIndex < len(word); charIndex++ {
					offsetX := row + (x * charIndex)
					offsetY := col + (y * charIndex)

					if offsetX < 0 || offsetX >= len(characters) || offsetY < 0 || offsetY >= len(characters[row]) {
						isXmas = false
						break
					}

					if characters[offsetX][offsetY] != rune(word[charIndex]) {
						isXmas = false
						break
					}
				}
				if isXmas {
					count++
				}
			}

		}
	}

	fmt.Println(count)
}
