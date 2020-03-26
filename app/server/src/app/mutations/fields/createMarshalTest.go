package mutations

import (
	"app/model"
	types "app/types"
	"github.com/graphql-go/graphql"
)

var CreateMarshalTest = &graphql.Field{
	Type:        types.MarshalTest,
	Description: "Create a new Marshal Test",
	Args: graphql.FieldConfigArgument{
		"stringValue": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return model.MarshalTest{
			StringValue: types.ToJsonNullString(params.Args["stringValue"].(string)),
		}, nil
	},
}
