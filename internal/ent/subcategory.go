// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/category"
	"kala/internal/ent/subcategory"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// SubCategory is the model entity for the SubCategory schema.
type SubCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubCategoryQuery when eager-loading is set.
	Edges        SubCategoryEdges `json:"edges"`
	category     *int
	selectValues sql.SelectValues
}

// SubCategoryEdges holds the relations/edges for other nodes in the graph.
type SubCategoryEdges struct {
	// Category holds the value of the category edge.
	Category *Category `json:"category,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubCategoryEdges) CategoryOrErr() (*Category, error) {
	if e.loadedTypes[0] {
		if e.Category == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: category.Label}
		}
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case subcategory.FieldName, subcategory.FieldDescription:
			values[i] = new(sql.NullString)
		case subcategory.FieldCreateTime, subcategory.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case subcategory.ForeignKeys[0]: // category
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubCategory fields.
func (sc *SubCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case subcategory.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				sc.CreateTime = value.Time
			}
		case subcategory.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				sc.UpdateTime = value.Time
			}
		case subcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sc.Name = value.String
			}
		case subcategory.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				sc.Description = value.String
			}
		case subcategory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field category", value)
			} else if value.Valid {
				sc.category = new(int)
				*sc.category = int(value.Int64)
			}
		default:
			sc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SubCategory.
// This includes values selected through modifiers, order, etc.
func (sc *SubCategory) Value(name string) (ent.Value, error) {
	return sc.selectValues.Get(name)
}

// QueryCategory queries the "category" edge of the SubCategory entity.
func (sc *SubCategory) QueryCategory() *CategoryQuery {
	return NewSubCategoryClient(sc.config).QueryCategory(sc)
}

// Update returns a builder for updating this SubCategory.
// Note that you need to call SubCategory.Unwrap() before calling this method if this SubCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SubCategory) Update() *SubCategoryUpdateOne {
	return NewSubCategoryClient(sc.config).UpdateOne(sc)
}

// Unwrap unwraps the SubCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SubCategory) Unwrap() *SubCategory {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SubCategory is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SubCategory) String() string {
	var builder strings.Builder
	builder.WriteString("SubCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("create_time=")
	builder.WriteString(sc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(sc.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sc.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(sc.Description)
	builder.WriteByte(')')
	return builder.String()
}

// SubCategories is a parsable slice of SubCategory.
type SubCategories []*SubCategory
