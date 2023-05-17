// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/comment"
	"kala/internal/ent/cons"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Cons is the model entity for the Cons schema.
type Cons struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Con holds the value of the "con" field.
	Con string `json:"con,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ConsQuery when eager-loading is set.
	Edges        ConsEdges `json:"edges"`
	comment      *int
	selectValues sql.SelectValues
}

// ConsEdges holds the relations/edges for other nodes in the graph.
type ConsEdges struct {
	// Comment holds the value of the comment edge.
	Comment *Comment `json:"comment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ConsEdges) CommentOrErr() (*Comment, error) {
	if e.loadedTypes[0] {
		if e.Comment == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: comment.Label}
		}
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cons) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cons.FieldID:
			values[i] = new(sql.NullInt64)
		case cons.FieldCon:
			values[i] = new(sql.NullString)
		case cons.ForeignKeys[0]: // comment
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cons fields.
func (c *Cons) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cons.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cons.FieldCon:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field con", values[i])
			} else if value.Valid {
				c.Con = value.String
			}
		case cons.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field comment", value)
			} else if value.Valid {
				c.comment = new(int)
				*c.comment = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cons.
// This includes values selected through modifiers, order, etc.
func (c *Cons) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryComment queries the "comment" edge of the Cons entity.
func (c *Cons) QueryComment() *CommentQuery {
	return NewConsClient(c.config).QueryComment(c)
}

// Update returns a builder for updating this Cons.
// Note that you need to call Cons.Unwrap() before calling this method if this Cons
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cons) Update() *ConsUpdateOne {
	return NewConsClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cons entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cons) Unwrap() *Cons {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cons is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cons) String() string {
	var builder strings.Builder
	builder.WriteString("Cons(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("con=")
	builder.WriteString(c.Con)
	builder.WriteByte(')')
	return builder.String()
}

// ConsSlice is a parsable slice of Cons.
type ConsSlice []*Cons
