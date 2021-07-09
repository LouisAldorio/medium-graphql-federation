package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *entityResolver) FindPostByID(ctx context.Context, id int) (*model.Post, error) {
	return &model.Post{
		ID: id,
	}, nil
}

func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*model.User, error) {
	return service.UserFindByID(id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
