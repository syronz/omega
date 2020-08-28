package aes

import (
	"omega/internal/consts"
	"omega/pkg/random"

	goaes "github.com/syronz/goAES"
)

// Encrypt is used when according to consts we want to encrypt an string with AES
func Encrypt(str string) (result string, err error) {
	var encryptor goaes.BuildModel
	encryptor, err = goaes.New().Key(consts.SecretKeyAES).IV(consts.IVAES).Build()
	if err != nil {
		return
	}

	result = encryptor.Encrypt(str)
	return
}

// EncryptTwice is used for licenses and each time it produce different output
func EncryptTwice(str string) (result string, err error) {
	randomStr := random.String(8)
	var customEncryptor goaes.BuildModel
	customEncryptor, err = goaes.New().Key(randomStr).IV(consts.IVAES).Build()
	if err != nil {
		return
	}

	preEncrypted := customEncryptor.Encrypt(str)

	result, err = Encrypt(randomStr[:4] + preEncrypted + randomStr[4:])
	return
}
