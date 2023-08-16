package util

import (
	"fmt"
	"log"

	"github.com/mcgtrt/azure-graphql/graph/model"
)

func MakeQueryFromUpdateParams(params *model.UpdateEmployeeParams) string {
	pMap := make(map[string]string)
	if params.FirstName != nil {
		pMap["FirstName"] = *params.FirstName
	}
	if params.LastName != nil {
		pMap["LastName"] = *params.LastName
	}
	if params.Username != nil {
		pMap["Username"] = *params.Username
	}
	if params.Dob != nil {
		pMap["DOB"] = *params.Dob
	}
	if params.DepartmentID != nil {
		pMap["DOB"] = *params.Dob
	}
	if params.Position != nil {
		pMap["Position"] = *params.Position
	}

	var (
		index = 1
		mLen  = len(pMap)
		query string
	)

	for k, v := range pMap {
		query += fmt.Sprintf("%s = '%s'", k, v)

		if mLen != index {
			query += ", "
		}
		index++
	}

	log.Printf("QUERY: %s\n", query)
	return query
}
