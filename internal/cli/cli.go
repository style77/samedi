package cli

import (
	"fmt"
)

type CommandRequest struct {
	Name string
	Args []string
}

type CommandMeta struct {
	Name        string
	Description string
	Aliases     []string
}

type CommandArgument struct {
	Name     string
	Required bool
	Position int
}

type Command struct {
	callback  func(...string)
	Meta      CommandMeta
	Arguments []CommandArgument
}

var Commands = make(map[string]Command)
var CommandAliases = make(map[string]string)

func ParseCommand(cmd []string) (CommandRequest, error) {
	name := cmd[1]
	args := make([]string, 0)

	if len(cmd) > 2 {
		args = cmd[2:]
	}

	if alias, ok := CommandAliases[name]; ok {
		name = alias
	}

	if _, ok := Commands[name]; !ok {
		return CommandRequest{}, fmt.Errorf("unknown command: %s", name)
	}
	return CommandRequest{
		Name: name,
		Args: args,
	}, nil
}

func (c CommandRequest) Handle() error {
	cmd, ok := Commands[c.Name]

	if !ok {
		return fmt.Errorf("unknown command: %s", c.Name)
	}

	command_args := c.Args

	if len(command_args) < len(cmd.Arguments) {
		command_args = RetrieveArgumentsInteractivally(cmd.Arguments[len(c.Args):])
	}

	cmd.callback(command_args...)
	return nil
}

// RegisterCommand registers a command with a function to be called when the command is executed.
func RegisterCommand(cmd string, f func(...string), aliases []string, args []CommandArgument, description string) {
	if _, ok := Commands[cmd]; ok {
		panic(fmt.Sprintf("command %s already registered", cmd))
	}

	Commands[cmd] = Command{
		callback:  f,
		Meta:      CommandMeta{Name: cmd, Description: description, Aliases: aliases},
		Arguments: args,
	}

	if aliases == nil {
		return
	}

	for _, alias := range aliases {
		if _, ok := Commands[alias]; ok {
			panic(fmt.Sprintf("command %s already registered", alias))
		}

		CommandAliases[alias] = cmd
	}
}