package swapi

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Client struct {
	base string
	http *http.Client
}

func NewClient() *Client {
	return &Client{
		base: "https://swapi.co/api",
		http: newHTTPClient(),
	}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
}

func (c *Client) NewRequest(ctx context.Context, url string) (*http.Request, error) {
	url = c.base + url
	r, err := http.NewRequest("GET", url, nil)
	log.Println(url)
	if err != nil {
		return nil, err
	}

	return r.WithContext(ctx), nil
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.http.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}
	}

	return resp, nil
}
