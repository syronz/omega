package filter

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	samples := []struct {
		inStr  string
		inCols []string
		out    string
		err    error
	}{
		{
			inStr:  "username[eq]'diako'",
			inCols: []string{"username"},
			out:    "username = 'diako'",
		},
		{
			inStr:  "created_by[eq]'makwan'[and]name[eq]'ako'",
			inCols: []string{"created_by", "users.name"},
			out:    "created_by = 'makwan' AND name = 'ako'",
		},
		{
			inStr:  "users.name[eq]'diako'[and](age[gte]25[or]role[eq]'admin')",
			inCols: []string{"users.name", "users.age", "users.role"},
			out:    "users.name = 'diako' AND (age >= 25 OR role = 'admin')",
		},
		{
			inStr:  "users.name[eq]'diako'[and](age[gte]25[or]role[eq]'admin')",
			inCols: []string{"users.age", "users.role"},
			out:    "",
			err:    fmt.Errorf("column 'users.name' not exist"),
		},
		{
			inStr:  "'; select * from bas_users",
			inCols: []string{"users.age", "users.role"},
			out:    "",
			err:    fmt.Errorf("column 'users.name' not exist"),
		},
		{
			inStr:  "bas_roles.name[eq]'admin'",
			inCols: []string{"name"},
			out:    "",
		},
	}

	for _, v := range samples {
		result, err := Parser(v.inStr, v.inCols)
		if result != v.out {
			t.Errorf("\nin: %q, %q\nout: %q, err:%q\nshould be: %q",
				v.inStr, v.inCols, result, err, v.out)
		}

	}

}
