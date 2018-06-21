package resolver

import (
	"context"

	"github.com/de3/gg/swapi"
)

type QueryResolver struct {
	client *swapi.Client
}

func New(client *swapi.Client) *QueryResolver {
	return &QueryResolver{
		client: client,
	}
}

type FilmsQuery struct {
	Title *string
}

func (r *QueryResolver) Films(ctx context.Context, args FilmsQuery) (*[]*filmResolver, error) {
	resp, err := r.client.SearchFilms(ctx, strValue(args.Title))
	if err != nil {
		return nil, err
	}

	return ListFilm(ctx, resp)
}

func strValue(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}
