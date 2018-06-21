package main

import (
	"log"
	"net/http"

	"github.com/de3/swapi-go-graphql/resolver"
	"github.com/de3/swapi-go-graphql/schema"
	"github.com/de3/swapi-go-graphql/swapi"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	client := swapi.NewClient()
	queryResolver := resolver.New(client)

	schema := graphql.MustParseSchema(schema.String(), queryResolver)

	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
