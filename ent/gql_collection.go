// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/group"
	"shrektionary_api/ent/user"
	"shrektionary_api/ent/word"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DefinitionQuery) CollectFields(ctx context.Context, satisfies ...string) (*DefinitionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return d, nil
	}
	if err := d.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DefinitionQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(definition.Columns))
		selectedFields = []string{definition.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "creator":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: d.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			d.withCreator = query
		case "description":
			if _, ok := fieldSeen[definition.FieldDescription]; !ok {
				selectedFields = append(selectedFields, definition.FieldDescription)
				fieldSeen[definition.FieldDescription] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		d.Select(selectedFields...)
	}
	return nil
}

type definitionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DefinitionPaginateOption
}

func newDefinitionPaginateArgs(rv map[string]any) *definitionPaginateArgs {
	args := &definitionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &DefinitionOrder{Field: &DefinitionOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithDefinitionOrder(order))
			}
		case *DefinitionOrder:
			if v != nil {
				args.opts = append(args.opts, WithDefinitionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*DefinitionWhereInput); ok {
		args.opts = append(args.opts, WithDefinitionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(group.Columns))
		selectedFields = []string{group.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "rootwords":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&WordClient{config: gr.config}).Query()
			)
			args := newWordPaginateArgs(fieldArgs(ctx, new(WordWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newWordPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					gr.loadTotal = append(gr.loadTotal, func(ctx context.Context, nodes []*Group) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"group_root_words"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(group.RootWordsColumn), ids...))
						})
						if err := query.GroupBy(group.RootWordsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					gr.loadTotal = append(gr.loadTotal, func(_ context.Context, nodes []*Group) error {
						for i := range nodes {
							n := len(nodes[i].Edges.RootWords)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Word")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(group.RootWordsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			gr.WithNamedRootWords(alias, func(wq *WordQuery) {
				*wq = *query
			})
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: gr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			gr.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "description":
			if _, ok := fieldSeen[group.FieldDescription]; !ok {
				selectedFields = append(selectedFields, group.FieldDescription)
				fieldSeen[group.FieldDescription] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		gr.Select(selectedFields...)
	}
	return nil
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]any) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "groups":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GroupClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedGroups(alias, func(wq *GroupQuery) {
				*wq = *query
			})
		case "definitions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DefinitionClient{config: u.config}).Query()
			)
			args := newDefinitionPaginateArgs(fieldArgs(ctx, new(DefinitionWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newDefinitionPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_definitions"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(user.DefinitionsColumn), ids...))
						})
						if err := query.GroupBy(user.DefinitionsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Definitions)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Definition")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.DefinitionsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			u.WithNamedDefinitions(alias, func(wq *DefinitionQuery) {
				*wq = *query
			})
		case "words":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&WordClient{config: u.config}).Query()
			)
			args := newWordPaginateArgs(fieldArgs(ctx, new(WordWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newWordPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"user_words"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(user.WordsColumn), ids...))
						})
						if err := query.GroupBy(user.WordsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Words)
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Word")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(user.WordsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			u.WithNamedWords(alias, func(wq *WordQuery) {
				*wq = *query
			})
		case "authid":
			if _, ok := fieldSeen[user.FieldAuthID]; !ok {
				selectedFields = append(selectedFields, user.FieldAuthID)
				fieldSeen[user.FieldAuthID] = struct{}{}
			}
		case "firstname":
			if _, ok := fieldSeen[user.FieldFirstName]; !ok {
				selectedFields = append(selectedFields, user.FieldFirstName)
				fieldSeen[user.FieldFirstName] = struct{}{}
			}
		case "lastname":
			if _, ok := fieldSeen[user.FieldLastName]; !ok {
				selectedFields = append(selectedFields, user.FieldLastName)
				fieldSeen[user.FieldLastName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (w *WordQuery) CollectFields(ctx context.Context, satisfies ...string) (*WordQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return w, nil
	}
	if err := w.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *WordQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(word.Columns))
		selectedFields = []string{word.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "creator":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: w.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			w.withCreator = query
		case "group":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GroupClient{config: w.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			w.withGroup = query
		case "definitions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DefinitionClient{config: w.config}).Query()
			)
			args := newDefinitionPaginateArgs(fieldArgs(ctx, new(DefinitionWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newDefinitionPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					w.loadTotal = append(w.loadTotal, func(ctx context.Context, nodes []*Word) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"word_definitions"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(word.DefinitionsColumn), ids...))
						})
						if err := query.GroupBy(word.DefinitionsColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				} else {
					w.loadTotal = append(w.loadTotal, func(_ context.Context, nodes []*Word) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Definitions)
							if nodes[i].Edges.totalCount[2] == nil {
								nodes[i].Edges.totalCount[2] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[2][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Definition")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(word.DefinitionsColumn, limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			w.WithNamedDefinitions(alias, func(wq *DefinitionQuery) {
				*wq = *query
			})
		case "descendants":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&WordClient{config: w.config}).Query()
			)
			args := newWordPaginateArgs(fieldArgs(ctx, new(WordWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newWordPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					w.loadTotal = append(w.loadTotal, func(ctx context.Context, nodes []*Word) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"word_id"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(word.DescendantsTable)
							s.Join(joinT).On(s.C(word.FieldID), joinT.C(word.DescendantsPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(word.DescendantsPrimaryKey[0]), ids...))
							s.Select(joinT.C(word.DescendantsPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(word.DescendantsPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[3] == nil {
								nodes[i].Edges.totalCount[3] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[3][alias] = n
						}
						return nil
					})
				} else {
					w.loadTotal = append(w.loadTotal, func(_ context.Context, nodes []*Word) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Descendants)
							if nodes[i].Edges.totalCount[3] == nil {
								nodes[i].Edges.totalCount[3] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[3][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, opCtx, *field, path, mayAddCondition(satisfies, "Word")...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				modify := limitRows(word.DescendantsPrimaryKey[0], limit, pager.orderExpr(query))
				query.modifiers = append(query.modifiers, modify)
			} else {
				query = pager.applyOrder(query)
			}
			w.WithNamedDescendants(alias, func(wq *WordQuery) {
				*wq = *query
			})
		case "parents":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&WordClient{config: w.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			w.WithNamedParents(alias, func(wq *WordQuery) {
				*wq = *query
			})
		case "description":
			if _, ok := fieldSeen[word.FieldDescription]; !ok {
				selectedFields = append(selectedFields, word.FieldDescription)
				fieldSeen[word.FieldDescription] = struct{}{}
			}
		case "descendantcount":
			if _, ok := fieldSeen[word.FieldDescendantCount]; !ok {
				selectedFields = append(selectedFields, word.FieldDescendantCount)
				fieldSeen[word.FieldDescendantCount] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		w.Select(selectedFields...)
	}
	return nil
}

type wordPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []WordPaginateOption
}

func newWordPaginateArgs(rv map[string]any) *wordPaginateArgs {
	args := &wordPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*WordWhereInput); ok {
		args.opts = append(args.opts, WithWordFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
