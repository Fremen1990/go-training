package homework_command_line_program

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func checkType(fileInfo os.FileInfo, fileType string) bool {
	switch fileType {
	case "file":
		return fileInfo.Mode().IsRegular()
	case "directory":
		return fileInfo.IsDir()
	case "link":
		return fileInfo.Mode()&os.ModeSymlink != 0
	default:
		return false
	}
}

func findLogic(path string, pattern string, fileType string) error {
	return filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matchesPattern, _ := filepath.Match(pattern, info.Name())
		matchesType := checkType(info, fileType)

		if matchesPattern && matchesType {
			fmt.Println(currentPath)
		}
		return nil
	})
}

func Find(args []string) error {
	if len(args) < 3 {
		return errors.New("usage: find <path> <pattern> <type> (type should be 'file', 'directory', or 'link')")
	}

	startPath := args[0]
	pattern := args[1]
	fileType := args[2]

	if fileType != "file" && fileType != "directory" && fileType != "link" {
		return fmt.Errorf("invalid type: %s. Allowed types are 'file', 'directory' or 'link'", fileType)
	}
	return findLogic(startPath, pattern, fileType)
}
