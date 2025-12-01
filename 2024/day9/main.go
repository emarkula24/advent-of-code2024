package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	s := parseFile("day9.txt")
	part1(s)
}

type File struct {
	Id     int
	Blocks int
	Number int
}

func part1(line string) {
	files := buildFileSlice(line)
	spread := spreadString(files)
	sorted := sort(spread)
	calculateChecksum(sorted)
}
func sort(line []int) []int {
	left, right := 0, len(line)-1
	for left < right {
		for left < right && line[left] != -1 {
			left++
		}
		for left < right && line[right] == -1 {
			right--
		}
		if left < right {
			line[left], line[right] = line[right], line[left]
			left++
			right--
		}
	}
	return line
}
func calculateChecksum(line []int) {
	checksum := 0
	for i, val := range line {
		if val == -1 {
			continue
		}
		checksum += i * val
	}
	fmt.Println("Checksum:", checksum)
}
func spreadString(files []File) []int {
	var completeLine []int
	for _, f := range files {
		for i := 0; i < f.Number; i++ {
			completeLine = append(completeLine, f.Id)
		}
		for i := 0; i < f.Blocks; i++ {
			completeLine = append(completeLine, -1) // represent dot
		}
	}
	return completeLine
}
func buildFileSlice(line string) []File {
	files := make([]File, 0, 100)
	var counter int = 0
	for i := 0; i < len(line); i += 2 {
		number := int(line[i] - '0')
		blocks := 0
		if i+1 < len(line) {
			blocks = int(line[i+1] - '0')
		}
		f := File{
			Id:     counter,
			Number: number,
			Blocks: blocks,
		}
		files = append(files, f)
		counter++

	}
	return files
}
func parseFile(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line
}
