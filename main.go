package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	ao "github.com/adam-hanna/arrayOperations"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("a.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	lines2, err := readLines("b.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	z, ok := ao.Difference(lines, lines2)
	if !ok {
		fmt.Println("Cannot find intersect")
	}

	slice, ok := z.Interface().([]string)
	if !ok {
		fmt.Println("Cannot convert to slice")
	}

	for _, line := range slice {
		fmt.Println(line)
	}
}
