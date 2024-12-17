package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
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

func checkX(input [][]string, x int, y int) int {
	// make sure the 4 corners are in bounds
	rows := len(input)
	cols := len(input[0])
	if x + 1 >= rows || x - 1 < 0 || y + 1 >= cols || y - 1 < 0 {
		return 0
	}

	// check if the top left and bottom right are S or M, and not the same
	topLeft := input[x-1][y-1]
	botRight := input[x+1][y+1]
	if !strings.Contains("SM", topLeft) || !strings.Contains("SM", botRight) || topLeft == botRight {
		return 0
	}

	// now check top right and bottom left
	topRight := input[x-1][y+1]
	botLeft := input[x+1][y-1]
	if !strings.Contains("SM", topRight) || !strings.Contains("SM", botLeft) || topRight == botLeft {
		return 0
	}

	return 1
}

func search(input [][]string) int {
	tot := 0
	for x, row := range input {
		for y, cell := range row {
			if cell == "A" {
				tot += checkX(input, x, y)
			}
		}
	}

	return tot
}

