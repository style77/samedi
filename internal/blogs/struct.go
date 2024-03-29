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
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	// Posts    []Post
}
