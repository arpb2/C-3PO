// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// CredentialsColumns holds the columns for the "credentials" table.
	CredentialsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "salt", Type: field.TypeBytes},
		{Name: "password_hash", Type: field.TypeBytes},
		{Name: "user_credentials", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CredentialsTable holds the schema information for the "credentials" table.
	CredentialsTable = &schema.Table{
		Name:       "credentials",
		Columns:    CredentialsColumns,
		PrimaryKey: []*schema.Column{CredentialsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "credentials_users_credentials",
				Columns: []*schema.Column{CredentialsColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "credential_user_credentials",
				Unique:  true,
				Columns: []*schema.Column{CredentialsColumns[3]},
			},
		},
	}
	// LevelsColumns holds the columns for the "levels" table.
	LevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// LevelsTable holds the schema information for the "levels" table.
	LevelsTable = &schema.Table{
		Name:        "levels",
		Columns:     LevelsColumns,
		PrimaryKey:  []*schema.Column{LevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"teacher", "student"}},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "surname", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UserLevelsColumns holds the columns for the "user_levels" table.
	UserLevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString},
		{Name: "workspace", Type: field.TypeString},
		{Name: "user_levels", Type: field.TypeInt, Nullable: true},
		{Name: "user_level_level", Type: field.TypeInt, Nullable: true},
	}
	// UserLevelsTable holds the schema information for the "user_levels" table.
	UserLevelsTable = &schema.Table{
		Name:       "user_levels",
		Columns:    UserLevelsColumns,
		PrimaryKey: []*schema.Column{UserLevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_levels_users_levels",
				Columns: []*schema.Column{UserLevelsColumns[5]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "user_levels_levels_level",
				Columns: []*schema.Column{UserLevelsColumns[6]},

				RefColumns: []*schema.Column{LevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "userlevel_user_levels_user_level_level",
				Unique:  true,
				Columns: []*schema.Column{UserLevelsColumns[5], UserLevelsColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CredentialsTable,
		LevelsTable,
		UsersTable,
		UserLevelsTable,
	}
)

func init() {
	CredentialsTable.ForeignKeys[0].RefTable = UsersTable
	UserLevelsTable.ForeignKeys[0].RefTable = UsersTable
	UserLevelsTable.ForeignKeys[1].RefTable = LevelsTable
}
