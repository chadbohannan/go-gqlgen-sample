//go:generate go run github.com/99designs/gqlgen

package go_gqlgen_sample

import (
	"context"
	"errors"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) NewArticle(ctx context.Context, input NewArticle) (*ArticleResponse, error) {
	if input.UserID < 0 || input.UserID >= len(users) {
		return nil, errors.New("invalid userID")
	}
	article := &Article{
		ID: len(articles),
		Text: input.Text,
		UserID: input.UserID,
	}
	articles = append(articles, article)

	response := &ArticleResponse{
		ID:   article.ID,
		Text: article.Text,
		User: *users[article.UserID],
	}
	return response, nil
}

func (r *mutationResolver) SetArticleText(ctx context.Context, input EditArticle) (*ArticleResponse, error) {
	if input.ID < 0 || input.ID >= len(articles) {
		return nil, errors.New("index out of bounds")
	}
	article := articles[input.ID]
	article.Text = input.Text
	response := &ArticleResponse{
		ID:   article.ID,
		Text: article.Text,
		User: *users[article.UserID],
	}
	return response, nil
}

func (r *mutationResolver) NewUser(ctx context.Context, input NewUser) (*User, error) {
	user := &User{
		ID: len(users),
		Name: input.Name,
	}
	users = append(users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Articles(ctx context.Context) ([]ArticleResponse, error) {
	var response []ArticleResponse
	for _, article := range articles {
		r := ArticleResponse{
			ID:   article.ID,
			Text: article.Text,
			User: *users[article.UserID],
		}
		response = append(response, r)
	}
	return response, nil
}

func (r *queryResolver) Article(ctx context.Context, id *int) (*ArticleResponse, error) {
	if *id < 0 || *id >= len(articles) {
		return nil, errors.New("index out of bounds")
	}
	article := articles[*id]
	response := &ArticleResponse{
		ID:   article.ID,
		Text: article.Text,
		User: *users[article.UserID],
	}
	return response, nil
}

// Users returns a list of literals, not pointers, so we need to make a deep copy
// the return signature is dictated by generated code; it's best not to fight it
func (r *queryResolver) Users(ctx context.Context) ([]User, error) {
	var response []User
	for _, user := range users {
		response = append(response, *user)
	}
	return response, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*User, error) {
	if *id < 0 || *id>= len(users) {
		return nil, errors.New("index out of bounds")
	}
	return users[*id], nil
}
