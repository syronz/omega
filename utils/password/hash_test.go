package password

import (
	"testing"
)

func TestHash(t *testing.T) {
	samples := []struct {
		in  string
		err error
	}{
		{"hi", nil},
		{"123456", nil},
	}

	for _, v := range samples {
		result, err := Hash(v.in, "this is salt")
		t.Log(v, result, err)
	}
}
