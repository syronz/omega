package filter

import "testing"

func TestParser(t *testing.T) {
	samples := []struct {
		inStr  string
		inCols string
		out    string
		err    error
	}{
		{
			inStr:  "username[eq]'diako'",
			inCols: "username",
			out:    "username = 'diako'",
		},
		{
			inStr:  "created_by[eq]'makwan'[and]name[eq]'ako'",
			inCols: "created_by, name",
			out:    "created_by = 'makwan' AND name = 'ako'",
		},
		{
			inStr:  "users.name[eq]'diako'[and](age[gte]25[or]role[eq]'admin')",
			inCols: "users.name, users.age, users.role",
			out:    "users.name = 'diako' AND (age >= 25 OR role = 'admin')",
		},
		{
			inStr:  "users.name[eq]'diako'[and](age[gte]25[or]role[eq]'admin')",
			inCols: "users.age, users.role",
			out:    "users.name = 'diako' AND (age >= 25 OR role = 'admin')",
		},
	}

	for i, v := range samples {
		result, err := Parser(v.inStr, v.inCols)
		if result != v.out {
			t.Errorf("\nin: %q, %q\nout: %q, err:%q\nshould be: %q",
				v.inStr, v.inCols, result, err, v.out)
		}

		t.Log(i, v)
	}

}
