package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"shrektionary_api/ent"
	"shrektionary_api/utils"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*ent.Group, error) {
	return r.client.Group.Create().SetInput(input).Save(ctx)
}

// CreateDefinition is the resolver for the createDefinition field.
func (r *mutationResolver) CreateDefinition(ctx context.Context, input ent.CreateDefinitionInput) (*ent.Definition, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Definition.Create().SetInput(input).Save(ctx)
}

// CreateWord is the resolver for the createWord field.
func (r *mutationResolver) CreateWord(ctx context.Context, input ent.CreateWordInput) (*ent.Word, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Word.Create().SetInput(input).Save(ctx)
}

// CreateRootWord is the resolver for the createRootWord field.
func (r *mutationResolver) CreateRootWord(ctx context.Context, input ent.CreateWordInput) (*ent.Word, error) {
	err := input.ValidateCreateInput()
	if err != nil {
		return nil, err
	}

	input.SetCreatorID(ctx)

	return r.client.Word.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input ent.UpdateUserInput) (*ent.User, error) {
	creatorID, err := utils.GetCreatorIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.User.UpdateOneID(creatorID).SetInput(input).Save(ctx)
}

// UpdateWord is the resolver for the updateWord field.
func (r *mutationResolver) UpdateWord(ctx context.Context, id int, input ent.UpdateWordInput) (*ent.Word, error) {
	return r.client.Word.UpdateOneID(id).SetInput(input).Save(ctx)
}

// UpdateDefinition is the resolver for the updateDefinition field.
func (r *mutationResolver) UpdateDefinition(ctx context.Context, id int, input ent.UpdateDefinitionInput) (*ent.Definition, error) {
	return r.client.Definition.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteWord is the resolver for the deleteWord field.
func (r *mutationResolver) DeleteWord(ctx context.Context, id int) (bool, error) {
	if err := r.client.Word.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// DeleteDefinition is the resolver for the deleteDefinition field.
func (r *mutationResolver) DeleteDefinition(ctx context.Context, id int) (bool, error) {
	if err := r.client.Definition.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
