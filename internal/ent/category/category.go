// Code generated by ent, DO NOT EDIT.

package category

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeSubCategory holds the string denoting the sub_category edge name in mutations.
	EdgeSubCategory = "sub_category"
	// EdgeImage holds the string denoting the image edge name in mutations.
	EdgeImage = "image"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// EdgeBrand holds the string denoting the brand edge name in mutations.
	EdgeBrand = "brand"
	// EdgeSeller holds the string denoting the seller edge name in mutations.
	EdgeSeller = "seller"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// SubCategoryTable is the table that holds the sub_category relation/edge.
	SubCategoryTable = "sub_categories"
	// SubCategoryInverseTable is the table name for the SubCategory entity.
	// It exists in this package in order to avoid circular dependency with the "subcategory" package.
	SubCategoryInverseTable = "sub_categories"
	// SubCategoryColumn is the table column denoting the sub_category relation/edge.
	SubCategoryColumn = "category"
	// ImageTable is the table that holds the image relation/edge. The primary key declared below.
	ImageTable = "category_images"
	// ImageInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	ImageInverseTable = "images"
	// ProductTable is the table that holds the product relation/edge. The primary key declared below.
	ProductTable = "product_category"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// BrandTable is the table that holds the brand relation/edge. The primary key declared below.
	BrandTable = "brand_category"
	// BrandInverseTable is the table name for the Brand entity.
	// It exists in this package in order to avoid circular dependency with the "brand" package.
	BrandInverseTable = "brands"
	// SellerTable is the table that holds the seller relation/edge. The primary key declared below.
	SellerTable = "seller_category"
	// SellerInverseTable is the table name for the Seller entity.
	// It exists in this package in order to avoid circular dependency with the "seller" package.
	SellerInverseTable = "sellers"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldDescription,
}

var (
	// ImagePrimaryKey and ImageColumn2 are the table columns denoting the
	// primary key for the image relation (M2M).
	ImagePrimaryKey = []string{"category", "image"}
	// ProductPrimaryKey and ProductColumn2 are the table columns denoting the
	// primary key for the product relation (M2M).
	ProductPrimaryKey = []string{"product_id", "category_id"}
	// BrandPrimaryKey and BrandColumn2 are the table columns denoting the
	// primary key for the brand relation (M2M).
	BrandPrimaryKey = []string{"brand_id", "category_id"}
	// SellerPrimaryKey and SellerColumn2 are the table columns denoting the
	// primary key for the seller relation (M2M).
	SellerPrimaryKey = []string{"seller_id", "category_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

// OrderOption defines the ordering options for the Category queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// BySubCategoryCount orders the results by sub_category count.
func BySubCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSubCategoryStep(), opts...)
	}
}

// BySubCategory orders the results by sub_category terms.
func BySubCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBrandCount orders the results by brand count.
func ByBrandCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBrandStep(), opts...)
	}
}

// ByBrand orders the results by brand terms.
func ByBrand(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBrandStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newSubCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubCategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SubCategoryTable, SubCategoryColumn),
	)
}
func newImageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ImageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ImageTable, ImagePrimaryKey...),
	)
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
	)
}
func newBrandStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BrandInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BrandTable, BrandPrimaryKey...),
	)
}
func newSellerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SellerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SellerTable, SellerPrimaryKey...),
	)
}
