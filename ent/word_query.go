// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/predicate"
	"shrektionary_api/ent/user"
	"shrektionary_api/ent/word"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WordQuery is the builder for querying Word entities.
type WordQuery struct {
	config
	ctx                  *QueryContext
	order                []word.OrderOption
	inters               []Interceptor
	predicates           []predicate.Word
	withCreator          *UserQuery
	withDefinitions      *DefinitionQuery
	withDescendants      *WordQuery
	withParent           *WordQuery
	withFKs              bool
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Word) error
	withNamedDefinitions map[string]*DefinitionQuery
	withNamedDescendants map[string]*WordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WordQuery builder.
func (wq *WordQuery) Where(ps ...predicate.Word) *WordQuery {
	wq.predicates = append(wq.predicates, ps...)
	return wq
}

// Limit the number of records to be returned by this query.
func (wq *WordQuery) Limit(limit int) *WordQuery {
	wq.ctx.Limit = &limit
	return wq
}

// Offset to start from.
func (wq *WordQuery) Offset(offset int) *WordQuery {
	wq.ctx.Offset = &offset
	return wq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wq *WordQuery) Unique(unique bool) *WordQuery {
	wq.ctx.Unique = &unique
	return wq
}

// Order specifies how the records should be ordered.
func (wq *WordQuery) Order(o ...word.OrderOption) *WordQuery {
	wq.order = append(wq.order, o...)
	return wq
}

// QueryCreator chains the current query on the "creator" edge.
func (wq *WordQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.CreatorTable, word.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDefinitions chains the current query on the "definitions" edge.
func (wq *WordQuery) QueryDefinitions() *DefinitionQuery {
	query := (&DefinitionClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(definition.Table, definition.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, word.DefinitionsTable, word.DefinitionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDescendants chains the current query on the "descendants" edge.
func (wq *WordQuery) QueryDescendants() *WordQuery {
	query := (&WordClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, word.DescendantsTable, word.DescendantsColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (wq *WordQuery) QueryParent() *WordQuery {
	query := (&WordClient{config: wq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(word.Table, word.FieldID, selector),
			sqlgraph.To(word.Table, word.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, word.ParentTable, word.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(wq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Word entity from the query.
// Returns a *NotFoundError when no Word was found.
func (wq *WordQuery) First(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(1).All(setContextOp(ctx, wq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{word.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wq *WordQuery) FirstX(ctx context.Context) *Word {
	node, err := wq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Word ID from the query.
// Returns a *NotFoundError when no Word ID was found.
func (wq *WordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(1).IDs(setContextOp(ctx, wq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{word.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wq *WordQuery) FirstIDX(ctx context.Context) int {
	id, err := wq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Word entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Word entity is found.
// Returns a *NotFoundError when no Word entities are found.
func (wq *WordQuery) Only(ctx context.Context) (*Word, error) {
	nodes, err := wq.Limit(2).All(setContextOp(ctx, wq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{word.Label}
	default:
		return nil, &NotSingularError{word.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wq *WordQuery) OnlyX(ctx context.Context) *Word {
	node, err := wq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Word ID in the query.
// Returns a *NotSingularError when more than one Word ID is found.
// Returns a *NotFoundError when no entities are found.
func (wq *WordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wq.Limit(2).IDs(setContextOp(ctx, wq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{word.Label}
	default:
		err = &NotSingularError{word.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wq *WordQuery) OnlyIDX(ctx context.Context) int {
	id, err := wq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Words.
func (wq *WordQuery) All(ctx context.Context) ([]*Word, error) {
	ctx = setContextOp(ctx, wq.ctx, "All")
	if err := wq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Word, *WordQuery]()
	return withInterceptors[[]*Word](ctx, wq, qr, wq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wq *WordQuery) AllX(ctx context.Context) []*Word {
	nodes, err := wq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Word IDs.
func (wq *WordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if wq.ctx.Unique == nil && wq.path != nil {
		wq.Unique(true)
	}
	ctx = setContextOp(ctx, wq.ctx, "IDs")
	if err = wq.Select(word.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wq *WordQuery) IDsX(ctx context.Context) []int {
	ids, err := wq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wq *WordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wq.ctx, "Count")
	if err := wq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wq, querierCount[*WordQuery](), wq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wq *WordQuery) CountX(ctx context.Context) int {
	count, err := wq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wq *WordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wq.ctx, "Exist")
	switch _, err := wq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wq *WordQuery) ExistX(ctx context.Context) bool {
	exist, err := wq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wq *WordQuery) Clone() *WordQuery {
	if wq == nil {
		return nil
	}
	return &WordQuery{
		config:          wq.config,
		ctx:             wq.ctx.Clone(),
		order:           append([]word.OrderOption{}, wq.order...),
		inters:          append([]Interceptor{}, wq.inters...),
		predicates:      append([]predicate.Word{}, wq.predicates...),
		withCreator:     wq.withCreator.Clone(),
		withDefinitions: wq.withDefinitions.Clone(),
		withDescendants: wq.withDescendants.Clone(),
		withParent:      wq.withParent.Clone(),
		// clone intermediate query.
		sql:  wq.sql.Clone(),
		path: wq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "creator" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithCreator(opts ...func(*UserQuery)) *WordQuery {
	query := (&UserClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withCreator = query
	return wq
}

// WithDefinitions tells the query-builder to eager-load the nodes that are connected to
// the "definitions" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithDefinitions(opts ...func(*DefinitionQuery)) *WordQuery {
	query := (&DefinitionClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withDefinitions = query
	return wq
}

// WithDescendants tells the query-builder to eager-load the nodes that are connected to
// the "descendants" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithDescendants(opts ...func(*WordQuery)) *WordQuery {
	query := (&WordClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withDescendants = query
	return wq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithParent(opts ...func(*WordQuery)) *WordQuery {
	query := (&WordClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wq.withParent = query
	return wq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Word.Query().
//		GroupBy(word.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wq *WordQuery) GroupBy(field string, fields ...string) *WordGroupBy {
	wq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WordGroupBy{build: wq}
	grbuild.flds = &wq.ctx.Fields
	grbuild.label = word.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.Word.Query().
//		Select(word.FieldDescription).
//		Scan(ctx, &v)
func (wq *WordQuery) Select(fields ...string) *WordSelect {
	wq.ctx.Fields = append(wq.ctx.Fields, fields...)
	sbuild := &WordSelect{WordQuery: wq}
	sbuild.label = word.Label
	sbuild.flds, sbuild.scan = &wq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WordSelect configured with the given aggregations.
func (wq *WordQuery) Aggregate(fns ...AggregateFunc) *WordSelect {
	return wq.Select().Aggregate(fns...)
}

func (wq *WordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wq); err != nil {
				return err
			}
		}
	}
	for _, f := range wq.ctx.Fields {
		if !word.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wq.path != nil {
		prev, err := wq.path(ctx)
		if err != nil {
			return err
		}
		wq.sql = prev
	}
	return nil
}

func (wq *WordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Word, error) {
	var (
		nodes       = []*Word{}
		withFKs     = wq.withFKs
		_spec       = wq.querySpec()
		loadedTypes = [4]bool{
			wq.withCreator != nil,
			wq.withDefinitions != nil,
			wq.withDescendants != nil,
			wq.withParent != nil,
		}
	)
	if wq.withCreator != nil || wq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, word.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Word).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Word{config: wq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wq.withCreator; query != nil {
		if err := wq.loadCreator(ctx, query, nodes, nil,
			func(n *Word, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	if query := wq.withDefinitions; query != nil {
		if err := wq.loadDefinitions(ctx, query, nodes,
			func(n *Word) { n.Edges.Definitions = []*Definition{} },
			func(n *Word, e *Definition) { n.Edges.Definitions = append(n.Edges.Definitions, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withDescendants; query != nil {
		if err := wq.loadDescendants(ctx, query, nodes,
			func(n *Word) { n.Edges.Descendants = []*Word{} },
			func(n *Word, e *Word) { n.Edges.Descendants = append(n.Edges.Descendants, e) }); err != nil {
			return nil, err
		}
	}
	if query := wq.withParent; query != nil {
		if err := wq.loadParent(ctx, query, nodes, nil,
			func(n *Word, e *Word) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedDefinitions {
		if err := wq.loadDefinitions(ctx, query, nodes,
			func(n *Word) { n.appendNamedDefinitions(name) },
			func(n *Word, e *Definition) { n.appendNamedDefinitions(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range wq.withNamedDescendants {
		if err := wq.loadDescendants(ctx, query, nodes,
			func(n *Word) { n.appendNamedDescendants(name) },
			func(n *Word, e *Word) { n.appendNamedDescendants(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range wq.loadTotal {
		if err := wq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wq *WordQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*Word, init func(*Word), assign func(*Word, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Word)
	for i := range nodes {
		if nodes[i].user_words == nil {
			continue
		}
		fk := *nodes[i].user_words
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_words" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wq *WordQuery) loadDefinitions(ctx context.Context, query *DefinitionQuery, nodes []*Word, init func(*Word), assign func(*Word, *Definition)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Word)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Definition(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(word.DefinitionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.word_definitions
		if fk == nil {
			return fmt.Errorf(`foreign-key "word_definitions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "word_definitions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WordQuery) loadDescendants(ctx context.Context, query *WordQuery, nodes []*Word, init func(*Word), assign func(*Word, *Word)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Word)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Word(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(word.DescendantsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.word_descendants
		if fk == nil {
			return fmt.Errorf(`foreign-key "word_descendants" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "word_descendants" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (wq *WordQuery) loadParent(ctx context.Context, query *WordQuery, nodes []*Word, init func(*Word), assign func(*Word, *Word)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Word)
	for i := range nodes {
		if nodes[i].word_descendants == nil {
			continue
		}
		fk := *nodes[i].word_descendants
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(word.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "word_descendants" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wq *WordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wq.querySpec()
	if len(wq.modifiers) > 0 {
		_spec.Modifiers = wq.modifiers
	}
	_spec.Node.Columns = wq.ctx.Fields
	if len(wq.ctx.Fields) > 0 {
		_spec.Unique = wq.ctx.Unique != nil && *wq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wq.driver, _spec)
}

func (wq *WordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(word.Table, word.Columns, sqlgraph.NewFieldSpec(word.FieldID, field.TypeInt))
	_spec.From = wq.sql
	if unique := wq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wq.path != nil {
		_spec.Unique = true
	}
	if fields := wq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, word.FieldID)
		for i := range fields {
			if fields[i] != word.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wq *WordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wq.driver.Dialect())
	t1 := builder.Table(word.Table)
	columns := wq.ctx.Fields
	if len(columns) == 0 {
		columns = word.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wq.sql != nil {
		selector = wq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wq.ctx.Unique != nil && *wq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wq.predicates {
		p(selector)
	}
	for _, p := range wq.order {
		p(selector)
	}
	if offset := wq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedDefinitions tells the query-builder to eager-load the nodes that are connected to the "definitions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithNamedDefinitions(name string, opts ...func(*DefinitionQuery)) *WordQuery {
	query := (&DefinitionClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedDefinitions == nil {
		wq.withNamedDefinitions = make(map[string]*DefinitionQuery)
	}
	wq.withNamedDefinitions[name] = query
	return wq
}

// WithNamedDescendants tells the query-builder to eager-load the nodes that are connected to the "descendants"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (wq *WordQuery) WithNamedDescendants(name string, opts ...func(*WordQuery)) *WordQuery {
	query := (&WordClient{config: wq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if wq.withNamedDescendants == nil {
		wq.withNamedDescendants = make(map[string]*WordQuery)
	}
	wq.withNamedDescendants[name] = query
	return wq
}

// WordGroupBy is the group-by builder for Word entities.
type WordGroupBy struct {
	selector
	build *WordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wgb *WordGroupBy) Aggregate(fns ...AggregateFunc) *WordGroupBy {
	wgb.fns = append(wgb.fns, fns...)
	return wgb
}

// Scan applies the selector query and scans the result into the given value.
func (wgb *WordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wgb.build.ctx, "GroupBy")
	if err := wgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WordQuery, *WordGroupBy](ctx, wgb.build, wgb, wgb.build.inters, v)
}

func (wgb *WordGroupBy) sqlScan(ctx context.Context, root *WordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wgb.fns))
	for _, fn := range wgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wgb.flds)+len(wgb.fns))
		for _, f := range *wgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WordSelect is the builder for selecting fields of Word entities.
type WordSelect struct {
	*WordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ws *WordSelect) Aggregate(fns ...AggregateFunc) *WordSelect {
	ws.fns = append(ws.fns, fns...)
	return ws
}

// Scan applies the selector query and scans the result into the given value.
func (ws *WordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ws.ctx, "Select")
	if err := ws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WordQuery, *WordSelect](ctx, ws.WordQuery, ws, ws.inters, v)
}

func (ws *WordSelect) sqlScan(ctx context.Context, root *WordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ws.fns))
	for _, fn := range ws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
