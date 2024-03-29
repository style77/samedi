package blogs

import (
	"time"

	"github.com/style77/samedi/internal/database"
)

type Blog struct {
	Database    *database.Database
	ID          int
	Name        string
	Description *string
	Title       string
	Author      string
	Logo        *string
	Github      *string
	Twitter     *string
	Linkedin    *string
	Language    string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	// Posts    []Post
}
