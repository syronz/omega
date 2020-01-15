package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash is used for hashing the password
func Hash(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(bytes), err
}
