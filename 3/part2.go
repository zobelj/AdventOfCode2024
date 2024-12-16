package main

import (
	"fmt"
	"os"
	"log"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	input := readInput("input.txt")

	matches := getMuls(input)
	matches = filterDonts(matches)

	tot := 0
	for _, match := range matches {
		tot += mul(match)
	}

	fmt.Println("total:", tot)
}

func readInput(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	input := string(file)

	return input
}

func getMuls(input string) []string {
	r := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\))`)
	matches := r.FindAllString(input, -1)

	return matches
}

func mul(input string) int {
	r := regexp.MustCompile(`[0-9]+`)
	matches := r.FindAllString(input, -1)

	tot := 1
	for _, v := range matches {
		val, _ := strconv.Atoi(v)
		tot *= val
	}

	return tot
}

func filterDonts(matches []string) []string {
	var keeps []string
	keep := true
	for _, val := range matches {
		if strings.Contains(val, "mul") && keep {
			keeps = append(keeps, val)
		} else if val == "don't()" {
			keep = false
		} else if val == "do()" {
			keep = true
		}
	}

	return keeps
}

