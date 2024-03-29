package blogs

import (
	"fmt"

	"github.com/style77/samedi/internal/database"
)

func GetBlog(name string) (*Blog, error) {
	database, err := database.Init()
	if err != nil {
		return nil, err
	}

	var blog Blog
	err = database.QueryRow("SELECT id, name, description, title, author, logo, github, twitter, linkedin, language, created_at, updated_at FROM blogs WHERE name = ?", name).Scan(
		&blog.ID, &blog.Name, &blog.Description, &blog.Title, &blog.Author, &blog.Logo, &blog.Github, &blog.Twitter, &blog.Linkedin, &blog.Language, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		fmt.Println("No blog found with name:", name)
		return nil, err
	}

	return &blog, nil
}
