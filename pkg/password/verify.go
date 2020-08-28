package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Verify is used for verifying the hash and password
func Verify(password, hash, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}
