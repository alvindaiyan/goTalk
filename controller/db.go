package controller

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

var once sync.Once

// user = gotalk
// password = gotalk123
type DBCon struct {
	instance *sql.DB
}

type DBHandler struct {
	ins      DBCon
	instance func(ins *DBCon, usr string, pwd string) *sql.DB
}

func (d DBHandler) Instance(usr string, pwd string) *sql.DB {
	return d.instance(&d.ins, usr, pwd)
}

// connect for pgql
func PostgresIns(con *DBCon, usr string, pwd string) *sql.DB {
	if con.instance == nil {
		once.Do(func() {
			db, err := sql.Open("postgres", "user="+usr+" password="+pwd+" dbname=goTalk-db sslmode=disable")
			checkErr(err)
			con.instance = db
		})
	}
	return con.instance
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
