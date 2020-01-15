package loghook

import (
	"testing"
)

func TestFindCaller(t *testing.T) {

	results := [5]string{
		"loghook/caller.go",
		"loghook/caller.go",
		"loghook/caller_test.go",
		"testing/testing.go",
		"runtime/asm_amd64.s",
	}

	for i, v := range results {
		file, functions, line := findCaller(i)
		if file != v {
			t.Errorf("The response for findCaller(%d) should be %q but returns %q",
				i, v, file)
		}
		t.Log(file, functions, line, v)
	}

}

func TestGetCaller(t *testing.T) {

	results := [4]string{
		"loghook/caller.go",
		"loghook/caller_test.go",
		"testing/testing.go",
		"runtime/asm_amd64.s",
	}

	for i, v := range results {
		p, f, l := getCaller(i)
		if f != v {
			t.Errorf("The response for getCaller(%d) should be %q but returns %q",
				i, v, p)
		}
		t.Log(p, f, l)
	}

}
