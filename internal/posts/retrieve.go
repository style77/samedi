package posts

import "github.com/style77/samedi/internal/database"

func GetPost(id int) (*Post, error) {
	database, err := database.Init()
	if err != nil {
		return nil, err
	}

	var post Post
	err = database.QueryRow("SELECT id, title, content, created_at, updated_at, blog_id FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Blog.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
