package swapi

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	var client *Client

	t.Run("NewClient", func(t *testing.T) {
		client = NewClient()
		if client == nil {
			t.Error("newclient is nil")
		}
	})

	t.Run("NewRequest", func(t *testing.T) {
		ctx := context.Background()
		r, err := client.NewRequest(ctx, "/films")
		if err != nil {
			t.Fatalf("error on NewRequest %+v\n", err)
		}

		if r.URL.String() != "https://swapi.co/api/films" {
			t.Fatalf("not match")
		}
	})
}
