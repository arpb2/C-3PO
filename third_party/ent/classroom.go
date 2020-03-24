// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/arpb2/C-3PO/third_party/ent/classroom"
	"github.com/arpb2/C-3PO/third_party/ent/level"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Classroom is the model entity for the Classroom schema.
type Classroom struct {
	config `json:"-"`
	// ID of the ent.
	ID uint `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClassroomQuery when eager-loading is set.
	Edges             ClassroomEdges `json:"edges"`
	classroom_teacher *uint
	classroom_level   *uint
}

// ClassroomEdges holds the relations/edges for other nodes in the graph.
type ClassroomEdges struct {
	// Teacher holds the value of the teacher edge.
	Teacher *User
	// Students holds the value of the students edge.
	Students []*User
	// Level holds the value of the level edge.
	Level *Level
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// TeacherOrErr returns the Teacher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassroomEdges) TeacherOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Teacher == nil {
			// The edge teacher was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Teacher, nil
	}
	return nil, &NotLoadedError{edge: "teacher"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e ClassroomEdges) StudentsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// LevelOrErr returns the Level value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ClassroomEdges) LevelOrErr() (*Level, error) {
	if e.loadedTypes[2] {
		if e.Level == nil {
			// The edge level was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: level.Label}
		}
		return e.Level, nil
	}
	return nil, &NotLoadedError{edge: "level"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Classroom) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // created_at
		&sql.NullTime{},  // updated_at
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Classroom) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // classroom_teacher
		&sql.NullInt64{}, // classroom_level
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Classroom fields.
func (c *Classroom) assignValues(values ...interface{}) error {
	if m, n := len(values), len(classroom.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = uint(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field created_at", values[0])
	} else if value.Valid {
		c.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updated_at", values[1])
	} else if value.Valid {
		c.UpdatedAt = value.Time
	}
	values = values[2:]
	if len(values) == len(classroom.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field classroom_teacher", value)
		} else if value.Valid {
			c.classroom_teacher = new(uint)
			*c.classroom_teacher = uint(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field classroom_level", value)
		} else if value.Valid {
			c.classroom_level = new(uint)
			*c.classroom_level = uint(value.Int64)
		}
	}
	return nil
}

// QueryTeacher queries the teacher edge of the Classroom.
func (c *Classroom) QueryTeacher() *UserQuery {
	return (&ClassroomClient{config: c.config}).QueryTeacher(c)
}

// QueryStudents queries the students edge of the Classroom.
func (c *Classroom) QueryStudents() *UserQuery {
	return (&ClassroomClient{config: c.config}).QueryStudents(c)
}

// QueryLevel queries the level edge of the Classroom.
func (c *Classroom) QueryLevel() *LevelQuery {
	return (&ClassroomClient{config: c.config}).QueryLevel(c)
}

// Update returns a builder for updating this Classroom.
// Note that, you need to call Classroom.Unwrap() before calling this method, if this Classroom
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Classroom) Update() *ClassroomUpdateOne {
	return (&ClassroomClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Classroom) Unwrap() *Classroom {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Classroom is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Classroom) String() string {
	var builder strings.Builder
	builder.WriteString("Classroom(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Classrooms is a parsable slice of Classroom.
type Classrooms []*Classroom

func (c Classrooms) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
