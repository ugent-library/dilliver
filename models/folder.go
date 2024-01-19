package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mozillazg/go-unidecode"
	"github.com/ugent-library/okay"
)

var reSlug = regexp.MustCompile("[^a-zA-Z0-9-]+")

type Folder struct {
	ID        string    `json:"id,omitempty"`
	SpaceID   string    `json:"space_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// relations (can be empty)
	Space *Space  `json:"space,omitempty"`
	Files []*File `json:"files,omitempty"`
}

func (f *Folder) TotalSize() (n int64) {
	for _, file := range f.Files {
		n += file.Size
	}
	return
}

func (f *Folder) TotalDownloads() (n int64) {
	for _, file := range f.Files {
		n += file.Downloads
	}
	return
}

func (f *Folder) Slug() string {
	return strings.Trim(reSlug.ReplaceAllString(unidecode.Unidecode(f.Name), "-"), "-")
}

func (f *Folder) Validate() error {
	return okay.Validate(
		okay.NotEmpty("name", f.Name),
		okay.LengthBetween("name", f.Name, 1, 100),
	)
}

func (f *Folder) Fake(faker *gofakeit.Faker) (any, error) {
	created := gofakeit.PastDate()
	return Folder{
		Name:      fmt.Sprintf("%d", gofakeit.Number(1234567, 9123456)),
		CreatedAt: created,
		UpdatedAt: gofakeit.DateRange(created, time.Now()),
		ExpiresAt: created.AddDate(0, 1, 0),
	}, nil
}
