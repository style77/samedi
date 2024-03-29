package posts

import "github.com/style77/samedi/internal/blogs"

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
	Blog      *blogs.Blog
}
