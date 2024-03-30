package app

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/internal/posts"
)

//go:embed templates/*
var templates embed.FS

func IndexHandler(blog *blogs.Blog) http.HandlerFunc {
	tmpl := template.Must(template.New("index.html").Funcs(
		template.FuncMap{
			"Cmp":  func(i *string, j string) bool { return *i == j },
			"Ncmp": func(i *string, j string) bool { return *i != j },
		},
	).ParseFS(templates, "templates/index.html"))

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

		currentYear := time.Now().Year()

		err = tmpl.Execute(w, struct {
			Blog        *blogs.Blog
			Posts       []posts.Post
			CurrentYear int
		}{
			Blog:        blog,
			Posts:       postsData,
			CurrentYear: currentYear,
		})
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
