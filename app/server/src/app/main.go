package main

import (
	"app/queries"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queries.RootQuery,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err.Error)
	}
}
