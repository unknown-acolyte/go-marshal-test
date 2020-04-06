package queries

import (
	"app/model"
	types "app/types"
	"github.com/graphql-go/graphql"
	"github.com/guregu/null"
)

var GetMarshalTests = &graphql.Field{
	Type:        graphql.NewList(types.MarshalTest),
	Description: "Get list of all Marshal Tests",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var marshalTestList []model.MarshalTest

		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: null.NewString("", true)})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: null.NewString("test1", true)})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: null.NewString("", false)})
		marshalTestList = append(marshalTestList, model.MarshalTest{StringValue: null.NewString("test2", false)})

		return marshalTestList, nil
	},
}
