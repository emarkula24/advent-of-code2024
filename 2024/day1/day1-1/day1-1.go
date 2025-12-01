package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file := ParseFile("../day1-input")
	defer file.Close()
	var listLeft, listRight []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			continue
		}
		fields := strings.Fields(line)
		val1, err1 := strconv.Atoi(fields[0])
		val2, err2 := strconv.Atoi(fields[1])
		if err1 != nil || err2 != nil {
			log.Printf("Skipping line with non-integer values: %q", line)
			continue
		}

		listLeft = append(listLeft, val1)
		listRight = append(listRight, val2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	slices.Sort(listLeft)
	slices.Sort(listRight)

	var numberDistance int

	for i := range listLeft {
		if listLeft[i] > listRight[i] {
			numberDistance += listLeft[i] - listRight[i]
		} else if listRight[i] > listLeft[i] {
			numberDistance += listRight[i] - listLeft[i]
		} else {
			continue
		}
	}

	fmt.Println(numberDistance)

	duration := time.Since(start)
	fmt.Printf("Execution time: %d μs", duration.Microseconds())

}
