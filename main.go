package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/adam-hanna/arrayOperations"
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

func readFiles(fileNames []string) ([][]string, error) {
	var files [][]string
	for _, fileName := range fileNames {
		fileLines, err := readLines(fileName)
		if err != nil {
			return nil, err
		}
		files = append(files, fileLines)
	}
	return files, nil
}

func main() {
	intersectPtr := flag.Bool("i", false, "Intersect two files")
	flag.Parse()
	args := flag.Args()

	files, err := readFiles(args)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	if *intersectPtr {
		slice := arrayOperations.IntersectString(files...)
		for _, line := range slice {
			fmt.Println(line)
		}
	} else {
		slice := arrayOperations.DifferenceString(files...)
		for _, line := range slice {
			fmt.Println(line)
		}
	}
	os.Exit(0)
}
