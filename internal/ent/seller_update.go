// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/address"
	"kala/internal/ent/category"
	"kala/internal/ent/predicate"
	"kala/internal/ent/product"
	"kala/internal/ent/seller"
	"kala/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerUpdate is the builder for updating Seller entities.
type SellerUpdate struct {
	config
	hooks    []Hook
	mutation *SellerMutation
}

// Where appends a list predicates to the SellerUpdate builder.
func (su *SellerUpdate) Where(ps ...predicate.Seller) *SellerUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SellerUpdate) SetUpdateTime(t time.Time) *SellerUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetName sets the "name" field.
func (su *SellerUpdate) SetName(s string) *SellerUpdate {
	su.mutation.SetName(s)
	return su
}

// SetDescription sets the "description" field.
func (su *SellerUpdate) SetDescription(s string) *SellerUpdate {
	su.mutation.SetDescription(s)
	return su
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (su *SellerUpdate) SetNillableDescription(s *string) *SellerUpdate {
	if s != nil {
		su.SetDescription(*s)
	}
	return su
}

// ClearDescription clears the value of the "description" field.
func (su *SellerUpdate) ClearDescription() *SellerUpdate {
	su.mutation.ClearDescription()
	return su
}

// SetRating sets the "rating" field.
func (su *SellerUpdate) SetRating(f float64) *SellerUpdate {
	su.mutation.ResetRating()
	su.mutation.SetRating(f)
	return su
}

// AddRating adds f to the "rating" field.
func (su *SellerUpdate) AddRating(f float64) *SellerUpdate {
	su.mutation.AddRating(f)
	return su
}

// SetRatingCount sets the "rating_count" field.
func (su *SellerUpdate) SetRatingCount(i int32) *SellerUpdate {
	su.mutation.ResetRatingCount()
	su.mutation.SetRatingCount(i)
	return su
}

// AddRatingCount adds i to the "rating_count" field.
func (su *SellerUpdate) AddRatingCount(i int32) *SellerUpdate {
	su.mutation.AddRatingCount(i)
	return su
}

// SetPhone sets the "phone" field.
func (su *SellerUpdate) SetPhone(s string) *SellerUpdate {
	su.mutation.SetPhone(s)
	return su
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (su *SellerUpdate) AddProductIDs(ids ...int) *SellerUpdate {
	su.mutation.AddProductIDs(ids...)
	return su
}

// AddProduct adds the "product" edges to the Product entity.
func (su *SellerUpdate) AddProduct(p ...*Product) *SellerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddProductIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (su *SellerUpdate) AddCategoryIDs(ids ...int) *SellerUpdate {
	su.mutation.AddCategoryIDs(ids...)
	return su
}

// AddCategory adds the "category" edges to the Category entity.
func (su *SellerUpdate) AddCategory(c ...*Category) *SellerUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCategoryIDs(ids...)
}

// AddAddresIDs adds the "address" edge to the Address entity by IDs.
func (su *SellerUpdate) AddAddresIDs(ids ...int) *SellerUpdate {
	su.mutation.AddAddresIDs(ids...)
	return su
}

// AddAddress adds the "address" edges to the Address entity.
func (su *SellerUpdate) AddAddress(a ...*Address) *SellerUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddAddresIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (su *SellerUpdate) SetUserID(id int) *SellerUpdate {
	su.mutation.SetUserID(id)
	return su
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (su *SellerUpdate) SetNillableUserID(id *int) *SellerUpdate {
	if id != nil {
		su = su.SetUserID(*id)
	}
	return su
}

// SetUser sets the "user" edge to the User entity.
func (su *SellerUpdate) SetUser(u *User) *SellerUpdate {
	return su.SetUserID(u.ID)
}

// Mutation returns the SellerMutation object of the builder.
func (su *SellerUpdate) Mutation() *SellerMutation {
	return su.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (su *SellerUpdate) ClearProduct() *SellerUpdate {
	su.mutation.ClearProduct()
	return su
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (su *SellerUpdate) RemoveProductIDs(ids ...int) *SellerUpdate {
	su.mutation.RemoveProductIDs(ids...)
	return su
}

// RemoveProduct removes "product" edges to Product entities.
func (su *SellerUpdate) RemoveProduct(p ...*Product) *SellerUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveProductIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (su *SellerUpdate) ClearCategory() *SellerUpdate {
	su.mutation.ClearCategory()
	return su
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (su *SellerUpdate) RemoveCategoryIDs(ids ...int) *SellerUpdate {
	su.mutation.RemoveCategoryIDs(ids...)
	return su
}

// RemoveCategory removes "category" edges to Category entities.
func (su *SellerUpdate) RemoveCategory(c ...*Category) *SellerUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCategoryIDs(ids...)
}

// ClearAddress clears all "address" edges to the Address entity.
func (su *SellerUpdate) ClearAddress() *SellerUpdate {
	su.mutation.ClearAddress()
	return su
}

// RemoveAddresIDs removes the "address" edge to Address entities by IDs.
func (su *SellerUpdate) RemoveAddresIDs(ids ...int) *SellerUpdate {
	su.mutation.RemoveAddresIDs(ids...)
	return su
}

// RemoveAddress removes "address" edges to Address entities.
func (su *SellerUpdate) RemoveAddress(a ...*Address) *SellerUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveAddresIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (su *SellerUpdate) ClearUser() *SellerUpdate {
	su.mutation.ClearUser()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SellerUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SellerUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SellerUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SellerUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SellerUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := seller.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SellerUpdate) check() error {
	if v, ok := su.mutation.Name(); ok {
		if err := seller.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Seller.name": %w`, err)}
		}
	}
	if v, ok := su.mutation.Rating(); ok {
		if err := seller.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Seller.rating": %w`, err)}
		}
	}
	if v, ok := su.mutation.RatingCount(); ok {
		if err := seller.RatingCountValidator(v); err != nil {
			return &ValidationError{Name: "rating_count", err: fmt.Errorf(`ent: validator failed for field "Seller.rating_count": %w`, err)}
		}
	}
	if v, ok := su.mutation.Phone(); ok {
		if err := seller.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Seller.phone": %w`, err)}
		}
	}
	return nil
}

func (su *SellerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(seller.Table, seller.Columns, sqlgraph.NewFieldSpec(seller.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.SetField(seller.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(seller.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Description(); ok {
		_spec.SetField(seller.FieldDescription, field.TypeString, value)
	}
	if su.mutation.DescriptionCleared() {
		_spec.ClearField(seller.FieldDescription, field.TypeString)
	}
	if value, ok := su.mutation.Rating(); ok {
		_spec.SetField(seller.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedRating(); ok {
		_spec.AddField(seller.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.RatingCount(); ok {
		_spec.SetField(seller.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := su.mutation.AddedRatingCount(); ok {
		_spec.AddField(seller.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := su.mutation.Phone(); ok {
		_spec.SetField(seller.FieldPhone, field.TypeString, value)
	}
	if su.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedProductIDs(); len(nodes) > 0 && !su.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !su.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedAddressIDs(); len(nodes) > 0 && !su.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seller.UserTable,
			Columns: []string{seller.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seller.UserTable,
			Columns: []string{seller.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seller.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SellerUpdateOne is the builder for updating a single Seller entity.
type SellerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SellerMutation
}

// SetUpdateTime sets the "update_time" field.
func (suo *SellerUpdateOne) SetUpdateTime(t time.Time) *SellerUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetName sets the "name" field.
func (suo *SellerUpdateOne) SetName(s string) *SellerUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetDescription sets the "description" field.
func (suo *SellerUpdateOne) SetDescription(s string) *SellerUpdateOne {
	suo.mutation.SetDescription(s)
	return suo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (suo *SellerUpdateOne) SetNillableDescription(s *string) *SellerUpdateOne {
	if s != nil {
		suo.SetDescription(*s)
	}
	return suo
}

// ClearDescription clears the value of the "description" field.
func (suo *SellerUpdateOne) ClearDescription() *SellerUpdateOne {
	suo.mutation.ClearDescription()
	return suo
}

// SetRating sets the "rating" field.
func (suo *SellerUpdateOne) SetRating(f float64) *SellerUpdateOne {
	suo.mutation.ResetRating()
	suo.mutation.SetRating(f)
	return suo
}

// AddRating adds f to the "rating" field.
func (suo *SellerUpdateOne) AddRating(f float64) *SellerUpdateOne {
	suo.mutation.AddRating(f)
	return suo
}

// SetRatingCount sets the "rating_count" field.
func (suo *SellerUpdateOne) SetRatingCount(i int32) *SellerUpdateOne {
	suo.mutation.ResetRatingCount()
	suo.mutation.SetRatingCount(i)
	return suo
}

// AddRatingCount adds i to the "rating_count" field.
func (suo *SellerUpdateOne) AddRatingCount(i int32) *SellerUpdateOne {
	suo.mutation.AddRatingCount(i)
	return suo
}

// SetPhone sets the "phone" field.
func (suo *SellerUpdateOne) SetPhone(s string) *SellerUpdateOne {
	suo.mutation.SetPhone(s)
	return suo
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (suo *SellerUpdateOne) AddProductIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.AddProductIDs(ids...)
	return suo
}

// AddProduct adds the "product" edges to the Product entity.
func (suo *SellerUpdateOne) AddProduct(p ...*Product) *SellerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddProductIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (suo *SellerUpdateOne) AddCategoryIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.AddCategoryIDs(ids...)
	return suo
}

// AddCategory adds the "category" edges to the Category entity.
func (suo *SellerUpdateOne) AddCategory(c ...*Category) *SellerUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCategoryIDs(ids...)
}

// AddAddresIDs adds the "address" edge to the Address entity by IDs.
func (suo *SellerUpdateOne) AddAddresIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.AddAddresIDs(ids...)
	return suo
}

// AddAddress adds the "address" edges to the Address entity.
func (suo *SellerUpdateOne) AddAddress(a ...*Address) *SellerUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddAddresIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (suo *SellerUpdateOne) SetUserID(id int) *SellerUpdateOne {
	suo.mutation.SetUserID(id)
	return suo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (suo *SellerUpdateOne) SetNillableUserID(id *int) *SellerUpdateOne {
	if id != nil {
		suo = suo.SetUserID(*id)
	}
	return suo
}

// SetUser sets the "user" edge to the User entity.
func (suo *SellerUpdateOne) SetUser(u *User) *SellerUpdateOne {
	return suo.SetUserID(u.ID)
}

// Mutation returns the SellerMutation object of the builder.
func (suo *SellerUpdateOne) Mutation() *SellerMutation {
	return suo.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (suo *SellerUpdateOne) ClearProduct() *SellerUpdateOne {
	suo.mutation.ClearProduct()
	return suo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (suo *SellerUpdateOne) RemoveProductIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.RemoveProductIDs(ids...)
	return suo
}

// RemoveProduct removes "product" edges to Product entities.
func (suo *SellerUpdateOne) RemoveProduct(p ...*Product) *SellerUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveProductIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (suo *SellerUpdateOne) ClearCategory() *SellerUpdateOne {
	suo.mutation.ClearCategory()
	return suo
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (suo *SellerUpdateOne) RemoveCategoryIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.RemoveCategoryIDs(ids...)
	return suo
}

// RemoveCategory removes "category" edges to Category entities.
func (suo *SellerUpdateOne) RemoveCategory(c ...*Category) *SellerUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCategoryIDs(ids...)
}

// ClearAddress clears all "address" edges to the Address entity.
func (suo *SellerUpdateOne) ClearAddress() *SellerUpdateOne {
	suo.mutation.ClearAddress()
	return suo
}

// RemoveAddresIDs removes the "address" edge to Address entities by IDs.
func (suo *SellerUpdateOne) RemoveAddresIDs(ids ...int) *SellerUpdateOne {
	suo.mutation.RemoveAddresIDs(ids...)
	return suo
}

// RemoveAddress removes "address" edges to Address entities.
func (suo *SellerUpdateOne) RemoveAddress(a ...*Address) *SellerUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveAddresIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (suo *SellerUpdateOne) ClearUser() *SellerUpdateOne {
	suo.mutation.ClearUser()
	return suo
}

// Where appends a list predicates to the SellerUpdate builder.
func (suo *SellerUpdateOne) Where(ps ...predicate.Seller) *SellerUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SellerUpdateOne) Select(field string, fields ...string) *SellerUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Seller entity.
func (suo *SellerUpdateOne) Save(ctx context.Context) (*Seller, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SellerUpdateOne) SaveX(ctx context.Context) *Seller {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SellerUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SellerUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SellerUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := seller.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SellerUpdateOne) check() error {
	if v, ok := suo.mutation.Name(); ok {
		if err := seller.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Seller.name": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Rating(); ok {
		if err := seller.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Seller.rating": %w`, err)}
		}
	}
	if v, ok := suo.mutation.RatingCount(); ok {
		if err := seller.RatingCountValidator(v); err != nil {
			return &ValidationError{Name: "rating_count", err: fmt.Errorf(`ent: validator failed for field "Seller.rating_count": %w`, err)}
		}
	}
	if v, ok := suo.mutation.Phone(); ok {
		if err := seller.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Seller.phone": %w`, err)}
		}
	}
	return nil
}

func (suo *SellerUpdateOne) sqlSave(ctx context.Context) (_node *Seller, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(seller.Table, seller.Columns, sqlgraph.NewFieldSpec(seller.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Seller.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, seller.FieldID)
		for _, f := range fields {
			if !seller.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != seller.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.SetField(seller.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(seller.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Description(); ok {
		_spec.SetField(seller.FieldDescription, field.TypeString, value)
	}
	if suo.mutation.DescriptionCleared() {
		_spec.ClearField(seller.FieldDescription, field.TypeString)
	}
	if value, ok := suo.mutation.Rating(); ok {
		_spec.SetField(seller.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedRating(); ok {
		_spec.AddField(seller.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.RatingCount(); ok {
		_spec.SetField(seller.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.AddedRatingCount(); ok {
		_spec.AddField(seller.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := suo.mutation.Phone(); ok {
		_spec.SetField(seller.FieldPhone, field.TypeString, value)
	}
	if suo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedProductIDs(); len(nodes) > 0 && !suo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !suo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedAddressIDs(); len(nodes) > 0 && !suo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seller.UserTable,
			Columns: []string{seller.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seller.UserTable,
			Columns: []string{seller.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Seller{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{seller.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}