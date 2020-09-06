package aes

import (
	"fmt"
	"omega/internal/consts"

	goaes "github.com/syronz/goAES"
)

// Decrypt is used when according to consts we want to encrypt an string with AES
func Decrypt(str string) (result string, err error) {
	var encryptor goaes.BuildModel
	encryptor, err = goaes.New().Key(consts.SecretKeyAES).IV(consts.IVAES).Length(len(str)).Build()
	if err != nil {
		return
	}

	result = encryptor.Decrypt(str)
	return
}

// DecryptTwice is used for licenses and each time it produce different output
func DecryptTwice(str string) (result string, err error) {

	var preDecrypted string
	if preDecrypted, err = Decrypt(str); err != nil {
		return
	}

	if len(preDecrypted) < 9 {
		err = fmt.Errorf("string is not valid")
		return
	}

	randStr := preDecrypted[:4] + preDecrypted[len(preDecrypted)-4:]

	var customDecryptor goaes.BuildModel
	customDecryptor, err = goaes.New().Key(randStr).IV(consts.IVAES).Length(len(str)).Build()
	if err != nil {
		return
	}

	result = customDecryptor.Decrypt(preDecrypted[4 : len(preDecrypted)-4])
	return
}
