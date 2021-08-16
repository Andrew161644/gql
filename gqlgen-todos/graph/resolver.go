package graph

import (
	"andrew.com/bff/BackRoundResolver"
	"andrew.com/bff/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var userAddedChan map[int]chan *model.User
var BResolver *BackRoundResolver.Backround_resolver

func init() {

	userAddedChan = map[int]chan *model.User{}
}

type Resolver struct {
	todos []*model.Todo
}
