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

	createBlog(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
}

func createBlog(name string, title string, author string, description string, logo string, github string, twitter string, linkedin string, language string) {
	db, err := database.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	blog, _ := GetBlog(name)
	if blog != nil {
		fmt.Println(fmt.Sprintf("Blog %s already exists", name))
		return
	}

	_, err = db.Exec(
		"INSERT INTO blogs (name, description, title, author, logo, github, twitter, linkedin, language) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		name, description, title, author, logo, github, twitter, linkedin, language)
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
