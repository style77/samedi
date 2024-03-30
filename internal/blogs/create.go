package blogs

import (
	"flag"
	"fmt"

	"github.com/style77/samedi/internal/database"
)

func CreateBlogCommand(args ...string) {
	if len(args) < 1 {
		fmt.Println("name is required")
		return
	}

	var description, logo, github, twitter, linkedin, language *string

	flags := flag.NewFlagSet("create_blog", flag.ExitOnError)
	flags.StringVar(description, "description", "", "description of the blog")
	flags.StringVar(logo, "logo", "", "logo of the blog")
	flags.StringVar(github, "github", "", "github of the blog")
	flags.StringVar(twitter, "twitter", "", "twitter of the blog")
	flags.StringVar(linkedin, "linkedin", "", "linkedin of the blog")
	flags.StringVar(language, "language", "en", "language of the blog")

	err := flags.Parse(args[1:])
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		return
	}

	createBlog(args[0], args[1], args[2], description, logo, github, twitter, linkedin, language)
}

func createBlog(name string, title string, author string, description *string, logo *string, github *string, twitter *string, linkedin *string, language *string) {
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
