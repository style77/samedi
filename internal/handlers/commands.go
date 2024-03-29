package handlers

import (
	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/cli"
	"github.com/style77/samedi/internal/web"
)

func Init() {
	cli.RegisterCommand("create_blog", blogs.CreateBlogCommand, []string{"createblog"}, []cli.CommandArgument{
		{Name: "name", Required: true, Position: 0},
		{Name: "description", Required: false, Position: 1},
	}, "Create a new blog")
	cli.RegisterCommand("serve", web.ServeCommand, []string{}, []cli.CommandArgument{
		{Name: "blog", Required: true, Position: 0},
		{Name: "host", Required: false, Position: 1, IsFlag: true, Default: "localhost"},
		{Name: "port", Required: false, Position: 2, IsFlag: true, Default: "8080"},
	}, "Start the web server")
	cli.RegisterCommand("help", HelpCommand, []string{}, []cli.CommandArgument{}, "Show this help message")
}
