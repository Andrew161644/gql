package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"andrew.com/bff/BackRoundResolver"
	"context"
	"fmt"
	"log"
	"math/rand"

	"andrew.com/bff/graph/generated"
	"andrew.com/bff/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	BResolver.MyEvent <- BackRoundResolver.EventSimple{
		Id:  input.UserID,
		Msg: input.Text,
	}
	for _, observer := range userAddedChan {
		observer <- &model.User{
			ID:   input.UserID,
			Name: input.Text,
		}
	}
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *subscriptionResolver) UserAdded(ctx context.Context) (<-chan *model.User, error) {
	userAddedEvent := make(chan *model.User, 1)
	go func() {
		<-ctx.Done()
	}()
	id := rand.Int()
	userAddedChan[id] = userAddedEvent
	log.Println(userAddedChan)
	return userAddedEvent, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
