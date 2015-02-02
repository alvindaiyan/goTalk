package config

import (
	"errors"
	model "github.com/model"
	"strconv"
)

type AppConfig struct {
	msgcs map[int]chan model.Message
}

const (
	MAX_CHAN     = 10000
	TOKEN_LENGTH = 32
)

func (app *AppConfig) Init() {
	msgcs := make(map[int]chan model.Message)
	app.msgcs = msgcs
}

func (app AppConfig) GetMsgcs() map[int]chan model.Message {
	return app.msgcs
}

// maybe use panic-recover process
// when error happens, should be panic, and recover to delete the existing chan and add new one?
func (app *AppConfig) SetMsgcs(uid int, msgc chan model.Message) error {
	if app.msgcs[uid] != nil {
		return errors.New("cannot create channel for user:" + strconv.Itoa(uid) + "  since it already exists")
	}
	app.msgcs[uid] = msgc
	return nil
}
