package app

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/posts"
)

//go:embed templates/*
var templates embed.FS

func IndexHandler(blog *blogs.Blog) http.HandlerFunc {
	tmpl, err := template.ParseFS(templates, "templates/index.html")
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		postsData, err := posts.GetPosts(blog.ID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Posts not found", http.StatusNotFound)
			return
		}

		err = tmpl.Execute(w, struct {
			Blog  *blogs.Blog
			Posts []posts.Post
		}{
			Blog:  blog,
			Posts: postsData,
		})
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
