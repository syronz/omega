package password

import (
	"testing"
)

func TestHashVerify(t *testing.T) {
	samples := []struct {
		password string
		salt     string
		hash     string
	}{
		{"123", "a", ""},
		{"user", "q7Gcqm9VXMVpf33PbFlYEpkMmDqOn1gRMVsavha9lQ8", ""},
		{"", "", ""},
	}

	for i, v := range samples {
		samples[i].hash, _ = Hash(v.password, v.salt)
	}

	for _, v := range samples {
		if !Verify(v.password, v.hash, v.salt) {
			t.Fatalf("for pass=%q and salt=%q is not working", v.password, v.salt)
		}
	}

}

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
		_, _ = result, err
	}
}
