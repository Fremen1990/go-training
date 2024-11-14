package homework_command_line_program

import (
	"errors"
	"fmt"
)

func CommandLineTools(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: <command> [arguments...]\nAvailable commands: echo, cat, find, grep")
	}

	command := args[1]
	commandArgs := args[2:]

	switch command {
	case "echo":
		Echo(commandArgs)
	case "cat":
		Cat(commandArgs)
	case "find":
		if err := Find(commandArgs); err != nil {
			return fmt.Errorf("Find error: %v", err)
		}
	case "grep":
		if err := GrepCommand(commandArgs); err != nil {
			return fmt.Errorf("Grep error: %v", err)
		}
	default:
		return fmt.Errorf("Unknown command: %s\nAvailable commands: echo, cat, find, grep", command)
	}

	return nil
}
