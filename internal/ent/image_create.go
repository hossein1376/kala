// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/brand"
	"kala/internal/ent/comment"
	"kala/internal/ent/image"
	"kala/internal/ent/product"
	"kala/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImageCreate is the builder for creating a Image entity.
type ImageCreate struct {
	config
	mutation *ImageMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *ImageCreate) SetName(s string) *ImageCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetPath sets the "path" field.
func (ic *ImageCreate) SetPath(s string) *ImageCreate {
	ic.mutation.SetPath(s)
	return ic
}

// SetCaption sets the "caption" field.
func (ic *ImageCreate) SetCaption(s string) *ImageCreate {
	ic.mutation.SetCaption(s)
	return ic
}

// SetNillableCaption sets the "caption" field if the given value is not nil.
func (ic *ImageCreate) SetNillableCaption(s *string) *ImageCreate {
	if s != nil {
		ic.SetCaption(*s)
	}
	return ic
}

// SetWidth sets the "width" field.
func (ic *ImageCreate) SetWidth(i int32) *ImageCreate {
	ic.mutation.SetWidth(i)
	return ic
}

// SetHigh sets the "high" field.
func (ic *ImageCreate) SetHigh(i int32) *ImageCreate {
	ic.mutation.SetHigh(i)
	return ic
}

// SetSizeKB sets the "size_kb" field.
func (ic *ImageCreate) SetSizeKB(f float64) *ImageCreate {
	ic.mutation.SetSizeKB(f)
	return ic
}

// SetUploadedAt sets the "uploaded_at" field.
func (ic *ImageCreate) SetUploadedAt(t time.Time) *ImageCreate {
	ic.mutation.SetUploadedAt(t)
	return ic
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ic *ImageCreate) SetUserID(id int) *ImageCreate {
	ic.mutation.SetUserID(id)
	return ic
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableUserID(id *int) *ImageCreate {
	if id != nil {
		ic = ic.SetUserID(*id)
	}
	return ic
}

// SetUser sets the "user" edge to the User entity.
func (ic *ImageCreate) SetUser(u *User) *ImageCreate {
	return ic.SetUserID(u.ID)
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (ic *ImageCreate) SetCommentID(id int) *ImageCreate {
	ic.mutation.SetCommentID(id)
	return ic
}

// SetNillableCommentID sets the "comment" edge to the Comment entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableCommentID(id *int) *ImageCreate {
	if id != nil {
		ic = ic.SetCommentID(*id)
	}
	return ic
}

// SetComment sets the "comment" edge to the Comment entity.
func (ic *ImageCreate) SetComment(c *Comment) *ImageCreate {
	return ic.SetCommentID(c.ID)
}

// SetBrandID sets the "brand" edge to the Brand entity by ID.
func (ic *ImageCreate) SetBrandID(id int) *ImageCreate {
	ic.mutation.SetBrandID(id)
	return ic
}

// SetNillableBrandID sets the "brand" edge to the Brand entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableBrandID(id *int) *ImageCreate {
	if id != nil {
		ic = ic.SetBrandID(*id)
	}
	return ic
}

// SetBrand sets the "brand" edge to the Brand entity.
func (ic *ImageCreate) SetBrand(b *Brand) *ImageCreate {
	return ic.SetBrandID(b.ID)
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (ic *ImageCreate) SetProductID(id int) *ImageCreate {
	ic.mutation.SetProductID(id)
	return ic
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (ic *ImageCreate) SetNillableProductID(id *int) *ImageCreate {
	if id != nil {
		ic = ic.SetProductID(*id)
	}
	return ic
}

// SetProduct sets the "product" edge to the Product entity.
func (ic *ImageCreate) SetProduct(p *Product) *ImageCreate {
	return ic.SetProductID(p.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (ic *ImageCreate) Mutation() *ImageMutation {
	return ic.mutation
}

// Save creates the Image in the database.
func (ic *ImageCreate) Save(ctx context.Context) (*Image, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ImageCreate) SaveX(ctx context.Context) *Image {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ImageCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ImageCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ImageCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Image.name"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := image.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Image.name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Image.path"`)}
	}
	if v, ok := ic.mutation.Path(); ok {
		if err := image.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Image.path": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Width(); !ok {
		return &ValidationError{Name: "width", err: errors.New(`ent: missing required field "Image.width"`)}
	}
	if v, ok := ic.mutation.Width(); ok {
		if err := image.WidthValidator(v); err != nil {
			return &ValidationError{Name: "width", err: fmt.Errorf(`ent: validator failed for field "Image.width": %w`, err)}
		}
	}
	if _, ok := ic.mutation.High(); !ok {
		return &ValidationError{Name: "high", err: errors.New(`ent: missing required field "Image.high"`)}
	}
	if v, ok := ic.mutation.High(); ok {
		if err := image.HighValidator(v); err != nil {
			return &ValidationError{Name: "high", err: fmt.Errorf(`ent: validator failed for field "Image.high": %w`, err)}
		}
	}
	if _, ok := ic.mutation.SizeKB(); !ok {
		return &ValidationError{Name: "size_kb", err: errors.New(`ent: missing required field "Image.size_kb"`)}
	}
	if v, ok := ic.mutation.SizeKB(); ok {
		if err := image.SizeKBValidator(v); err != nil {
			return &ValidationError{Name: "size_kb", err: fmt.Errorf(`ent: validator failed for field "Image.size_kb": %w`, err)}
		}
	}
	if _, ok := ic.mutation.UploadedAt(); !ok {
		return &ValidationError{Name: "uploaded_at", err: errors.New(`ent: missing required field "Image.uploaded_at"`)}
	}
	return nil
}

func (ic *ImageCreate) sqlSave(ctx context.Context) (*Image, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ImageCreate) createSpec() (*Image, *sqlgraph.CreateSpec) {
	var (
		_node = &Image{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(image.Table, sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(image.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Path(); ok {
		_spec.SetField(image.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := ic.mutation.Caption(); ok {
		_spec.SetField(image.FieldCaption, field.TypeString, value)
		_node.Caption = value
	}
	if value, ok := ic.mutation.Width(); ok {
		_spec.SetField(image.FieldWidth, field.TypeInt32, value)
		_node.Width = value
	}
	if value, ok := ic.mutation.High(); ok {
		_spec.SetField(image.FieldHigh, field.TypeInt32, value)
		_node.High = value
	}
	if value, ok := ic.mutation.SizeKB(); ok {
		_spec.SetField(image.FieldSizeKB, field.TypeFloat64, value)
		_node.SizeKB = value
	}
	if value, ok := ic.mutation.UploadedAt(); ok {
		_spec.SetField(image.FieldUploadedAt, field.TypeTime, value)
		_node.UploadedAt = value
	}
	if nodes := ic.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.UserTable,
			Columns: []string{image.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_image = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.CommentTable,
			Columns: []string{image.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.comment_image = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.BrandTable,
			Columns: []string{image.BrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.brand_image = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.ProductTable,
			Columns: []string{image.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_image = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ImageCreateBulk is the builder for creating many Image entities in bulk.
type ImageCreateBulk struct {
	config
	builders []*ImageCreate
}

// Save creates the Image entities in the database.
func (icb *ImageCreateBulk) Save(ctx context.Context) ([]*Image, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Image, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ImageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ImageCreateBulk) SaveX(ctx context.Context) []*Image {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ImageCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ImageCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
