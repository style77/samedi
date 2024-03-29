package cli

import (
	"bufio"
	"fmt"
	"os"
)

func RetrieveArgumentsInteractively(arguments []CommandArgument) []string {
	args := make([]string, len(arguments))

	scanner := bufio.NewScanner(os.Stdin)

	for i, arg := range arguments {
		if arg.IsFlag || !arg.Required {
			args[i] = arg.Default
			continue
		}
		fmt.Printf("Enter %s: ", arg.Name)
		scanner.Scan()
		val := scanner.Text()
		if arg.Required && val == "" {
			fmt.Println("This argument is required")
			i--
			continue
		}

		args[i] = val
	}

	return args
}
