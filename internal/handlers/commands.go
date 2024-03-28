package handlers

import (
	"github.com/style77/samedi/internal/cli"
)

func Init() {
	cli.RegisterCommand("new_post", NewPostCommand, []string{"newpost", "new", "post"}, []cli.CommandArgument{
		{
			Name:     "title",
			Required: true,
			Position: 0,
		},
	}, "Create a new post")
	cli.RegisterCommand("help", HelpCommand, []string{}, []cli.CommandArgument{}, "Show this help message")
}
