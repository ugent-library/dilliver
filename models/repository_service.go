package models

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/ugent-library/deliver/ent"
	"github.com/ugent-library/deliver/ent/file"
	"github.com/ugent-library/deliver/ent/folder"
	entmigrate "github.com/ugent-library/deliver/ent/migrate"
	"github.com/ugent-library/deliver/ent/space"
	"github.com/ugent-library/deliver/ent/user"
	"github.com/ugent-library/deliver/validate"
)

var ErrNotFound = errors.New("not found")

type RepositoryConfig struct {
	DB string
}

type RepositoryService interface {
	UserByRememberToken(context.Context, string) (*User, error)
	CreateOrUpdateUser(context.Context, *User) error
	RenewUserRememberToken(context.Context, string) error
	Spaces(context.Context) ([]*Space, error)
	SpacesByUsername(context.Context, string) ([]*Space, error)
	SpaceByID(context.Context, string) (*Space, error)
	SpaceByName(context.Context, string) (*Space, error)
	CreateSpace(context.Context, *Space) error
	UpdateSpace(context.Context, *Space) error
	FolderByID(context.Context, string) (*Folder, error)
	CreateFolder(context.Context, *Folder) error
	UpdateFolder(context.Context, *Folder) error
	DeleteFolder(context.Context, string) error
	DeleteExpiredFolders(context.Context) error
	FileByID(context.Context, string) (*File, error)
	CreateFile(context.Context, *File) error
	DeleteFile(context.Context, string) error
	AddFileDownload(context.Context, string) error
}

func NewRepositoryService(c RepositoryConfig) (RepositoryService, error) {
	db, err := sql.Open("pgx", c.DB)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(entdialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

	err = client.Schema.Create(context.TODO(),
		entmigrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

	return &repositoryService{
		db: client,
	}, nil
}

type repositoryService struct {
	db *ent.Client
}

func (r *repositoryService) UserByRememberToken(ctx context.Context, token string) (*User, error) {
	row, err := r.db.User.Query().
		Where(user.RememberTokenEQ(token)).
		First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rowToUser(row), nil

}

// TODO rewrite this when ent supports the Save method on Update; until then
// we have to do an extra select
// https://github.com/ent/ent/issues/2600
func (r *repositoryService) CreateOrUpdateUser(ctx context.Context, u *User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	token, err := NewRememberToken()
	if err != nil {
		return err
	}
	id, err := r.db.User.Create().
		SetUsername(u.Username).
		SetName(u.Name).
		SetEmail(u.Email).
		SetRememberToken(token).
		OnConflict(
			entsql.ConflictColumns(user.FieldUsername),
		).
		Update(func(u *ent.UserUpsert) {
			u.UpdateName()
			u.UpdateEmail()
		}).ID(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return validate.NewErrors(validate.ErrNotUnique("username"))
		}
		return err
	}
	row, err := r.db.User.Get(ctx, id)
	if err != nil {
		return err
	}
	*u = *rowToUser(row)
	return nil
}

func (r *repositoryService) RenewUserRememberToken(ctx context.Context, id string) error {
	newToken, err := NewRememberToken()
	if err != nil {
		return err
	}
	err = r.db.User.
		UpdateOneID(id).
		SetRememberToken(newToken).
		Exec(ctx)
	return err
}

func (r *repositoryService) Spaces(ctx context.Context) ([]*Space, error) {
	rows, err := r.db.Space.Query().
		Order(ent.Asc(space.FieldName)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	spaces := make([]*Space, len(rows))
	for i, row := range rows {
		spaces[i] = rowToSpace(row)
	}
	return spaces, nil
}

func (r *repositoryService) SpacesByUsername(ctx context.Context, username string) ([]*Space, error) {
	rows, err := r.db.Space.Query().
		Where(func(s *entsql.Selector) {
			s.Where(sqljson.ValueContains(space.FieldAdmins, username))
		}).
		Order(ent.Asc(space.FieldName)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	spaces := make([]*Space, len(rows))
	for i, row := range rows {
		spaces[i] = rowToSpace(row)
	}
	return spaces, nil
}

func (r *repositoryService) SpaceByID(ctx context.Context, id string) (*Space, error) {
	row, err := r.db.Space.Query().
		Where(space.IDEQ(id)).
		WithFolders(func(q *ent.FolderQuery) {
			q.Order(ent.Asc(folder.FieldExpiresAt))
			q.WithFiles(func(q *ent.FileQuery) {
				// TODO why does this give the error
				// unexpected foreign-key "folder_id" returned  for node
				// q.Select(file.FieldSize)
			})
		}).
		First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rowToSpace(row), nil
}

func (r *repositoryService) SpaceByName(ctx context.Context, name string) (*Space, error) {
	row, err := r.db.Space.Query().
		Where(space.NameEQ(name)).
		WithFolders(func(q *ent.FolderQuery) {
			q.Order(ent.Asc(folder.FieldExpiresAt))
			q.WithFiles(func(q *ent.FileQuery) {
				// TODO why does this give the error
				// unexpected foreign-key "folder_id" returned  for node
				// q.Select(file.FieldSize)
			})
		}).
		First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rowToSpace(row), nil
}

func (r *repositoryService) CreateSpace(ctx context.Context, s *Space) error {
	if err := s.Validate(); err != nil {
		return err
	}
	row, err := r.db.Space.Create().
		SetName(s.Name).
		SetAdmins(s.Admins).
		Save(ctx)
	if err != nil {
		return err
	}
	*s = *rowToSpace(row)
	return nil
}

func (r *repositoryService) UpdateSpace(ctx context.Context, s *Space) error {
	if err := s.Validate(); err != nil {
		return err
	}
	row, err := r.db.Space.UpdateOneID(s.ID).
		SetAdmins(s.Admins).
		Save(ctx)
	if err != nil {
		return err
	}
	*s = *rowToSpace(row)
	return nil
}

func (r *repositoryService) FolderByID(ctx context.Context, id string) (*Folder, error) {
	row, err := r.db.Folder.Query().
		Where(folder.IDEQ(id)).
		WithSpace().
		WithFiles(func(q *ent.FileQuery) {
			q.Order(ent.Asc(file.FieldName))
		}).
		First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rowToFolder(row), nil
}

func (r *repositoryService) CreateFolder(ctx context.Context, f *Folder) error {
	if err := f.Validate(); err != nil {
		return err
	}
	row, err := r.db.Folder.Create().
		SetSpaceID(f.SpaceID).
		SetName(f.Name).
		SetExpiresAt(f.ExpiresAt).
		Save(ctx)
	if err != nil {
		// TODO does ent support unwrapping sql errors?
		// https://stackoverflow.com/questions/70859712/how-do-you-handle-database-errors-in-go-without-getting-coupled-to-the-sql-drive
		// https://github.com/ent/ent/issues/2328
		// see also UpdateFolder and CreateUser
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return validate.NewErrors(validate.ErrNotUnique("name"))
		}
		return err
	}
	*f = *rowToFolder(row)
	return nil
}

func (r *repositoryService) UpdateFolder(ctx context.Context, f *Folder) error {
	if err := f.Validate(); err != nil {
		return err
	}
	row, err := r.db.Folder.UpdateOneID(f.ID).
		SetName(f.Name).
		Save(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return validate.NewErrors(validate.ErrNotUnique("name"))
		}
		return err
	}
	*f = *rowToFolder(row)
	return nil
}

func (r *repositoryService) DeleteFolder(ctx context.Context, folderID string) error {
	err := r.db.Folder.
		DeleteOneID(folderID).
		Exec(ctx)
	return err
}

func (r *repositoryService) DeleteExpiredFolders(ctx context.Context) error {
	_, err := r.db.Folder.
		Delete().
		Where(folder.ExpiresAtLT(time.Now())).
		Exec(ctx)
	return err
}

func (r *repositoryService) CreateFile(ctx context.Context, f *File) error {
	if err := f.Validate(); err != nil {
		return err
	}
	row, err := r.db.File.Create().
		SetID(f.ID).
		SetFolderID(f.FolderID).
		SetMd5(f.MD5).
		SetName(f.Name).
		SetContentType(f.ContentType).
		SetSize(f.Size).
		Save(ctx)
	if err != nil {
		return err
	}
	*f = *rowToFile(row)
	return nil
}

func (r *repositoryService) FileByID(ctx context.Context, id string) (*File, error) {
	row, err := r.db.File.Query().
		Where(file.IDEQ(id)).
		WithFolder(func(q *ent.FolderQuery) {
			q.WithSpace()
		}).
		First(ctx)
	if err != nil {
		var e *ent.NotFoundError
		if errors.As(err, &e) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return rowToFile(row), nil
}

func (r *repositoryService) DeleteFile(ctx context.Context, id string) error {
	err := r.db.File.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

func (r *repositoryService) AddFileDownload(ctx context.Context, id string) error {
	err := r.db.File.
		UpdateOneID(id).
		AddDownloads(1).
		Exec(ctx)
	return err
}

func rowToUser(row *ent.User) *User {
	u := &User{
		ID:            row.ID,
		Username:      row.Username,
		Name:          row.Name,
		Email:         row.Email,
		RememberToken: row.RememberToken,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
	return u
}

func rowToSpace(row *ent.Space) *Space {
	s := &Space{
		ID:        row.ID,
		Name:      row.Name,
		Admins:    row.Admins,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
	if row.Edges.Folders != nil {
		s.Folders = make([]*Folder, len(row.Edges.Folders))
		for i, r := range row.Edges.Folders {
			f := &Folder{
				ID:        r.ID,
				SpaceID:   r.SpaceID,
				Name:      r.Name,
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
				ExpiresAt: r.ExpiresAt,
			}
			if r.Edges.Files != nil {
				f.FileCount = len(r.Edges.Files)
				for _, r := range r.Edges.Files {
					f.Size += r.Size
				}
			}

			s.Folders[i] = f
		}
	}
	return s
}

func rowToFolder(row *ent.Folder) *Folder {
	f := &Folder{
		ID:        row.ID,
		SpaceID:   row.SpaceID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		ExpiresAt: row.ExpiresAt,
	}
	if row.Edges.Space != nil {
		f.Space = rowToSpace(row.Edges.Space)
	}
	if row.Edges.Files != nil {
		f.FileCount = len(row.Edges.Files)
		f.Files = make([]*File, len(row.Edges.Files))
		for i, r := range row.Edges.Files {
			ff := rowToFile(r)
			f.Size += ff.Size
			f.Files[i] = ff
		}
	}
	return f
}

func rowToFile(row *ent.File) *File {
	f := &File{
		ID:          row.ID,
		FolderID:    row.FolderID,
		MD5:         row.Md5,
		Name:        row.Name,
		Size:        row.Size,
		ContentType: row.ContentType,
		Downloads:   row.Downloads,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
	if row.Edges.Folder != nil {
		f.Folder = rowToFolder(row.Edges.Folder)
	}
	return f
}
