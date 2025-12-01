package main

import (
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseFile(filename string) *os.File {
	f, err := os.Open(filename)
	Check(err)
	return f
}
