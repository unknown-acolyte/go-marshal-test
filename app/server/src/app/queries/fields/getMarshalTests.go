package queries

import (
	"app/model"
	types "app/types"
	"github.com/graphql-go/graphql"
)

var GetMarshalTests = &graphql.Field{
	Type:        graphql.NewList(types.MarshalTest),
	Description: "Get list of all Marshal Tests",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var marshalTestList []model.MarshalTest

		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: types.NullableString{String: "", Valid: true}})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: types.NullableString{String: "test1", Valid: true}})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: types.NullableString{String: "", Valid: false}})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: types.NullableString{String: "test2", Valid: false}})

		return marshalTestList, nil
	},
}
