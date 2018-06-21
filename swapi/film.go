package swapi

import (
	"context"
	"net/url"
)

type Film struct {
	Title     string `json:"title"`
	EpisodeID int64  `json:"episode_id"`
	Director  string `json:"director"`
	URL       string `json:"url"`
}

type FilmResponse struct {
	Count int64  `json:"count"`
	Films []Film `json:"results"`
}

func (c *Client) SearchFilms(ctx context.Context, title string) (FilmResponse, error) {
	q := url.Values{"search": {title}}

	req, err := c.NewRequest(ctx, "/films?"+q.Encode())
	if err != nil {
		return FilmResponse{}, err
	}

	var resp FilmResponse
	if _, err := c.Do(req, &resp); err != nil {
		return FilmResponse{}, err
	}

	return resp, nil
}
