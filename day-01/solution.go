package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func solvePartOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	numberByte := make([]byte, 5)
	skipByte := make([]byte, 3)
	skipEnd := make([]byte, 1)
	firstArray := [1000]int{}
	secondArray := [1000]int{}

	for i := 0; i < 1000; i++ {

		data, err := file.Read(numberByte)
		if err != nil {
			panic(err)
		}

		number, err := strconv.Atoi(string(numberByte[:data]))
		if err != nil {
			panic(err)
		}

		firstArray[i] = number

		_, err = file.Read(skipByte)
		if err != nil {
			panic(err)
		}

		data, err = file.Read(numberByte)
		if err != nil {
			panic(err)
		}
		number, err = strconv.Atoi(string(numberByte[:data]))
		if err != nil {
			panic(err)
		}
		secondArray[i] = number

		_, err = file.Read(skipEnd)
		if err != nil {
			panic(err)
		}
	}

	slices.Sort(firstArray[:])
	slices.Sort(secondArray[:])

	distance := 0
	for i := 0; i < 1000; i++ {
		distance += int(math.Abs(float64(firstArray[i] - secondArray[i])))
	}
	fmt.Println("Distance:", distance)

}

func solvePartTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	numberByte := make([]byte, 5)
	skipByte := make([]byte, 3)
	skipEnd := make([]byte, 1)
	firstArray := [1000]int{}
	secondArray := [1000]int{}

	for i := 0; i < 1000; i++ {

		data, err := file.Read(numberByte)
		if err != nil {
			panic(err)
		}

		number, err := strconv.Atoi(string(numberByte[:data]))
		if err != nil {
			panic(err)
		}

		firstArray[i] = number

		_, err = file.Read(skipByte)
		if err != nil {
			panic(err)
		}

		data, err = file.Read(numberByte)
		if err != nil {
			panic(err)
		}
		number, err = strconv.Atoi(string(numberByte[:data]))
		if err != nil {
			panic(err)
		}
		secondArray[i] = number

		_, err = file.Read(skipEnd)
		if err != nil {
			panic(err)
		}
	}

	m := make(map[int]int)

	for i := 0; i < 1000; i++ {
		count, exist := m[secondArray[i]]
		if exist {
			m[secondArray[i]] = count + 1
		} else {
			m[secondArray[i]] = 1
		}
	}

	similarity := 0
	for i := 0; i < 1000; i++ {
		similarity += firstArray[i] * m[firstArray[i]]
	}

	fmt.Println("Similarity:", similarity)
}

func main() {
	isPartTwo := flag.Bool("partTwo", false, "Solve part two")
	claude := flag.Bool("claude", false, "Solve it the claude way")

	flag.Parse()
	if *claude {
		if *isPartTwo {
			solvePartTwoClaude()
		} else {
			solvePartOneClaude()
		}
	} else {
		if *isPartTwo {
			solvePartTwo()
		} else {
			solvePartOne()
		}
	}
}
