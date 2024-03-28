package handlers

import "fmt"

func NewPostCommand(args ...string) {
	title := args[0]

	NewPost(title)
}

func NewPost(title string) {
	fmt.Println(title)
}
