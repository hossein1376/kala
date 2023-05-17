// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/attribute"
	"kala/internal/ent/attributevalue"
	"kala/internal/ent/predicate"
	"kala/internal/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeValueUpdate is the builder for updating AttributeValue entities.
type AttributeValueUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeValueMutation
}

// Where appends a list predicates to the AttributeValueUpdate builder.
func (avu *AttributeValueUpdate) Where(ps ...predicate.AttributeValue) *AttributeValueUpdate {
	avu.mutation.Where(ps...)
	return avu
}

// SetValue sets the "value" field.
func (avu *AttributeValueUpdate) SetValue(s string) *AttributeValueUpdate {
	avu.mutation.SetValue(s)
	return avu
}

// SetAttribute sets the "attribute" field.
func (avu *AttributeValueUpdate) SetAttribute(i int) *AttributeValueUpdate {
	avu.mutation.SetAttribute(i)
	return avu
}

// SetProduct sets the "product" field.
func (avu *AttributeValueUpdate) SetProduct(i int) *AttributeValueUpdate {
	avu.mutation.SetProduct(i)
	return avu
}

// SetAttributesID sets the "attributes" edge to the Attribute entity by ID.
func (avu *AttributeValueUpdate) SetAttributesID(id int) *AttributeValueUpdate {
	avu.mutation.SetAttributesID(id)
	return avu
}

// SetAttributes sets the "attributes" edge to the Attribute entity.
func (avu *AttributeValueUpdate) SetAttributes(a *Attribute) *AttributeValueUpdate {
	return avu.SetAttributesID(a.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (avu *AttributeValueUpdate) SetProductsID(id int) *AttributeValueUpdate {
	avu.mutation.SetProductsID(id)
	return avu
}

// SetProducts sets the "products" edge to the Product entity.
func (avu *AttributeValueUpdate) SetProducts(p *Product) *AttributeValueUpdate {
	return avu.SetProductsID(p.ID)
}

// Mutation returns the AttributeValueMutation object of the builder.
func (avu *AttributeValueUpdate) Mutation() *AttributeValueMutation {
	return avu.mutation
}

// ClearAttributes clears the "attributes" edge to the Attribute entity.
func (avu *AttributeValueUpdate) ClearAttributes() *AttributeValueUpdate {
	avu.mutation.ClearAttributes()
	return avu
}

// ClearProducts clears the "products" edge to the Product entity.
func (avu *AttributeValueUpdate) ClearProducts() *AttributeValueUpdate {
	avu.mutation.ClearProducts()
	return avu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (avu *AttributeValueUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, avu.sqlSave, avu.mutation, avu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (avu *AttributeValueUpdate) SaveX(ctx context.Context) int {
	affected, err := avu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (avu *AttributeValueUpdate) Exec(ctx context.Context) error {
	_, err := avu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avu *AttributeValueUpdate) ExecX(ctx context.Context) {
	if err := avu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (avu *AttributeValueUpdate) check() error {
	if v, ok := avu.mutation.Value(); ok {
		if err := attributevalue.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AttributeValue.value": %w`, err)}
		}
	}
	if _, ok := avu.mutation.AttributesID(); avu.mutation.AttributesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttributeValue.attributes"`)
	}
	if _, ok := avu.mutation.ProductsID(); avu.mutation.ProductsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttributeValue.products"`)
	}
	return nil
}

func (avu *AttributeValueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := avu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attributevalue.Table, attributevalue.Columns, sqlgraph.NewFieldSpec(attributevalue.FieldID, field.TypeInt))
	if ps := avu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := avu.mutation.Value(); ok {
		_spec.SetField(attributevalue.FieldValue, field.TypeString, value)
	}
	if avu.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.AttributesTable,
			Columns: []string{attributevalue.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := avu.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.AttributesTable,
			Columns: []string{attributevalue.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if avu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.ProductsTable,
			Columns: []string{attributevalue.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := avu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.ProductsTable,
			Columns: []string{attributevalue.ProductsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, avu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attributevalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	avu.mutation.done = true
	return n, nil
}

// AttributeValueUpdateOne is the builder for updating a single AttributeValue entity.
type AttributeValueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeValueMutation
}

// SetValue sets the "value" field.
func (avuo *AttributeValueUpdateOne) SetValue(s string) *AttributeValueUpdateOne {
	avuo.mutation.SetValue(s)
	return avuo
}

// SetAttribute sets the "attribute" field.
func (avuo *AttributeValueUpdateOne) SetAttribute(i int) *AttributeValueUpdateOne {
	avuo.mutation.SetAttribute(i)
	return avuo
}

// SetProduct sets the "product" field.
func (avuo *AttributeValueUpdateOne) SetProduct(i int) *AttributeValueUpdateOne {
	avuo.mutation.SetProduct(i)
	return avuo
}

// SetAttributesID sets the "attributes" edge to the Attribute entity by ID.
func (avuo *AttributeValueUpdateOne) SetAttributesID(id int) *AttributeValueUpdateOne {
	avuo.mutation.SetAttributesID(id)
	return avuo
}

// SetAttributes sets the "attributes" edge to the Attribute entity.
func (avuo *AttributeValueUpdateOne) SetAttributes(a *Attribute) *AttributeValueUpdateOne {
	return avuo.SetAttributesID(a.ID)
}

// SetProductsID sets the "products" edge to the Product entity by ID.
func (avuo *AttributeValueUpdateOne) SetProductsID(id int) *AttributeValueUpdateOne {
	avuo.mutation.SetProductsID(id)
	return avuo
}

// SetProducts sets the "products" edge to the Product entity.
func (avuo *AttributeValueUpdateOne) SetProducts(p *Product) *AttributeValueUpdateOne {
	return avuo.SetProductsID(p.ID)
}

// Mutation returns the AttributeValueMutation object of the builder.
func (avuo *AttributeValueUpdateOne) Mutation() *AttributeValueMutation {
	return avuo.mutation
}

// ClearAttributes clears the "attributes" edge to the Attribute entity.
func (avuo *AttributeValueUpdateOne) ClearAttributes() *AttributeValueUpdateOne {
	avuo.mutation.ClearAttributes()
	return avuo
}

// ClearProducts clears the "products" edge to the Product entity.
func (avuo *AttributeValueUpdateOne) ClearProducts() *AttributeValueUpdateOne {
	avuo.mutation.ClearProducts()
	return avuo
}

// Where appends a list predicates to the AttributeValueUpdate builder.
func (avuo *AttributeValueUpdateOne) Where(ps ...predicate.AttributeValue) *AttributeValueUpdateOne {
	avuo.mutation.Where(ps...)
	return avuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (avuo *AttributeValueUpdateOne) Select(field string, fields ...string) *AttributeValueUpdateOne {
	avuo.fields = append([]string{field}, fields...)
	return avuo
}

// Save executes the query and returns the updated AttributeValue entity.
func (avuo *AttributeValueUpdateOne) Save(ctx context.Context) (*AttributeValue, error) {
	return withHooks(ctx, avuo.sqlSave, avuo.mutation, avuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (avuo *AttributeValueUpdateOne) SaveX(ctx context.Context) *AttributeValue {
	node, err := avuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (avuo *AttributeValueUpdateOne) Exec(ctx context.Context) error {
	_, err := avuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (avuo *AttributeValueUpdateOne) ExecX(ctx context.Context) {
	if err := avuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (avuo *AttributeValueUpdateOne) check() error {
	if v, ok := avuo.mutation.Value(); ok {
		if err := attributevalue.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AttributeValue.value": %w`, err)}
		}
	}
	if _, ok := avuo.mutation.AttributesID(); avuo.mutation.AttributesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttributeValue.attributes"`)
	}
	if _, ok := avuo.mutation.ProductsID(); avuo.mutation.ProductsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AttributeValue.products"`)
	}
	return nil
}

func (avuo *AttributeValueUpdateOne) sqlSave(ctx context.Context) (_node *AttributeValue, err error) {
	if err := avuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attributevalue.Table, attributevalue.Columns, sqlgraph.NewFieldSpec(attributevalue.FieldID, field.TypeInt))
	id, ok := avuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AttributeValue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := avuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attributevalue.FieldID)
		for _, f := range fields {
			if !attributevalue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attributevalue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := avuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := avuo.mutation.Value(); ok {
		_spec.SetField(attributevalue.FieldValue, field.TypeString, value)
	}
	if avuo.mutation.AttributesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.AttributesTable,
			Columns: []string{attributevalue.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := avuo.mutation.AttributesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.AttributesTable,
			Columns: []string{attributevalue.AttributesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if avuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.ProductsTable,
			Columns: []string{attributevalue.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := avuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attributevalue.ProductsTable,
			Columns: []string{attributevalue.ProductsColumn},
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
	_node = &AttributeValue{config: avuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, avuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attributevalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	avuo.mutation.done = true
	return _node, nil
}
