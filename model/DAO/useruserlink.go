package DAO

import (
	db "github.com/model"
	"github.com/util"
	"log"
)

type UserUserLink struct {
	Id      int
	UserId1 int
	UserId2 int
}

type UserUserLinkDAO struct {
}

func NewUserUserLinkDAO() *UserUserLinkDAO {
	return &UserUserLinkDAO{}
}

func (u *UserUserLinkDAO) Save(uul UserUserLink) int {
	time := util.GetCurrentTime("2006-01-02")
	stmt, err := db.Instance().Prepare("INSERT INTO useruserlink(user1id, user2id, created) VALUES($1,$2, $3) RETURNING id")
	db.CheckErr(err)
	res, err := stmt.Exec(uul.UserId1, uul.UserId2, time)
	db.CheckErr(err)
	id, err := res.RowsAffected()
	db.CheckErr(err)
	log.Println(id)
	return int(id)
}

func (u *UserUserLinkDAO) Delete(uul UserUserLink) {

}

func (u *UserUserLinkDAO) GetAll(userid int) ([]int, error) {
	stmt, err := db.Instance().Prepare("select user1id, user2id from useruserlink where user1id = $1 or user2id = $1")
	db.CheckErr(err)

	rows, err := stmt.Query(userid)
	var friendsIds []int
	for rows.Next() {
		var user1id int
		var user2id int
		err = rows.Scan(&user1id, &user2id)
		db.CheckErr(err)
		if user1id != userid && !util.Contains(friendsIds, user1id) {
			friendsIds = append(friendsIds, user1id)
		} else if user2id != userid && !util.Contains(friendsIds, user2id) {
			friendsIds = append(friendsIds, user2id)
		}
	}

	return friendsIds, nil
}
