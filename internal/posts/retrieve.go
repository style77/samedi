package posts

import (
	"github.com/style77/samedi/internal/database"
)

func GetPost(id int, blogId int) (*Post, error) {
	database, err := database.Init()
	if err != nil {
		return nil, err
	}

	var post Post
	err = database.QueryRow("SELECT id, title, content, created_at, updated_at FROM posts WHERE id = ? AND blog_id = ?", id, blogId).Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPosts(blogId int) ([]Post, error) {
	database, err := database.Init()
	if err != nil {
		return nil, err
	}

	rows, err := database.Query("SELECT id, title, content, created_at, updated_at FROM posts WHERE blog_id = ? ORDER BY created_at desc", blogId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
