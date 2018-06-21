package resolver

import (
	"context"

	"github.com/de3/swapi-go-graphql/swapi"
	graphql "github.com/graph-gophers/graphql-go"
)

type filmResolver struct {
	film swapi.Film
}

func SingleFilm(ctx context.Context, film swapi.Film) (*filmResolver, error) {
	return &filmResolver{film: film}, nil
}

func ListFilm(ctx context.Context, films swapi.FilmResponse) (*[]*filmResolver, error) {
	var result []*filmResolver

	for _, film := range films.Films {
		resolver, _ := SingleFilm(ctx, film)

		result = append(result, resolver)
	}

	return &result, nil
}

func (r *filmResolver) ID() graphql.ID {
	return graphql.ID(r.film.URL)
}

func (r *filmResolver) Episode() int32 {
	return int32(r.film.EpisodeID)
}

func (r *filmResolver) Title() string {
	return r.film.Title
}

func (r *filmResolver) DirectorName() string {
	return r.film.Director
}
