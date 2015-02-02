package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/config"
)

func PerformLogin(uname string, pwd string) (string, bool) {
	// this method is not finished
	if pwd == "password" {
		token, err := TokenGenerator()
		if err != nil {
			return token, false
		} else {
			return token, true
		}
	} else {
		return "wrong password", false
	}
}

func TokenGenerator() (string, error) {
	// change the length of the generated random string here
	size := config.TOKEN_LENGTH

	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		return "error", errors.New("cannot generate token for user")
	}
	rs := base64.URLEncoding.EncodeToString(rb)
	return rs, nil
}

func ValidateToken(token string) bool {
	return true
}

func encrypt(str string) string {
	return str
}

func decrypt(str string) string {
	return str
}
