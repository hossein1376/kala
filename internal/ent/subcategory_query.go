// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"kala/internal/ent/category"
	"kala/internal/ent/image"
	"kala/internal/ent/predicate"
	"kala/internal/ent/product"
	"kala/internal/ent/subcategory"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubCategoryQuery is the builder for querying SubCategory entities.
type SubCategoryQuery struct {
	config
	ctx          *QueryContext
	order        []subcategory.OrderOption
	inters       []Interceptor
	predicates   []predicate.SubCategory
	withImage    *ImageQuery
	withProduct  *ProductQuery
	withCategory *CategoryQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubCategoryQuery builder.
func (scq *SubCategoryQuery) Where(ps ...predicate.SubCategory) *SubCategoryQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *SubCategoryQuery) Limit(limit int) *SubCategoryQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *SubCategoryQuery) Offset(offset int) *SubCategoryQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *SubCategoryQuery) Unique(unique bool) *SubCategoryQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *SubCategoryQuery) Order(o ...subcategory.OrderOption) *SubCategoryQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryImage chains the current query on the "image" edge.
func (scq *SubCategoryQuery) QueryImage() *ImageQuery {
	query := (&ImageClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, subcategory.ImageTable, subcategory.ImagePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the "product" edge.
func (scq *SubCategoryQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, subcategory.ProductTable, subcategory.ProductPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the "category" edge.
func (scq *SubCategoryQuery) QueryCategory() *CategoryQuery {
	query := (&CategoryClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subcategory.Table, subcategory.FieldID, selector),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subcategory.CategoryTable, subcategory.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SubCategory entity from the query.
// Returns a *NotFoundError when no SubCategory was found.
func (scq *SubCategoryQuery) First(ctx context.Context) (*SubCategory, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *SubCategoryQuery) FirstX(ctx context.Context) *SubCategory {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SubCategory ID from the query.
// Returns a *NotFoundError when no SubCategory ID was found.
func (scq *SubCategoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *SubCategoryQuery) FirstIDX(ctx context.Context) int {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SubCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SubCategory entity is found.
// Returns a *NotFoundError when no SubCategory entities are found.
func (scq *SubCategoryQuery) Only(ctx context.Context) (*SubCategory, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subcategory.Label}
	default:
		return nil, &NotSingularError{subcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *SubCategoryQuery) OnlyX(ctx context.Context) *SubCategory {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SubCategory ID in the query.
// Returns a *NotSingularError when more than one SubCategory ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *SubCategoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subcategory.Label}
	default:
		err = &NotSingularError{subcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *SubCategoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SubCategories.
func (scq *SubCategoryQuery) All(ctx context.Context) ([]*SubCategory, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SubCategory, *SubCategoryQuery]()
	return withInterceptors[[]*SubCategory](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *SubCategoryQuery) AllX(ctx context.Context) []*SubCategory {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SubCategory IDs.
func (scq *SubCategoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(subcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *SubCategoryQuery) IDsX(ctx context.Context) []int {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *SubCategoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*SubCategoryQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *SubCategoryQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *SubCategoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *SubCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *SubCategoryQuery) Clone() *SubCategoryQuery {
	if scq == nil {
		return nil
	}
	return &SubCategoryQuery{
		config:       scq.config,
		ctx:          scq.ctx.Clone(),
		order:        append([]subcategory.OrderOption{}, scq.order...),
		inters:       append([]Interceptor{}, scq.inters...),
		predicates:   append([]predicate.SubCategory{}, scq.predicates...),
		withImage:    scq.withImage.Clone(),
		withProduct:  scq.withProduct.Clone(),
		withCategory: scq.withCategory.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithImage tells the query-builder to eager-load the nodes that are connected to
// the "image" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *SubCategoryQuery) WithImage(opts ...func(*ImageQuery)) *SubCategoryQuery {
	query := (&ImageClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withImage = query
	return scq
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *SubCategoryQuery) WithProduct(opts ...func(*ProductQuery)) *SubCategoryQuery {
	query := (&ProductClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withProduct = query
	return scq
}

// WithCategory tells the query-builder to eager-load the nodes that are connected to
// the "category" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *SubCategoryQuery) WithCategory(opts ...func(*CategoryQuery)) *SubCategoryQuery {
	query := (&CategoryClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withCategory = query
	return scq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SubCategory.Query().
//		GroupBy(subcategory.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *SubCategoryQuery) GroupBy(field string, fields ...string) *SubCategoryGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubCategoryGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = subcategory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.SubCategory.Query().
//		Select(subcategory.FieldCreateTime).
//		Scan(ctx, &v)
func (scq *SubCategoryQuery) Select(fields ...string) *SubCategorySelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &SubCategorySelect{SubCategoryQuery: scq}
	sbuild.label = subcategory.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubCategorySelect configured with the given aggregations.
func (scq *SubCategoryQuery) Aggregate(fns ...AggregateFunc) *SubCategorySelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *SubCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !subcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *SubCategoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SubCategory, error) {
	var (
		nodes       = []*SubCategory{}
		withFKs     = scq.withFKs
		_spec       = scq.querySpec()
		loadedTypes = [3]bool{
			scq.withImage != nil,
			scq.withProduct != nil,
			scq.withCategory != nil,
		}
	)
	if scq.withCategory != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subcategory.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SubCategory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SubCategory{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withImage; query != nil {
		if err := scq.loadImage(ctx, query, nodes,
			func(n *SubCategory) { n.Edges.Image = []*Image{} },
			func(n *SubCategory, e *Image) { n.Edges.Image = append(n.Edges.Image, e) }); err != nil {
			return nil, err
		}
	}
	if query := scq.withProduct; query != nil {
		if err := scq.loadProduct(ctx, query, nodes,
			func(n *SubCategory) { n.Edges.Product = []*Product{} },
			func(n *SubCategory, e *Product) { n.Edges.Product = append(n.Edges.Product, e) }); err != nil {
			return nil, err
		}
	}
	if query := scq.withCategory; query != nil {
		if err := scq.loadCategory(ctx, query, nodes, nil,
			func(n *SubCategory, e *Category) { n.Edges.Category = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *SubCategoryQuery) loadImage(ctx context.Context, query *ImageQuery, nodes []*SubCategory, init func(*SubCategory), assign func(*SubCategory, *Image)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*SubCategory)
	nids := make(map[int]map[*SubCategory]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(subcategory.ImageTable)
		s.Join(joinT).On(s.C(image.FieldID), joinT.C(subcategory.ImagePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(subcategory.ImagePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(subcategory.ImagePrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*SubCategory]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Image](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "image" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (scq *SubCategoryQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*SubCategory, init func(*SubCategory), assign func(*SubCategory, *Product)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*SubCategory)
	nids := make(map[int]map[*SubCategory]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(subcategory.ProductTable)
		s.Join(joinT).On(s.C(product.FieldID), joinT.C(subcategory.ProductPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(subcategory.ProductPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(subcategory.ProductPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*SubCategory]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Product](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "product" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (scq *SubCategoryQuery) loadCategory(ctx context.Context, query *CategoryQuery, nodes []*SubCategory, init func(*SubCategory), assign func(*SubCategory, *Category)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SubCategory)
	for i := range nodes {
		if nodes[i].category == nil {
			continue
		}
		fk := *nodes[i].category
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(category.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "category" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *SubCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *SubCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subcategory.Table, subcategory.Columns, sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeInt))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subcategory.FieldID)
		for i := range fields {
			if fields[i] != subcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *SubCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(subcategory.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = subcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubCategoryGroupBy is the group-by builder for SubCategory entities.
type SubCategoryGroupBy struct {
	selector
	build *SubCategoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *SubCategoryGroupBy) Aggregate(fns ...AggregateFunc) *SubCategoryGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *SubCategoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubCategoryQuery, *SubCategoryGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *SubCategoryGroupBy) sqlScan(ctx context.Context, root *SubCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubCategorySelect is the builder for selecting fields of SubCategory entities.
type SubCategorySelect struct {
	*SubCategoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *SubCategorySelect) Aggregate(fns ...AggregateFunc) *SubCategorySelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *SubCategorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubCategoryQuery, *SubCategorySelect](ctx, scs.SubCategoryQuery, scs, scs.inters, v)
}

func (scs *SubCategorySelect) sqlScan(ctx context.Context, root *SubCategoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
