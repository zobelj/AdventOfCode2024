package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"bufio"
	"slices"
	"strings"
	"strconv"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	rules := getRules(file)
	updates := getUpdates(file)

	// check each update against the rules
	updates = filterUpdates(updates, rules)
	updates = fixUpdates(updates, rules)
	tot := sumMiddles(updates)

	fmt.Println("total:", tot)
}

func getMiddleInt(update []int) int {
	length := len(update)
	midIdx := (length - 1) / 2

	return update[midIdx]
}

func sumMiddles(updates[][]int) int {
	tot := 0
	for _, update := range updates {
		tot += getMiddleInt(update)
	}

	return tot
}

func isCorrect(update []int, rules map[int][]int) bool {
	size := len(update)
	for i := 0; i < size; i++ {
		checkVals := rules[update[i]]

		for _, val := range checkVals {
			idx := slices.Index(update, val)

			if (idx != -1) && (idx <= i) {
				return false
			}
		}
	}

	return true
}

func filterUpdates(updates [][]int, rules map[int][]int) [][]int {
	correctUpdates := [][]int{}

	for _, update := range updates {
		if !isCorrect(update, rules) {
			correctUpdates = append(correctUpdates, update)
		}
	}

	return correctUpdates
}

func fixUpdate(update []int, rules map[int][]int) []int {
	fixedUpdate := []int{}
	fixedUpdate = append(fixedUpdate, update[0])

	for _, val := range update[1:] {
		added := false
		for idx2, rule := range fixedUpdate {
			// if rule is in rules[val], we need to go insert just before rule
			// if val is in rules[rule], we have to keep going
			if !added && slices.Contains(rules[val], rule) {
				//fixedUpdate = append(fixedUpdate[:idx2+1], fixedUpdate[idx2:]...)
				//fixedUpdate[idx2] = val
				fixedUpdate = slices.Insert(fixedUpdate, idx2, val)
				added = true
			}
		}
		if !added {
			fixedUpdate = append(fixedUpdate, val)
		}
	}

	return fixedUpdate
}

func fixUpdates(updates [][]int, rules map[int][]int) [][]int {
	fixedUpdates := [][]int{}

	for _, val := range updates {
		fixedUpdates = append(fixedUpdates, fixUpdate(val, rules))
	}

	return fixedUpdates
}

func getRules(file *os.File) map[int][]int {
	// read in the file one line at a time
	// done when you reach an empty line
	scanner := bufio.NewScanner(file)
	rules := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			rulePair := strings.Split(line, "|")
			rule0, _ := strconv.Atoi(rulePair[0])
			rule1, _ := strconv.Atoi(rulePair[1])

			if rules[rule0] == nil {
				rules[rule0] = make([]int, 1)
			}
			rules[rule0] = append(rules[rule0], rule1)
		} else {
			break
		}
	}

	// reset the file pointer
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	return rules
}

func getUpdates(file *os.File) [][]int {
	// read in the file one line at a time
	// only start after the empty line
	// each line is its own slice of integers
	scanner := bufio.NewScanner(file)
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.Contains(line, "|") {
			continue
		}

		update := []int{}
		valStrings := strings.Split(line, ",")
		for _, str := range valStrings {
			valInt, _ := strconv.Atoi(str)
			update = append(update, valInt)
		}
		updates = append(updates, update)
	}

	return updates
}

