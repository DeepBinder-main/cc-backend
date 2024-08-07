// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type IntRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type Query struct {
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Aggregate string

const (
	AggregateUser    Aggregate = "USER"
	AggregateProject Aggregate = "PROJECT"
	AggregateCluster Aggregate = "CLUSTER"
)

var AllAggregate = []Aggregate{
	AggregateUser,
	AggregateProject,
	AggregateCluster,
}

func (e Aggregate) IsValid() bool {
	switch e {
	case AggregateUser, AggregateProject, AggregateCluster:
		return true
	}
	return false
}

func (e Aggregate) String() string {
	return string(e)
}

func (e *Aggregate) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Aggregate(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Aggregate", str)
	}
	return nil
}

func (e Aggregate) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
