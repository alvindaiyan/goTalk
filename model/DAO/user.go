package DAO

import (
	db "github.com/model"
	"github.com/util"
	"log"
)

type User struct {
	Id    int
	Name  string
	Pwd   string
	Token string
}

type UserDAO struct {
}

const (
	TABLE_NAME = "userinfo"
)

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// save user
func (u *UserDAO) Save(user User) int {
	time := util.GetCurrentTime("2006-01-02")
	stmt, err := db.Instance().Prepare("INSERT INTO userinfo(username,password,created) VALUES($1,$2,$3) RETURNING uid")
	db.CheckErr(err)
	hashpwd, err := encrypt(user.Pwd, KEY)
	if err != nil {
		log.Println("error encrypt the password")
		return -1
	}
	res, err := stmt.Exec(user.Name, hashpwd, time)
	db.CheckErr(err)
	id, err := res.LastInsertId()
	db.CheckErr(err)
	log.Println(id)
	return int(id)

}

// delete user
func (u *UserDAO) Delete(id int) {
	stmt, err := db.Instance().Prepare("delete from userinfo where uid=$1")
	db.CheckErr(err)

	res, err := stmt.Exec(id)
	affect, err := res.RowsAffected()
	db.CheckErr(err)

	log.Println(affect)
}

// get user by id
func (u *UserDAO) Get(id int) User {
	stmt, err := db.Instance().Prepare("select uid, username, password from userinfo where uid=$1")
	db.CheckErr(err)

	rows, err := stmt.Query(id)

	var usr User
	for rows.Next() {
		var uid int
		var username string
		var password string
		err = rows.Scan(&uid, &username, &password)
		db.CheckErr(err)
		usr.Id = uid
		usr.Name = username
		usr.Pwd = password
	}

	return usr
}

func GetUserIdByName(uname string) (int, error) {
	// defer func() {
	// 	if x := recover(); x != nil {
	// 		return -1, errors.New("cannot find the user")
	// 	}
	// }()

	stmt, err := db.Instance().Prepare("select uid from userinfo where username=$1")
	db.CheckErr(err)

	rows, err := stmt.Query(uname)

	var uid int
	for rows.Next() {
		err = rows.Scan(&uid)
		db.CheckErr(err)
	}

	return uid, nil
}

// get user by id
func (u *UserDAO) GetUserByName(uname string) User {
	stmt, err := db.Instance().Prepare("select uid, username, password from userinfo where username=$1")
	db.CheckErr(err)

	rows, err := stmt.Query(uname)

	var usr User
	for rows.Next() {
		var uid int
		var username string
		var password string
		err = rows.Scan(&uid, &username, &password)
		db.CheckErr(err)
		usr.Id = uid
		usr.Name = username
		usr.Pwd = password
	}

	return usr
}
