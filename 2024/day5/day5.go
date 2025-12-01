package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rules, pages := parseFile("day5.txt")
	// fmt.Println(rules)
	// fmt.Println(pages)
	result1 := part1(rules, pages)
	result2 := part2(rules, pages)
	fmt.Println(result1)
	fmt.Println(result2)
}

func part1(rules [][2]int, pages [][]int) int {
	ruleMap := buildRuleMap(rules)
	sum := 0
	for _, page := range pages {
		if isPageValid(ruleMap, page) {
			sum += middleValue(page)
		}

	}
	return sum
}
func part2(rules [][2]int, pages [][]int) int {
	ruleMap := buildRuleMap(rules)
	sum := 0
	for _, page := range pages {
		if !isPageValid(ruleMap, page) {
			fixed := fixPage(ruleMap, page)
			sum += middleValue(fixed)
		}

	}
	return sum
}
func fixPage(ruleMap map[int][]int, page []int) []int {
	sort.Slice(page, func(i, j int) bool {
		a, b := page[i], page[j]

		if slices.Contains(ruleMap[a], b) {
			return true
		}
		if slices.Contains(ruleMap[b], a) {
			return false
		}

		return a < b
	})
	return page
}
func isPageValid(ruleMap map[int][]int, page []int) bool {
	pos := make(map[int]int)
	for i, num := range page {
		pos[num] = i
	}

	for a, bs := range ruleMap {
		for _, b := range bs {
			posA, okA := pos[a]
			posB, okB := pos[b]
			if okA && okB && posA > posB {
				return false
			}
		}
	}
	return true
}

func middleValue(page []int) int {
	return page[len(page)/2]
}

func buildRuleMap(rules [][2]int) map[int][]int {
	ruleMap := make(map[int][]int)
	for _, r := range rules {
		ruleMap[r[0]] = append(ruleMap[r[0]], r[1])
	}
	return ruleMap
}

func parseFile(fileName string) (rules [][2]int, pages [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.Contains(line, "|"):
			parts := strings.Split(line, "|")

			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			rules = append(rules, [2]int{a, b})

		case strings.Contains(line, ","):
			parts := strings.Split(line, ",")
			var row []int
			for _, p := range parts {
				val, _ := strconv.Atoi(strings.TrimSpace(p))
				row = append(row, val)
			}
			pages = append(pages, row)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return rules, pages
}
