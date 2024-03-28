package handlers

import (
	"fmt"

	"github.com/style77/samedi/internal/cli"
)

func HelpCommand(args ...string) {
	fmt.Println(GenerateHelp())
}

func GenerateUsage(cmd string) string {
	if _, ok := cli.Commands[cmd]; !ok {
		return fmt.Sprintf("Unknown command: %s", cmd)
	}

	message := fmt.Sprintf("Usage: %s", cmd)
	for _, arg := range cli.Commands[cmd].Arguments {
		message += fmt.Sprintf(" <%s>", arg.Name)
	}

	return message
}

func GenerateHelp() string {
	message := "Usage: samedi <command> [args]\n\n"

	message += "Commands:\n"

	for name, cmd := range cli.Commands {
		var cmd_args string
		for _, arg := range cmd.Arguments {
			cmd_args += fmt.Sprintf(" <%s>", arg.Name)
		}

		message += fmt.Sprintf("  %s%s - %s\n", name, cmd_args, cmd.Meta.Description)
	}

	return message
}
