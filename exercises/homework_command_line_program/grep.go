package homework_command_line_program

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func searchFile(filePath string, pattern string) error {
	file, err := os.Open(filePath)

	if err != nil {
		return fmt.Errorf("could not open file %s: %v", file, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, pattern) {
			fmt.Printf("%s:%d %s\n", filePath, lineNumber, line)
			lineNumber++
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %s: %v", file, err)
	}
	return nil
}

func searchDirectory(dirPath string, pattern string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return searchFile(path, pattern)
		}
		return nil
	})
}

func GrepCommand(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("usage: grep <pattern> <file/directory> [additional files/directories...]")
	}

	pattern := args[1]
	paths := args[2:]

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("could not access %s: %v", path, err)
			continue
		}

		if info.IsDir() {
			if err := searchDirectory(path, pattern); err != nil {
				log.Printf("error searching directory %s: %v", path, err)
			}
		} else {
			if err := searchFile(path, pattern); err != nil {
				log.Printf("error searching file %s: %v", path, err)
			}
		}
	}
	return nil
}

//
//func Grep(args []string) error {
//	if len(args) < 3 {
//		return fmt.Errorf("usage: grep <pattern> <file/directory> [additional files/directories...]")
//	}
//
//	pattern := args[1]
//	paths := args[2:]
//
//	return grepLogic(pattern, paths)
//}
