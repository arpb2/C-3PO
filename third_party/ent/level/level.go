// Code generated by entc, DO NOT EDIT.

package level

import (
	"time"
)

const (
	// Label holds the string label denoting the level type in the database.
	Label = "level"
	// FieldID holds the string denoting the id field in the database.
	FieldID          = "id"          // FieldCreatedAt holds the string denoting the created_at vertex property in the database.
	FieldCreatedAt   = "created_at"  // FieldUpdatedAt holds the string denoting the updated_at vertex property in the database.
	FieldUpdatedAt   = "updated_at"  // FieldName holds the string denoting the name vertex property in the database.
	FieldName        = "name"        // FieldDescription holds the string denoting the description vertex property in the database.
	FieldDescription = "description" // FieldDefinition holds the string denoting the definition vertex property in the database.
	FieldDefinition  = "definition"

	// Table holds the table name of the level in the database.
	Table = "levels"
)

// Columns holds all SQL columns for level fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldDefinition,
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)
