// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"shrektionary_api/ent/group"
	"shrektionary_api/ent/user"
	"shrektionary_api/ent/word"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Word is the model entity for the Word schema.
type Word struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// DescendantCount holds the value of the "descendantCount" field.
	DescendantCount int `json:"descendantCount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WordQuery when eager-loading is set.
	Edges            WordEdges `json:"edges"`
	group_root_words *int
	user_words       *int
	selectValues     sql.SelectValues
}

// WordEdges holds the relations/edges for other nodes in the graph.
type WordEdges struct {
	// Creator holds the value of the creator edge.
	Creator *User `json:"creator,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Definitions holds the value of the definitions edge.
	Definitions []*Definition `json:"definitions,omitempty"`
	// Descendants holds the value of the descendants edge.
	Descendants []*Word `json:"descendants,omitempty"`
	// Parents holds the value of the parents edge.
	Parents []*Word `json:"parents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedDefinitions map[string][]*Definition
	namedDescendants map[string][]*Word
	namedParents     map[string][]*Word
}

// CreatorOrErr returns the Creator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WordEdges) CreatorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Creator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Creator, nil
	}
	return nil, &NotLoadedError{edge: "creator"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WordEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// DefinitionsOrErr returns the Definitions value or an error if the edge
// was not loaded in eager-loading.
func (e WordEdges) DefinitionsOrErr() ([]*Definition, error) {
	if e.loadedTypes[2] {
		return e.Definitions, nil
	}
	return nil, &NotLoadedError{edge: "definitions"}
}

// DescendantsOrErr returns the Descendants value or an error if the edge
// was not loaded in eager-loading.
func (e WordEdges) DescendantsOrErr() ([]*Word, error) {
	if e.loadedTypes[3] {
		return e.Descendants, nil
	}
	return nil, &NotLoadedError{edge: "descendants"}
}

// ParentsOrErr returns the Parents value or an error if the edge
// was not loaded in eager-loading.
func (e WordEdges) ParentsOrErr() ([]*Word, error) {
	if e.loadedTypes[4] {
		return e.Parents, nil
	}
	return nil, &NotLoadedError{edge: "parents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Word) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case word.FieldID, word.FieldDescendantCount:
			values[i] = new(sql.NullInt64)
		case word.FieldDescription:
			values[i] = new(sql.NullString)
		case word.FieldCreateTime, word.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case word.ForeignKeys[0]: // group_root_words
			values[i] = new(sql.NullInt64)
		case word.ForeignKeys[1]: // user_words
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Word fields.
func (w *Word) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case word.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case word.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				w.CreateTime = value.Time
			}
		case word.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				w.UpdateTime = value.Time
			}
		case word.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				w.Description = value.String
			}
		case word.FieldDescendantCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field descendantCount", values[i])
			} else if value.Valid {
				w.DescendantCount = int(value.Int64)
			}
		case word.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_root_words", value)
			} else if value.Valid {
				w.group_root_words = new(int)
				*w.group_root_words = int(value.Int64)
			}
		case word.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_words", value)
			} else if value.Valid {
				w.user_words = new(int)
				*w.user_words = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Word.
// This includes values selected through modifiers, order, etc.
func (w *Word) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryCreator queries the "creator" edge of the Word entity.
func (w *Word) QueryCreator() *UserQuery {
	return NewWordClient(w.config).QueryCreator(w)
}

// QueryGroup queries the "group" edge of the Word entity.
func (w *Word) QueryGroup() *GroupQuery {
	return NewWordClient(w.config).QueryGroup(w)
}

// QueryDefinitions queries the "definitions" edge of the Word entity.
func (w *Word) QueryDefinitions() *DefinitionQuery {
	return NewWordClient(w.config).QueryDefinitions(w)
}

// QueryDescendants queries the "descendants" edge of the Word entity.
func (w *Word) QueryDescendants() *WordQuery {
	return NewWordClient(w.config).QueryDescendants(w)
}

// QueryParents queries the "parents" edge of the Word entity.
func (w *Word) QueryParents() *WordQuery {
	return NewWordClient(w.config).QueryParents(w)
}

// Update returns a builder for updating this Word.
// Note that you need to call Word.Unwrap() before calling this method if this Word
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Word) Update() *WordUpdateOne {
	return NewWordClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Word entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Word) Unwrap() *Word {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Word is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Word) String() string {
	var builder strings.Builder
	builder.WriteString("Word(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("create_time=")
	builder.WriteString(w.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(w.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(w.Description)
	builder.WriteString(", ")
	builder.WriteString("descendantCount=")
	builder.WriteString(fmt.Sprintf("%v", w.DescendantCount))
	builder.WriteByte(')')
	return builder.String()
}

// NamedDefinitions returns the Definitions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (w *Word) NamedDefinitions(name string) ([]*Definition, error) {
	if w.Edges.namedDefinitions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := w.Edges.namedDefinitions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (w *Word) appendNamedDefinitions(name string, edges ...*Definition) {
	if w.Edges.namedDefinitions == nil {
		w.Edges.namedDefinitions = make(map[string][]*Definition)
	}
	if len(edges) == 0 {
		w.Edges.namedDefinitions[name] = []*Definition{}
	} else {
		w.Edges.namedDefinitions[name] = append(w.Edges.namedDefinitions[name], edges...)
	}
}

// NamedDescendants returns the Descendants named value or an error if the edge was not
// loaded in eager-loading with this name.
func (w *Word) NamedDescendants(name string) ([]*Word, error) {
	if w.Edges.namedDescendants == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := w.Edges.namedDescendants[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (w *Word) appendNamedDescendants(name string, edges ...*Word) {
	if w.Edges.namedDescendants == nil {
		w.Edges.namedDescendants = make(map[string][]*Word)
	}
	if len(edges) == 0 {
		w.Edges.namedDescendants[name] = []*Word{}
	} else {
		w.Edges.namedDescendants[name] = append(w.Edges.namedDescendants[name], edges...)
	}
}

// NamedParents returns the Parents named value or an error if the edge was not
// loaded in eager-loading with this name.
func (w *Word) NamedParents(name string) ([]*Word, error) {
	if w.Edges.namedParents == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := w.Edges.namedParents[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (w *Word) appendNamedParents(name string, edges ...*Word) {
	if w.Edges.namedParents == nil {
		w.Edges.namedParents = make(map[string][]*Word)
	}
	if len(edges) == 0 {
		w.Edges.namedParents[name] = []*Word{}
	} else {
		w.Edges.namedParents[name] = append(w.Edges.namedParents[name], edges...)
	}
}

// Words is a parsable slice of Word.
type Words []*Word
