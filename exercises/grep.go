package exercises

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func Grep() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: grep <pattern> <path>")
	}

	pattern, err := regexp.Compile(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to compile regex pattern: %v", err)
	}

	path := os.Args[2]

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			log.Printf("Failed to open file %s: %v", path, err)
			return nil
		}

		defer func(file *os.File) {
			if file.Close() != nil {
				log.Fatal("Error closing file: ", file)
			}
		}(file)

		lineNumber := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if scannerErr := scanner.Err(); scannerErr != nil {
				log.Printf("Error reading file %s: %v", path, scannerErr)
			}
			lineNumber++
			if pattern.MatchString(line) {
				fmt.Printf("%s (line: %d): %s\n", path, lineNumber, line)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %s: %v", path, err)
	}
}
