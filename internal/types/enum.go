package types

import "strings"

// Enum is used for define all types
type Enum string

// JoinEnum convert slice of enums to a string seperated by comma
func JoinEnum(arr []Enum) string {
	var strArr []string

	for _, v := range arr {
		strArr = append(strArr, string(v))
	}

	return strings.Join(strArr, ", ")
}
