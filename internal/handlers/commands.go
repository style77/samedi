package handlers

import (
	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/cli"
	"github.com/style77/samedi/internal/posts"
	"github.com/style77/samedi/internal/web"
)

func Init() {
	cli.RegisterCommand("create_blog", blogs.CreateBlogCommand, []string{"createblog"}, []cli.CommandArgument{
		{Name: "name", Required: true, Position: 0},
		{Name: "title", Required: true, Position: 1},
		{Name: "author", Required: true, Position: 2},
		{Name: "description", Required: false, Position: 3},
		{Name: "logo", Required: false, Position: 4, IsFlag: true},
		{Name: "github", Required: false, Position: 5, IsFlag: true},
		{Name: "twitter", Required: false, Position: 6, IsFlag: true},
		{Name: "linkedin", Required: false, Position: 7, IsFlag: true},
		{Name: "language", Required: false, Position: 8, IsFlag: true, Default: "en"},
	}, "Create a new blog")
	cli.RegisterCommand("create_post", posts.CreatePostCommand, []string{"createpost"}, []cli.CommandArgument{
		{Name: "blog", Required: true, Position: 0},
		{Name: "title", Required: true, Position: 1},
	}, "Create a new post")
	cli.RegisterCommand("serve", web.ServeCommand, []string{}, []cli.CommandArgument{
		{Name: "blog", Required: true, Position: 0},
		{Name: "host", Required: false, Position: 1, IsFlag: true, Default: "localhost"},
		{Name: "port", Required: false, Position: 2, IsFlag: true, Default: "8080"},
	}, "Start the web server")
	cli.RegisterCommand("help", HelpCommand, []string{}, []cli.CommandArgument{}, "Show this help message")
}
