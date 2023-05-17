// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/comment"
	"kala/internal/ent/pros"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Pros is the model entity for the Pros schema.
type Pros struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pro holds the value of the "pro" field.
	Pro string `json:"pro,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProsQuery when eager-loading is set.
	Edges        ProsEdges `json:"edges"`
	comment      *int
	selectValues sql.SelectValues
}

// ProsEdges holds the relations/edges for other nodes in the graph.
type ProsEdges struct {
	// Comment holds the value of the comment edge.
	Comment *Comment `json:"comment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProsEdges) CommentOrErr() (*Comment, error) {
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
func (*Pros) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pros.FieldID:
			values[i] = new(sql.NullInt64)
		case pros.FieldPro:
			values[i] = new(sql.NullString)
		case pros.ForeignKeys[0]: // comment
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pros fields.
func (pr *Pros) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pros.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case pros.FieldPro:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pro", values[i])
			} else if value.Valid {
				pr.Pro = value.String
			}
		case pros.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field comment", value)
			} else if value.Valid {
				pr.comment = new(int)
				*pr.comment = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pros.
// This includes values selected through modifiers, order, etc.
func (pr *Pros) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryComment queries the "comment" edge of the Pros entity.
func (pr *Pros) QueryComment() *CommentQuery {
	return NewProsClient(pr.config).QueryComment(pr)
}

// Update returns a builder for updating this Pros.
// Note that you need to call Pros.Unwrap() before calling this method if this Pros
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Pros) Update() *ProsUpdateOne {
	return NewProsClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Pros entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Pros) Unwrap() *Pros {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pros is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Pros) String() string {
	var builder strings.Builder
	builder.WriteString("Pros(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("pro=")
	builder.WriteString(pr.Pro)
	builder.WriteByte(')')
	return builder.String()
}

// ProsSlice is a parsable slice of Pros.
type ProsSlice []*Pros
