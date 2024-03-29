package posts

import (
	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/database"
	"github.com/style77/samedi/internal/helpers"
)

func CreatePostCommand(args ...string) {
	blogName := args[0]
	title := args[1]

	content, err := helpers.GetTextInput(helpers.GetAvailableEditor())
	if err != nil {
		return
	}

	blog, err := blogs.GetBlog(blogName)
	if err != nil {
		return
	}
	createPost(title, content, blog)
}

func createPost(title string, content string, blog *blogs.Blog) {
	database, err := database.Init()
	if err != nil {
		return
	}

	_, err = database.Exec("INSERT INTO posts (title, content, blog_id) VALUES (?, ?, ?)", title, content, blog.ID)
	if err != nil {
		return
	}

	return
}
