// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Book struct {
	ID        string       `json:"Id"`
	Name      string       `json:"Name"`
	Type      *BookType    `json:"Type,omitempty"`
	Publisher []*Publisher `json:"Publisher"`
}

type Mutation struct {
}

type NewBook struct {
	Name string    `json:"Name"`
	Type *BookType `json:"Type,omitempty"`
}

type Publisher struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Query struct {
}

type BookType string

const (
	BookTypeHistory BookType = "History"
	BookTypeScience BookType = "Science"
	BookTypeMath    BookType = "Math"
)

var AllBookType = []BookType{
	BookTypeHistory,
	BookTypeScience,
	BookTypeMath,
}

func (e BookType) IsValid() bool {
	switch e {
	case BookTypeHistory, BookTypeScience, BookTypeMath:
		return true
	}
	return false
}

func (e BookType) String() string {
	return string(e)
}

func (e *BookType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BookType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BookType", str)
	}
	return nil
}

func (e BookType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
