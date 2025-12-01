package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := parseFile("input.txt")
	part1(instructions)
	part2(instructions)
}
func part2(instructions []string) {
	current := 50
	count := 0
	for _, instruction := range instructions {

		if strings.HasPrefix(instruction, "R") {
			split := strings.Split(instruction, "R")
			number := split[1]
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			next := current
			next += numberInt
			if next > 99 {
				for next > 99 {
					next -= 100
					count += 1
				}

			}
			current = next

		}
		if strings.HasPrefix(instruction, "L") {
			split := strings.Split(instruction, "L")
			number := split[1]
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			next := current
			next -= numberInt

			if next < 0 {
				for next < 0 {
					next += 100
					count++
				}

			}
			current = next
		}

	}

	log.Print("Part 2 solution: ", count)

}
func part1(instructions []string) {
	current := 50
	count := 0
	for _, instruction := range instructions {

		if strings.HasPrefix(instruction, "R") {
			split := strings.Split(instruction, "R")
			number := split[1]
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			next := current
			next += numberInt
			if next >= 99 {
				for next >= 99 {
					next -= 99
					next -= 1
				}
				current = 0
				current = next
			} else {
				current = next
			}

		}
		if strings.HasPrefix(instruction, "L") {
			split := strings.Split(instruction, "L")
			number := split[1]
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			next := current
			next -= numberInt
			if next < 0 {
				for next < 0 {
					next += 99
					next += 1
				}
				current = 99
				current = next
			} else {
				current = next
			}
		}
		if current == 0 {
			count += 1
		}

	}
	log.Print("Part 1 solution: ", count)

}
func parseFile(filename string) (instructions []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	return instructions
}
