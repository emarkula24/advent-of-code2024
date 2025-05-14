package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1("./day3-input")
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int
	for scanner.Scan() {
		r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
		if err != nil {
			panic(err)
		}
		line := scanner.Text()
		matches := r.FindAllString(line, -1)
		for _, match := range matches {

			r := regexp.MustCompile(`\d{1,3}`)
			numbers := r.FindAllString(match, -1)
			var intNumbers []int

			for _, numStr := range numbers {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err) // Handle conversion error
				}
				intNumbers = append(intNumbers, num)
			}

			total += intNumbers[0] * intNumbers[1]

		}
	}
	fmt.Println(total)
}
