// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Definitions struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	User *User  `json:"user"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewWord struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Word struct {
	ID          string         `json:"id"`
	Text        string         `json:"text"`
	User        *User          `json:"user"`
	Definitions []*Definitions `json:"definitions,omitempty"`
}
