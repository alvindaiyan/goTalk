package DAO

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	// model "github.com/model/DAO"
)

const (
	TOKEN_LENGTH = 32
)

func PerformLogin(uname string, pwd string) (string, bool) {
	// this method is not finished
	userDao := NewUserDAO()
	user := userDao.GetUserByName(uname)
	if pwd == user.Pwd {
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

func PerformLogout(token string) {
	// todo this is a sub
}

func TokenGenerator() (string, error) {
	// change the length of the generated random string here
	size := TOKEN_LENGTH

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
