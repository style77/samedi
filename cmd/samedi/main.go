package main

import (
	"fmt"
	"os"

	"github.com/style77/samedi/internal/cli"
	"github.com/style77/samedi/internal/handlers"
)

func handleCommand(req cli.CommandRequest) {
	err := req.Handle()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handlers.Init()

	args := os.Args

	if len(args) < 2 {
		fmt.Println("no command provided")
	}

	req, err := cli.ParseCommand(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	handleCommand(req)
}
