package blogs

import (
	"fmt"

	"github.com/style77/samedi/internal/database"
)

func CreateBlogCommand(args ...string) {
	if len(args) < 1 {
		fmt.Println("name is required")
		return
	}

	name := args[0]
	description := ""
	if len(args) > 1 {
		description = args[1]
	}

	createBlog(name, description)
}

func createBlog(name string, description string) {
	db, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Exec("INSERT INTO blogs (name, description) VALUES (?, ?)", name, description)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Blog %s created successfully", name))

	err = db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
