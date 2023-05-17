// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeImage holds the string denoting the image edge name in mutations.
	EdgeImage = "image"
	// EdgeSeller holds the string denoting the seller edge name in mutations.
	EdgeSeller = "seller"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// EdgeLogs holds the string denoting the logs edge name in mutations.
	EdgeLogs = "logs"
	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CommentTable is the table that holds the comment relation/edge. The primary key declared below.
	CommentTable = "user_comments"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// ImageTable is the table that holds the image relation/edge.
	ImageTable = "images"
	// ImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImageInverseTable = "images"
	// ImageColumn is the table column denoting the image relation/edge.
	ImageColumn = "user_image"
	// SellerTable is the table that holds the seller relation/edge.
	SellerTable = "sellers"
	// SellerInverseTable is the table name for the Seller entity.
	// It exists in this package in order to avoid circular dependency with the "seller" package.
	SellerInverseTable = "sellers"
	// SellerColumn is the table column denoting the seller relation/edge.
	SellerColumn = "user_id"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "orders"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "user_order"
	// LogsTable is the table that holds the logs relation/edge.
	LogsTable = "logs"
	// LogsInverseTable is the table name for the Logs entity.
	// It exists in this package in order to avoid circular dependency with the "logs" package.
	LogsInverseTable = "logs"
	// LogsColumn is the table column denoting the logs relation/edge.
	LogsColumn = "user"
	// AddressTable is the table that holds the address relation/edge.
	AddressTable = "addresses"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUsername,
	FieldPassword,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPhone,
	FieldRole,
	FieldStatus,
}

var (
	// CommentPrimaryKey and CommentColumn2 are the table columns denoting the
	// primary key for the comment relation (M2M).
	CommentPrimaryKey = []string{"user", "comment"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus bool
)

// Role defines the type for the "role" enum field.
type Role string

// RoleUser is the default value of the Role enum.
const DefaultRole = RoleUser

// Role values.
const (
	RoleAdmin  Role = "admin"
	RoleSeller Role = "seller"
	RoleUser   Role = "user"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleSeller, RoleUser:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByCommentCount orders the results by comment count.
func ByCommentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCommentStep(), opts...)
	}
}

// ByComment orders the results by comment terms.
func ByComment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByImageCount orders the results by image count.
func ByImageCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newImageStep(), opts...)
	}
}

// ByImage orders the results by image terms.
func ByImage(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newImageStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySellerCount orders the results by seller count.
func BySellerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSellerStep(), opts...)
	}
}

// BySeller orders the results by seller terms.
func BySeller(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSellerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderCount orders the results by order count.
func ByOrderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderStep(), opts...)
	}
}

// ByOrder orders the results by order terms.
func ByOrder(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLogsCount orders the results by logs count.
func ByLogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLogsStep(), opts...)
	}
}

// ByLogs orders the results by logs terms.
func ByLogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAddressCount orders the results by address count.
func ByAddressCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAddressStep(), opts...)
	}
}

// ByAddress orders the results by address terms.
func ByAddress(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAddressStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCommentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, CommentTable, CommentPrimaryKey...),
	)
}
func newImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ImageTable, ImageColumn),
	)
}
func newSellerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SellerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SellerTable, SellerColumn),
	)
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderTable, OrderColumn),
	)
}
func newLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LogsTable, LogsColumn),
	)
}
func newAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AddressTable, AddressColumn),
	)
}