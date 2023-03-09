// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "md5", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt64},
		{Name: "content_type", Type: field.TypeString},
		{Name: "downloads", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "folder_id", Type: field.TypeString},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_folders_files",
				Columns:    []*schema.Column{FilesColumns[8]},
				RefColumns: []*schema.Column{FoldersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FoldersColumns holds the columns for the "folders" table.
	FoldersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "space_id", Type: field.TypeString},
	}
	// FoldersTable holds the schema information for the "folders" table.
	FoldersTable = &schema.Table{
		Name:       "folders",
		Columns:    FoldersColumns,
		PrimaryKey: []*schema.Column{FoldersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "folders_spaces_folders",
				Columns:    []*schema.Column{FoldersColumns[5]},
				RefColumns: []*schema.Column{SpacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "folder_space_id_name",
				Unique:  true,
				Columns: []*schema.Column{FoldersColumns[5], FoldersColumns[1]},
			},
		},
	}
	// SpacesColumns holds the columns for the "spaces" table.
	SpacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "admins", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// SpacesTable holds the schema information for the "spaces" table.
	SpacesTable = &schema.Table{
		Name:       "spaces",
		Columns:    SpacesColumns,
		PrimaryKey: []*schema.Column{SpacesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "remember_token", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FilesTable,
		FoldersTable,
		SpacesTable,
		UsersTable,
	}
)

func init() {
	FilesTable.ForeignKeys[0].RefTable = FoldersTable
	FoldersTable.ForeignKeys[0].RefTable = SpacesTable
}
