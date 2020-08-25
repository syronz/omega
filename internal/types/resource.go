package types

import (
	"fmt"
	"strings"
)

// Resource is a special type for checking the access
type Resource string

func (p *Resource) String() string {
	return fmt.Sprint(*p)
}

// ResourceJoin make a string from an array of resources
func ResourceJoin(resources []Resource) string {
	var strArr []string

	for _, v := range resources {
		strArr = append(strArr, string(v))
	}

	return strings.Join(strArr, ", ")
}
