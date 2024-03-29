package web

import (
	"fmt"
	"net/http"

	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/web/app"
)

type Server struct {
	Host string
	Port int
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) Serve(blog *blogs.Blog) {
	handler := http.NewServeMux()
	handler.HandleFunc("GET /", app.IndexHandler(blog))
	handler.HandleFunc("GET /post/{id}/", app.PostHandler(blog))
	handler.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/app/static"))))

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), handler)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Server listening on http://%s:%d\n", s.Host, s.Port)

	select {}
}
