// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/brand"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Brand is the model entity for the Brand schema.
type Brand struct {
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
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Rating holds the value of the "rating" field.
	Rating float64 `json:"rating,omitempty"`
	// RatingCount holds the value of the "rating_count" field.
	RatingCount int32 `json:"rating_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BrandQuery when eager-loading is set.
	Edges        BrandEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BrandEdges holds the relations/edges for other nodes in the graph.
type BrandEdges struct {
	// Image holds the value of the image edge.
	Image []*Image `json:"image,omitempty"`
	// Category holds the value of the category edge.
	Category []*Category `json:"category,omitempty"`
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) ImageOrErr() ([]*Image, error) {
	if e.loadedTypes[0] {
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) CategoryOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[2] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Brand) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case brand.FieldStatus:
			values[i] = new(sql.NullBool)
		case brand.FieldRating:
			values[i] = new(sql.NullFloat64)
		case brand.FieldID, brand.FieldRatingCount:
			values[i] = new(sql.NullInt64)
		case brand.FieldName, brand.FieldDescription:
			values[i] = new(sql.NullString)
		case brand.FieldCreateTime, brand.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Brand fields.
func (b *Brand) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case brand.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case brand.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				b.CreateTime = value.Time
			}
		case brand.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				b.UpdateTime = value.Time
			}
		case brand.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case brand.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case brand.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.Bool
			}
		case brand.FieldRating:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rating", values[i])
			} else if value.Valid {
				b.Rating = value.Float64
			}
		case brand.FieldRatingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rating_count", values[i])
			} else if value.Valid {
				b.RatingCount = int32(value.Int64)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Brand.
// This includes values selected through modifiers, order, etc.
func (b *Brand) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryImage queries the "image" edge of the Brand entity.
func (b *Brand) QueryImage() *ImageQuery {
	return NewBrandClient(b.config).QueryImage(b)
}

// QueryCategory queries the "category" edge of the Brand entity.
func (b *Brand) QueryCategory() *CategoryQuery {
	return NewBrandClient(b.config).QueryCategory(b)
}

// QueryProduct queries the "product" edge of the Brand entity.
func (b *Brand) QueryProduct() *ProductQuery {
	return NewBrandClient(b.config).QueryProduct(b)
}

// Update returns a builder for updating this Brand.
// Note that you need to call Brand.Unwrap() before calling this method if this Brand
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Brand) Update() *BrandUpdateOne {
	return NewBrandClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Brand entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Brand) Unwrap() *Brand {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Brand is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Brand) String() string {
	var builder strings.Builder
	builder.WriteString("Brand(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("create_time=")
	builder.WriteString(b.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(b.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("rating=")
	builder.WriteString(fmt.Sprintf("%v", b.Rating))
	builder.WriteString(", ")
	builder.WriteString("rating_count=")
	builder.WriteString(fmt.Sprintf("%v", b.RatingCount))
	builder.WriteByte(')')
	return builder.String()
}

// Brands is a parsable slice of Brand.
type Brands []*Brand