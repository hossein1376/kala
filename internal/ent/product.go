// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/brand"
	"kala/internal/ent/product"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
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
	// Review holds the value of the "review" field.
	Review string `json:"review,omitempty"`
	// Rating holds the value of the "rating" field.
	Rating float64 `json:"rating,omitempty"`
	// RatingCount holds the value of the "rating_count" field.
	RatingCount int32 `json:"rating_count,omitempty"`
	// Price holds the value of the "price" field.
	Price int32 `json:"price,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int32 `json:"quantity,omitempty"`
	// Available holds the value of the "available" field.
	Available bool `json:"available,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges          ProductEdges `json:"edges"`
	brand_product  *int
	seller_product *int
	selectValues   sql.SelectValues
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Values holds the value of the values edge.
	Values []*AttributeValue `json:"values,omitempty"`
	// Comment holds the value of the comment edge.
	Comment []*Comment `json:"comment,omitempty"`
	// Image holds the value of the image edge.
	Image []*Image `json:"image,omitempty"`
	// Order holds the value of the order edge.
	Order []*Order `json:"order,omitempty"`
	// Brand holds the value of the brand edge.
	Brand *Brand `json:"brand,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ValuesOrErr returns the Values value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ValuesOrErr() ([]*AttributeValue, error) {
	if e.loadedTypes[0] {
		return e.Values, nil
	}
	return nil, &NotLoadedError{edge: "values"}
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) CommentOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ImageOrErr() ([]*Image, error) {
	if e.loadedTypes[2] {
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) OrderOrErr() ([]*Order, error) {
	if e.loadedTypes[3] {
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// BrandOrErr returns the Brand value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) BrandOrErr() (*Brand, error) {
	if e.loadedTypes[4] {
		if e.Brand == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: brand.Label}
		}
		return e.Brand, nil
	}
	return nil, &NotLoadedError{edge: "brand"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldAvailable:
			values[i] = new(sql.NullBool)
		case product.FieldRating:
			values[i] = new(sql.NullFloat64)
		case product.FieldID, product.FieldRatingCount, product.FieldPrice, product.FieldQuantity:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldDescription, product.FieldReview:
			values[i] = new(sql.NullString)
		case product.FieldCreateTime, product.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case product.ForeignKeys[0]: // brand_product
			values[i] = new(sql.NullInt64)
		case product.ForeignKeys[1]: // seller_product
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case product.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		case product.FieldReview:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review", values[i])
			} else if value.Valid {
				pr.Review = value.String
			}
		case product.FieldRating:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field rating", values[i])
			} else if value.Valid {
				pr.Rating = value.Float64
			}
		case product.FieldRatingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rating_count", values[i])
			} else if value.Valid {
				pr.RatingCount = int32(value.Int64)
			}
		case product.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				pr.Price = int32(value.Int64)
			}
		case product.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				pr.Quantity = int32(value.Int64)
			}
		case product.FieldAvailable:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field available", values[i])
			} else if value.Valid {
				pr.Available = value.Bool
			}
		case product.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field brand_product", value)
			} else if value.Valid {
				pr.brand_product = new(int)
				*pr.brand_product = int(value.Int64)
			}
		case product.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field seller_product", value)
			} else if value.Valid {
				pr.seller_product = new(int)
				*pr.seller_product = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryValues queries the "values" edge of the Product entity.
func (pr *Product) QueryValues() *AttributeValueQuery {
	return NewProductClient(pr.config).QueryValues(pr)
}

// QueryComment queries the "comment" edge of the Product entity.
func (pr *Product) QueryComment() *CommentQuery {
	return NewProductClient(pr.config).QueryComment(pr)
}

// QueryImage queries the "image" edge of the Product entity.
func (pr *Product) QueryImage() *ImageQuery {
	return NewProductClient(pr.config).QueryImage(pr)
}

// QueryOrder queries the "order" edge of the Product entity.
func (pr *Product) QueryOrder() *OrderQuery {
	return NewProductClient(pr.config).QueryOrder(pr)
}

// QueryBrand queries the "brand" edge of the Product entity.
func (pr *Product) QueryBrand() *BrandQuery {
	return NewProductClient(pr.config).QueryBrand(pr)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteString(", ")
	builder.WriteString("review=")
	builder.WriteString(pr.Review)
	builder.WriteString(", ")
	builder.WriteString("rating=")
	builder.WriteString(fmt.Sprintf("%v", pr.Rating))
	builder.WriteString(", ")
	builder.WriteString("rating_count=")
	builder.WriteString(fmt.Sprintf("%v", pr.RatingCount))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", pr.Price))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", pr.Quantity))
	builder.WriteString(", ")
	builder.WriteString("available=")
	builder.WriteString(fmt.Sprintf("%v", pr.Available))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product
