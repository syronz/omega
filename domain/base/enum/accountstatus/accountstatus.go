package accountstatus

import (
	"omega/internal/types"
	"strings"
)

const (
	Active   types.Enum = "active"
	Inactive types.Enum = "inactive"
)

var List = []types.Enum{
	Active,
	Inactive,
}

// Join make a string for showing in the api
func Join() string {
	var strArr []string

	for _, v := range List {
		strArr = append(strArr, string(v))
	}

	return strings.Join(strArr, ", ")
}
