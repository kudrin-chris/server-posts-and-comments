package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"server/graph/model"
)

// CreateComment is the resolver for the CreateComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.InputComment) (*model.Comment, error) {
	newComment, err := r.CommentsService.CreateComment(input.FromInput())
	if err != nil {
		panic(err)
	}
	return &newComment, nil
}

// CommentsSubscription is the resolver for the CommentsSubscription field.
func (r *subscriptionResolver) CommentsSubscription(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	panic(fmt.Errorf("not implemented: CommentsSubscription - CommentsSubscription"))
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
