package types

import (
	"github.com/graphql-go/graphql"
)

var MarshalTest = graphql.NewObject(graphql.ObjectConfig{
	Name: "MarshalTest",
	Fields: graphql.Fields{
		"stringValue": &graphql.Field{
			Type: graphql.String,
		},
	},
})
