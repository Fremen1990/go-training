package homework_command_line_program

import (
	"fmt"
)

func Echo(args []string) {
	fmt.Println("Echoing:", args)
	for _, arg := range args {
		fmt.Println(arg)
	}
}
