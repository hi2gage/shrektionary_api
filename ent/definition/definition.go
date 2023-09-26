// Code generated by ent, DO NOT EDIT.

package definition

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the definition type in the database.
	Label = "definition"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeWord holds the string denoting the word edge name in mutations.
	EdgeWord = "word"
	// Table holds the table name of the definition in the database.
	Table = "definitions"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "definitions"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "user_definitions"
	// WordTable is the table that holds the word relation/edge.
	WordTable = "definitions"
	// WordInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	WordInverseTable = "words"
	// WordColumn is the table column denoting the word relation/edge.
	WordColumn = "word_definitions"
)

// Columns holds all SQL columns for definition fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "definitions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_definitions",
	"word_definitions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

// OrderOption defines the ordering options for the Definition queries.
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

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByWordField orders the results by word field.
func ByWordField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWordStep(), sql.OrderByField(field, opts...))
	}
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
	)
}
func newWordStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WordInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, WordTable, WordColumn),
	)
}
