package web

import (
	"flag"
	"fmt"
	"net"

	"github.com/style77/samedi/internal/blogs"
	"github.com/style77/samedi/web"
)

func checkIfPortIsAvailable(host string, port int) bool {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: net.ParseIP(host), Port: port})
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func ServeCommand(args ...string) {
	var host string
	var port int

	flags := flag.NewFlagSet("server", flag.ExitOnError)
	flags.StringVar(&host, "host", "localhost", "host to serve on")
	flags.IntVar(&port, "port", 5400, "port to serve on")
	err := flags.Parse(args[1:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	blog, err := blogs.GetBlog(args[0])
	if err != nil {
		return
	}

	for checkIfPortIsAvailable(host, port) {
		port++
	}

	if port != 5400 {
		fmt.Println("Port 5400 is already in use, using port", port)
	}

	server := web.NewServer(host, port)

	server.Serve(blog)
}
