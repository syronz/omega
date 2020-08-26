package corerr

import "fmt"

type notFound struct {
	Part  string
	Field string
	Value string
}

func (p notFound) Error() string {
	return fmt.Sprintf("record not found in %v for %v with value %v", p.Part, p.Field, p.Value)
}

// NewNotFound is used for returning the notFound error
func NewNotFound(part, field, value string) error {
	return notFound{
		part,
		field,
		value,
	}
}
