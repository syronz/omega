package filter

import "testing"

func TestParser(t *testing.T) {
	samples := []struct {
		in  string
		out string
	}{
		// {
		// 	in:  "username[eq]'diako'",
		// 	out: "username = 'diako'",
		// },
		{
			in:  "created_by[eq]'makwan'[and]name[eq]'ako'",
			out: "username = 'diako'",
		},
	}

	for i, v := range samples {
		result := Parser(v.in)
		if result != v.out {
			t.Errorf("Parser(%q) = %q, which is wrong, correct result is %q",
				v.in, result, v.out)
		}

		t.Log(i, v)
	}

}
