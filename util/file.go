package util

import (
	"bufio"
	"os"
)

// Returns the file already opened.
func File(name string) *os.File {
	f, err := os.Open(name)
	if err != nil {
		panic("File not found")
	}
	return f
}

// Returns a slice with the entire file read.
func ReadString(name string) []string {
	var input []string
	file, err := os.Open(name)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
