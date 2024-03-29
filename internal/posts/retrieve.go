package posts

import "github.com/style77/samedi/internal/database"

func GetPost(id int, blogId int) (*Post, error) {
	database, err := database.Init()
	if err != nil {
		return nil, err
	}

	var post Post
	err = database.QueryRow("SELECT id, title, content, created_at, updated_at FROM posts WHERE id = ? AND blog_id = ?", id, blogId).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Blog.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
