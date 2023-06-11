// Code generated by ent, DO NOT EDIT.

package ent

// CreateDefinitionInput represents a mutation input for creating definitions.
type CreateDefinitionInput struct {
	Description string
	WordID      *int
}

// Mutate applies the CreateDefinitionInput on the DefinitionMutation builder.
func (i *CreateDefinitionInput) Mutate(m *DefinitionMutation) {
	m.SetDescription(i.Description)
	if v := i.WordID; v != nil {
		m.SetWordID(*v)
	}
}

// SetInput applies the change-set in the CreateDefinitionInput on the DefinitionCreate builder.
func (c *DefinitionCreate) SetInput(i CreateDefinitionInput) *DefinitionCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	Title string
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	m.SetTitle(i.Title)
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateWordInput represents a mutation input for creating words.
type CreateWordInput struct {
	Description   string
	DefinitionIDs []int
}

// Mutate applies the CreateWordInput on the WordMutation builder.
func (i *CreateWordInput) Mutate(m *WordMutation) {
	m.SetDescription(i.Description)
	if v := i.DefinitionIDs; len(v) > 0 {
		m.AddDefinitionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateWordInput on the WordCreate builder.
func (c *WordCreate) SetInput(i CreateWordInput) *WordCreate {
	i.Mutate(c.Mutation())
	return c
}