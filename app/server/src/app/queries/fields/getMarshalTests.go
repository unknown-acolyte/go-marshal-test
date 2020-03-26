package queries

import (
	"app/data"
	"app/model"
	types "app/types"
	"encoding/json"
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

const query = "select string_value from marshal_test"

var GetMarshalTests = &graphql.Field{
	Type:        graphql.NewList(types.MarshalTest),
	Description: "Get list of all Marshal Tests",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		rows, err := mysql.DB.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		var marshalTestList []model.MarshalTest
		for rows.Next() {
			// create object
			var marshalTest model.MarshalTest

			// scan row values into object
			err = rows.Scan(&marshalTest.StringValue)
			if err != nil {
				panic(err.Error())
			}

			// logging
			v, err := json.Marshal(marshalTest)
			if err != nil {
				log.Error(err.Error)
			}
			log.Trace(string(v))

			// add object to list
			marshalTestList = append(marshalTestList, marshalTest)
		}

		return marshalTestList, nil
	},
}
