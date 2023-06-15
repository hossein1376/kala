// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kala/internal/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password []byte `json:"-"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Comment holds the value of the comment edge.
	Comment []*Comment `json:"comment,omitempty"`
	// Image holds the value of the image edge.
	Image []*Image `json:"image,omitempty"`
	// Seller holds the value of the seller edge.
	Seller []*Seller `json:"seller,omitempty"`
	// Order holds the value of the order edge.
	Order []*Order `json:"order,omitempty"`
	// Logs holds the value of the logs edge.
	Logs []*Logs `json:"logs,omitempty"`
	// Address holds the value of the address edge.
	Address []*Address `json:"address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CommentOrErr() ([]*Comment, error) {
	if e.loadedTypes[0] {
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ImageOrErr() ([]*Image, error) {
	if e.loadedTypes[1] {
		return e.Image, nil
	}
	return nil, &NotLoadedError{edge: "image"}
}

// SellerOrErr returns the Seller value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SellerOrErr() ([]*Seller, error) {
	if e.loadedTypes[2] {
		return e.Seller, nil
	}
	return nil, &NotLoadedError{edge: "seller"}
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) OrderOrErr() ([]*Order, error) {
	if e.loadedTypes[3] {
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// LogsOrErr returns the Logs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LogsOrErr() ([]*Logs, error) {
	if e.loadedTypes[4] {
		return e.Logs, nil
	}
	return nil, &NotLoadedError{edge: "logs"}
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AddressOrErr() ([]*Address, error) {
	if e.loadedTypes[5] {
		return e.Address, nil
	}
	return nil, &NotLoadedError{edge: "address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldPassword:
			values[i] = new([]byte)
		case user.FieldStatus:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldEmail, user.FieldPhone, user.FieldRole:
			values[i] = new(sql.NullString)
		case user.FieldCreateTime, user.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				u.CreateTime = value.Time
			}
		case user.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				u.UpdateTime = value.Time
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				u.Password = *value
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				u.Role = user.Role(value.String)
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = value.Bool
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryComment queries the "comment" edge of the User entity.
func (u *User) QueryComment() *CommentQuery {
	return NewUserClient(u.config).QueryComment(u)
}

// QueryImage queries the "image" edge of the User entity.
func (u *User) QueryImage() *ImageQuery {
	return NewUserClient(u.config).QueryImage(u)
}

// QuerySeller queries the "seller" edge of the User entity.
func (u *User) QuerySeller() *SellerQuery {
	return NewUserClient(u.config).QuerySeller(u)
}

// QueryOrder queries the "order" edge of the User entity.
func (u *User) QueryOrder() *OrderQuery {
	return NewUserClient(u.config).QueryOrder(u)
}

// QueryLogs queries the "logs" edge of the User entity.
func (u *User) QueryLogs() *LogsQuery {
	return NewUserClient(u.config).QueryLogs(u)
}

// QueryAddress queries the "address" edge of the User entity.
func (u *User) QueryAddress() *AddressQuery {
	return NewUserClient(u.config).QueryAddress(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("create_time=")
	builder.WriteString(u.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(u.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(u.LastName)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
