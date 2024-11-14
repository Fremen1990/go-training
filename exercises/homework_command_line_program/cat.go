package homework_command_line_program

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// Options to read file
// 1. os.ReadFile
// 2. os.Open
// 3. use 'bufio' to scan text line by line

func OpenAndReadFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		bytesRead, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}

		text := string(buffer[:bytesRead])

		if bytesRead > 0 {
			fmt.Printf("%s", text)
		}
	}
}

func OpenAndScanFile(fileName string, flag string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()

		switch flag {
		case "":
			fmt.Printf("%s", line)
		case "-n":
			fmt.Printf("%d. %s\n", lineNumber, line)
			lineNumber++
		case "-nb":
			if line == "" {
				fmt.Printf("%s\n", line)
			} else {
				fmt.Printf("%d. %s\n", lineNumber, line)
				lineNumber++
			}
		default:
			fmt.Println("Flag not recognized, please use --help to see documentation :D ")
			return
		}

	}
}

func Cat(args []string) {
	//args := os.Args[1:]
	if len(args) > 0 {
		//fileNameFromCommandLine := strings.Join(args, " ")
		fileNameFromCommandLine := args[0]
		//OpenAndReadFile(fileNameFromCommandLine)
		if len(args) == 1 {
			OpenAndScanFile(fileNameFromCommandLine, "")
		} else if len(args) == 2 {
			OpenAndScanFile(fileNameFromCommandLine, args[1])
		}
	}

}
