package DAO

import (
// db "github.com/model"
// "github.com/util"
// "log"
)

type UserUserLink struct {
	Id      int
	UserId1 string
	UserId2 string
}

type UserUserLinkDAO struct {
}

func NewUserUserLinkDAO() *UserUserLink {
	return &UserUserLink{}
}

func (u *UserUserLinkDAO) Save(uul UserUserLink) int {
	return 0
}

func (u *UserUserLinkDAO) Delete(id int) {

}

func (u *UserUserLinkDAO) Get(userid int) ([]int, error) {
	return nil, nil
}
