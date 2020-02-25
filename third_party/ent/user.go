// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/arpb2/C-3PO/third_party/ent/credential"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Surname holds the value of the "surname" field.
	Surname string `json:"surname,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Levels holds the value of the levels edge.
	Levels []*UserLevel
	// Credentials holds the value of the credentials edge.
	Credentials *Credential
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// LevelsOrErr returns the Levels value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LevelsOrErr() ([]*UserLevel, error) {
	if e.loadedTypes[0] {
		return e.Levels, nil
	}
	return nil, &NotLoadedError{edge: "levels"}
}

// CredentialsOrErr returns the Credentials value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CredentialsOrErr() (*Credential, error) {
	if e.loadedTypes[1] {
		if e.Credentials == nil {
			// The edge credentials was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: credential.Label}
		}
		return e.Credentials, nil
	}
	return nil, &NotLoadedError{edge: "credentials"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // email
		&sql.NullString{}, // name
		&sql.NullString{}, // surname
		&sql.NullTime{},   // created_at
		&sql.NullTime{},   // updated_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = uint(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[0])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field surname", values[2])
	} else if value.Valid {
		u.Surname = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[3])
	} else if value.Valid {
		u.CreatedAt = value.Time
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[4])
	} else if value.Valid {
		u.UpdatedAt = value.Time
	}
	return nil
}

// QueryLevels queries the levels edge of the User.
func (u *User) QueryLevels() *UserLevelQuery {
	return (&UserClient{u.config}).QueryLevels(u)
}

// QueryCredentials queries the credentials edge of the User.
func (u *User) QueryCredentials() *CredentialQuery {
	return (&UserClient{u.config}).QueryCredentials(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", surname=")
	builder.WriteString(u.Surname)
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}