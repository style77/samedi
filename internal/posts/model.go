package posts

import (
	"time"

	"github.com/style77/samedi/internal/blogs"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt *time.Time
	Blog      *blogs.Blog
}
