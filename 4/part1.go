package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	input := readInput("input.txt")
	tot := search(input)

	fmt.Println("total:", tot)
}

func readInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]string
	for scanner.Scan() {
		rowStr := scanner.Text()
		var row []string
		for _, char := range rowStr{
			row = append(row, string(char))
		}
		input = append(input, row)
	}

	return input
}

func findLetter(input [][]string, x int, y int, deltaX int, deltaY int, target string) int {
	newX := x + deltaX
	newY := y + deltaY

	if newX < 0 || newY < 0 {
		return 0
	}

	if input[newX][newY] == target {
		if target == "M" {
			return findLetter(input, newX, newY, deltaX, deltaY, "A")
		} else if target == "A" {
			return findLetter(input, newX, newY, deltaX, deltaY, "S")
		}
		return 1
	}
	return 0
}

func findWords(input [][]string, x int, y int) int {
	// given an X, find the number of "XMAS" branching off
	// returns an int
	tot := 0
	rows := len(input)
	cols := len(input[0])
	for deltaX := -1; deltaX <= 1; deltaX++ {
		for deltaY := -1; deltaY <= 1; deltaY++ {
			if deltaX == 0 && deltaY == 0 {
				continue
			}

			if (x + 3 * deltaX) >= rows || (y + 3 * deltaY) >= cols || (x + 3 * deltaX) < 0 || (y + 3 * deltaY) < 0{
				continue
			}

			tot += findLetter(input, x, y, deltaX, deltaY, "M")
		}
	}

	return tot
}

func search(input [][]string) int {
	tot := 0
	for x, row := range input {
		for y, cell := range row {
			if cell == "X" {
				tot += findWords(input, x, y)
			}
		}
	}

	return tot
}

