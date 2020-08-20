package aes

import "testing"

func TestEncrypt(t *testing.T) {
	r, err := Encrypt("123456")
	if err != nil || r != "70133af07bfa" {
		t.Errorf("result should be 70133af07bfa but it is %q", r)
	}
}

func TestDecrypt(t *testing.T) {
	r, err := Decrypt("70133af07bfa")
	if err != nil || r != "123456" {
		t.Errorf("result should be 123456 but it is %q", r)
	}
}

func TestEncryptTwice(t *testing.T) {
	r, err := EncryptTwice("3")
	if err != nil || r != "70133af07bfa" {
		t.Errorf("result should be 70133af07bfa but it is %q", r)
	}
}

//result should be 70133af07bfa but it is "781939f27bfdc28b0dc8f50e5723d0457366dbd64ddd"
func TestDecryptTwice(t *testing.T) {
	// r, err := DecryptTwice("781939f27bfdc28b0dc8f50e5723d0457366dbd64ddd")
	// r, err := DecryptTwice("79143bf27afac08477e3f5")
	r, err := DecryptTwice("35753db37bfd84e44cf2")
	//
	if err != nil || r != "123456" {
		t.Errorf("result should be 70133af07bfa but it is %q", r)
	}
}
