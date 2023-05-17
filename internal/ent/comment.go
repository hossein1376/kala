// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/comment"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Status holds the value of the "status" field.
	Status comment.Status `json:"status,omitempty"`
	// Likes holds the value of the "likes" field.
	Likes int32 `json:"likes,omitempty"`
	// Dislikes holds the value of the "dislikes" field.
	Dislikes int32 `json:"dislikes,omitempty"`
	// Rating holds the value of the "rating" field.
	Rating int8 `json:"rating,omitempty"`
	// RatingCount holds the value of the "rating_count" field.
	RatingCount int32 `json:"rating_count,omitempty"`
	// VerifiedBuyer holds the value of the "verified_buyer" field.
	VerifiedBuyer bool `json:"verified_buyer,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges        CommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Image holds the value of the image edge.
	Image []*Image `json:"image,omitempty"`
	// Cons holds the value of the cons edge.
	Cons []*Cons `json:"cons,omitempty"`
	// Pros holds the value of the pros edge.
	Pros []*Pros `json:"pros,omitempty"`
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) ImageOrErr() ([]*Image, error) {
	if e.loadedTypes[0] {
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// ConsOrErr returns the Cons value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) ConsOrErr() ([]*Cons, error) {
	if e.loadedTypes[1] {
		return e.Cons, nil
	}
	return nil, &NotLoadedError{edge: "cons"}
}

// ProsOrErr returns the Pros value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) ProsOrErr() ([]*Pros, error) {
	if e.loadedTypes[2] {
		return e.Pros, nil
	}
	return nil, &NotLoadedError{edge: "pros"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[3] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e CommentEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[4] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldVerifiedBuyer:
			values[i] = new(sql.NullBool)
		case comment.FieldID, comment.FieldLikes, comment.FieldDislikes, comment.FieldRating, comment.FieldRatingCount:
			values[i] = new(sql.NullInt64)
		case comment.FieldContent, comment.FieldStatus:
			values[i] = new(sql.NullString)
		case comment.FieldCreateTime, comment.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case comment.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case comment.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = comment.Status(value.String)
			}
		case comment.FieldLikes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field likes", values[i])
			} else if value.Valid {
				c.Likes = int32(value.Int64)
			}
		case comment.FieldDislikes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field dislikes", values[i])
			} else if value.Valid {
				c.Dislikes = int32(value.Int64)
			}
		case comment.FieldRating:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rating", values[i])
			} else if value.Valid {
				c.Rating = int8(value.Int64)
			}
		case comment.FieldRatingCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rating_count", values[i])
			} else if value.Valid {
				c.RatingCount = int32(value.Int64)
			}
		case comment.FieldVerifiedBuyer:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified_buyer", values[i])
			} else if value.Valid {
				c.VerifiedBuyer = value.Bool
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (c *Comment) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryImage queries the "image" edge of the Comment entity.
func (c *Comment) QueryImage() *ImageQuery {
	return NewCommentClient(c.config).QueryImage(c)
}

// QueryCons queries the "cons" edge of the Comment entity.
func (c *Comment) QueryCons() *ConsQuery {
	return NewCommentClient(c.config).QueryCons(c)
}

// QueryPros queries the "pros" edge of the Comment entity.
func (c *Comment) QueryPros() *ProsQuery {
	return NewCommentClient(c.config).QueryPros(c)
}

// QueryUser queries the "user" edge of the Comment entity.
func (c *Comment) QueryUser() *UserQuery {
	return NewCommentClient(c.config).QueryUser(c)
}

// QueryProduct queries the "product" edge of the Comment entity.
func (c *Comment) QueryProduct() *ProductQuery {
	return NewCommentClient(c.config).QueryProduct(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(c.Content)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("likes=")
	builder.WriteString(fmt.Sprintf("%v", c.Likes))
	builder.WriteString(", ")
	builder.WriteString("dislikes=")
	builder.WriteString(fmt.Sprintf("%v", c.Dislikes))
	builder.WriteString(", ")
	builder.WriteString("rating=")
	builder.WriteString(fmt.Sprintf("%v", c.Rating))
	builder.WriteString(", ")
	builder.WriteString("rating_count=")
	builder.WriteString(fmt.Sprintf("%v", c.RatingCount))
	builder.WriteString(", ")
	builder.WriteString("verified_buyer=")
	builder.WriteString(fmt.Sprintf("%v", c.VerifiedBuyer))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment