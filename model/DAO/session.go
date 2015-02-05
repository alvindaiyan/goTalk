package DAO

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// encrypt method is aes256
const (
	TOKEN_LENGTH = 32
)

var (
	KEY = []byte("XY0nG86PSRJqMGz957Yza1D34393MPII") // 32 bytes
)

func PerformLogin(uname string, pwd string) (string, bool) {
	// this method is not finished
	userDao := NewUserDAO()
	user := userDao.GetUserByName(uname)
	hashpwd, err := decrypt(user.Pwd, KEY)
	if err != nil {
		return "", false
	}
	if pwd == hashpwd {
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

func encrypt(str string, key []byte) (string, error) {
	text := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return string(ciphertext[:]), nil
}

func decrypt(str string, key []byte) (string, error) {
	text := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data[:]), nil
}
