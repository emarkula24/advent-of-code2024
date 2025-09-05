package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sum int
var sum2 int

func main() {
	a := parseFile("day7.txt")

	for _, line := range a {
		solveEquation(line)
	}

	fmt.Println("Sum of all equation values:", sum)
	fmt.Println("Sum of all concat values:", sum2)
}

func solveEquation(line []int) bool {
	if len(line) < 2 {
		return false
	}

	target := line[0]
	numbers := line[1:]

	return helper(numbers, target, numbers[0], 1)
}

func helper(numbers []int, target, currentVal, index int) bool {
	if index == len(numbers) {
		if currentVal == target {
			sum += currentVal
			sum2 += currentVal
			return true
		}
		return false
	}

	nextNum := numbers[index]

	if helper(numbers, target, currentVal+nextNum, index+1) {
		return true
	}

	if helper(numbers, target, currentVal*nextNum, index+1) {
		return true
	}
	// part1 answer can be viewed by commentin out this part marked with comments.
	concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentVal, nextNum))
	if helper(numbers, target, concat, index+1) {
		return true
	}
	//

	return false
}

func parseFile(filename string) (equations [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		left := strings.TrimSpace(parts[0])
		right := strings.Fields(parts[1])

		combined := append([]string{left}, right...)
		lineInt := convertSliceToInt(combined)
		equations = append(equations, lineInt)
	}

	return equations

}

func convertSliceToInt(slice []string) (intSlice []int) {
	for _, str := range slice {
		str = strings.TrimSpace(str)
		j, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, j)
	}
	return intSlice
}
