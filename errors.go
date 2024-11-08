package main

import (
	"errors"
	"fmt"
)

func errorsExamples() {
	result, err := safeDiv(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	if _, err := readText("/Users/file.txtx"); err != nil {
		if errors.Is(err, fileNotFound) {
			fmt.Println("File not found")
		}
	}

	// panic("Fatal error") // błąd krytyczny, przerwanie działania programu
}

func safeDiv(value, divident float64) (float64, error) {
	if divident == 0 {
		return 0.0, errors.New("Division by zero")
	}
	return value / divident, nil
}

var fileNotFound = fmt.Errorf("file not found")
var readError = fmt.Errorf("read error")

func readText(path string) (string, error) {
	if path == "" {
		return "", readError
	}
	// ... read
	return "lines", nil
}

type appError struct {
	code         int
	descriptions string
}

func (e *appError) Error() string {
	return e.descriptions
}
