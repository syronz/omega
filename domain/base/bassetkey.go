package base

import (
	"omega/internal/types"
	"strings"
)

// settings key for base domain
const (
	DefaultRegisteredRole types.Setting = "default_registered_role"
)

var List = []types.Setting{
	DefaultRegisteredRole,
}

// Join make a string for showing in the api
func Join() string {
	var strArr []string

	for _, v := range List {
		strArr = append(strArr, string(v))
	}

	return strings.Join(strArr, ", ")
}
