// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// ClassroomsColumns holds the columns for the "classrooms" table.
	ClassroomsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "classroom_teacher", Type: field.TypeInt, Nullable: true},
		{Name: "classroom_level", Type: field.TypeInt, Nullable: true},
	}
	// ClassroomsTable holds the schema information for the "classrooms" table.
	ClassroomsTable = &schema.Table{
		Name:       "classrooms",
		Columns:    ClassroomsColumns,
		PrimaryKey: []*schema.Column{ClassroomsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "classrooms_users_teacher",
				Columns: []*schema.Column{ClassroomsColumns[3]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "classrooms_levels_level",
				Columns: []*schema.Column{ClassroomsColumns[4]},

				RefColumns: []*schema.Column{LevelsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
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
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "definition", Type: field.TypeJSON},
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
		{Name: "classroom_students", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "users_classrooms_students",
				Columns: []*schema.Column{UsersColumns[7]},

				RefColumns: []*schema.Column{ClassroomsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UserLevelsColumns holds the columns for the "user_levels" table.
	UserLevelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "code", Type: field.TypeString, Size: 2147483647},
		{Name: "workspace", Type: field.TypeString, Size: 2147483647},
		{Name: "user_level_developer", Type: field.TypeInt, Nullable: true},
		{Name: "user_level_level", Type: field.TypeInt, Nullable: true},
	}
	// UserLevelsTable holds the schema information for the "user_levels" table.
	UserLevelsTable = &schema.Table{
		Name:       "user_levels",
		Columns:    UserLevelsColumns,
		PrimaryKey: []*schema.Column{UserLevelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "user_levels_users_developer",
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
				Name:    "userlevel_user_level_developer_user_level_level",
				Unique:  true,
				Columns: []*schema.Column{UserLevelsColumns[5], UserLevelsColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ClassroomsTable,
		CredentialsTable,
		LevelsTable,
		UsersTable,
		UserLevelsTable,
	}
)

func init() {
	ClassroomsTable.ForeignKeys[0].RefTable = UsersTable
	ClassroomsTable.ForeignKeys[1].RefTable = LevelsTable
	CredentialsTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = ClassroomsTable
	UserLevelsTable.ForeignKeys[0].RefTable = UsersTable
	UserLevelsTable.ForeignKeys[1].RefTable = LevelsTable
}
