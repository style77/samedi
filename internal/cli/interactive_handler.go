package cli

import (
	"fmt"
)

func RetrieveArgumentsInteractivally(arguments []CommandArgument) []string {
	args := make([]string, len(arguments))

	for i, arg := range arguments {
		fmt.Printf("Enter %s: ", arg.Name)
		val := &args[i]
		if arg.Required && val == nil {
			fmt.Println("This argument is required")
			i--
			continue
		}

		fmt.Scanln(val)
	}

	return args
}
