package web

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/style77/samedi/internal/blogs"
)

type Server struct {
	Host string
	Port int
}

func ServeCommand(args ...string) {
	var host string
	var port int

	flags := flag.NewFlagSet("server", flag.ExitOnError)
	flags.StringVar(&host, "host", "localhost", "host to serve on")
	flags.IntVar(&port, "port", 8080, "port to serve on")
	err := flags.Parse(args[1:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	blog, err := blogs.GetBlog(args[0])
	if err != nil {
		return
	}

	fmt.Println(args[1:], host, port)

	server := NewServer(host, port)

	server.Serve(blog)
}

func NewServer(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (s *Server) Serve(blog *blogs.Blog) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), handler)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Printf("Server listening on http://%s:%d\n", s.Host, s.Port)

	select {}
}
