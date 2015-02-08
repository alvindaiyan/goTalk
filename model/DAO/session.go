package DAO

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// encrypt method is aes256
const (
	TOKEN_LENGTH = 32
)

var (
	KEY            = []byte("XY0nG86PSRJqMGz957Yza1D34393MPII") // 32 bytes
	provides       = make(map[string]Provider)
	globalSessions *Manager
)

type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxlifetime int64
}

func NewManager(providerName, cookieName string, maklifetime int64) (*Manager, error) {
	return nil, nil
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider" + name)
	}
	provides[name] = provider
}

func (manager *Manager) sessionId() (string, error) {
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

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value != "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryEscape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}

	return
}

func login(w http.ResponseWriter, r *http.Request) {
	session := globalSessions.SessionStart(w, r)
	r.ParseForm()
	session.Set("username", r.Form["username"])
	http.Redirect(w, r, "/", 302)
}

func PerformLogin(uname string, pwd string) (string, bool) {
	// this method is not finished
	userDao := NewUserDAO()
	manager, _ := NewManager("providerName", "cookieName", 0)
	user := userDao.GetUserByName(uname)
	hashpwd, err := decrypt(user.Pwd, KEY)
	if err != nil {
		return "", false
	}
	if pwd == hashpwd {
		token, err := manager.sessionId()
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

func init() {
	globalSessions, _ = NewManager("memory", "gosessionid", 3600)
}
