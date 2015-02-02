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
	uname    string
	pwd      string
	dbname   string
	instance *sql.DB
}

type DBHandler struct {
	ins      DBCon
	instance func(ins *DBCon) *sql.DB
}

func (d DBHandler) Instance() *sql.DB {
	return d.instance(&d.ins)
}

// connect for pgql
func PostgresIns(con *DBCon) *sql.DB {
	if con.instance == nil {
		once.Do(func() {
			// todo: read from a config file
			////////////////////////////////
			con.uname = "postgres"
			con.pwd = "Password1"
			con.dbname = "godb"
			////////////////////////////////
			db, err := sql.Open("postgres", "user="+con.uname+" password="+con.pwd+" dbname="+con.dbname+" sslmode=disable")
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
