package models

import (
	"time"

	"github.com/ugent-library/deliver/validate"
)

type Folder struct {
	ID        string    `json:"id,omitempty"`
	SpaceID   string    `json:"space_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// relations (can be empty)
	Size      int64   `json:"size"`
	FileCount int     `json:"file_count"`
	Space     *Space  `json:"space,omitempty"`
	Files     []*File `json:"files,omitempty"`
}

func (f *Folder) Validate() error {
	return validate.Validate(
		validate.NotEmpty("name", f.Name),
		validate.LengthIn("name", f.Name, 1, 100),
	)
}
