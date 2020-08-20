package search

import (
	"omega/internal/param"
	"testing"
)

// TODO: test should be completed
func TestParser(t *testing.T) {
	params := param.Param{
		PreCondition: "age > 100",
		Search:       "term",
	}

	pattern := `(ii.type like '%[1]v' OR
		m.name like '%[1]v' OR
		c.name like '%%%[1]v')`

	// result := Parse(params, pattern)

	params.Search = "user.phone>4350907"
	_ = Parse(params, pattern)
}
