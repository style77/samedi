package app

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/russross/blackfriday/v2"
	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/posts"
)

func PostHandler(blog *blogs.Blog) http.HandlerFunc {
	tmpl := template.Must(template.New("post.html").Funcs(
		template.FuncMap{
			"Cmp":  func(i *string, j string) bool { return *i == j },
			"Ncmp": func(i *string, j string) bool { return *i != j },
		},
	).ParseFS(templates, "templates/post.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		postID := r.PathValue("id")
		id, err := strconv.Atoi(postID)
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		post, err := posts.GetPost(id, blog.ID)
		if err != nil || post == nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		htmlContent := blackfriday.Run([]byte(post.Content))
		currentYear := time.Now().Year()

		err = tmpl.Execute(w, struct {
			Blog            *blogs.Blog
			Post            *posts.Post
			HTMLPostContent template.HTML
			CurrentYear     int
		}{
			Blog:            blog,
			Post:            post,
			HTMLPostContent: template.HTML(htmlContent),
			CurrentYear:     currentYear,
		})
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
