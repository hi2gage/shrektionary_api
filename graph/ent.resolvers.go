package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"shrektionary_api/ent"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Definitions is the resolver for the definitions field.
func (r *queryResolver) Definitions(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.DefinitionOrder, where *ent.DefinitionWhereInput) (*ent.DefinitionConnection, error) {
	return r.client.Definition.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithDefinitionOrder(ent.DefaultDefinitionOrder),
			ent.WithDefinitionFilter(where.Filter),
		)
}

// Groups is the resolver for the groups field.
func (r *queryResolver) Groups(ctx context.Context) ([]*ent.Group, error) {
	return r.client.Group.Query().All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.client.User.Query().All(ctx)
}

// Words is the resolver for the words field.
func (r *queryResolver) Words(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.WordWhereInput) (*ent.WordConnection, error) {
	return r.client.Word.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithWordOrder(ent.DefaultWordOrder),
			ent.WithWordFilter(where.Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
