package controller

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	once  sync.Once
	dbCon *sql.DB
)

type ConstrainValue struct {
	value    string
	dataType DataType
}

type DataType int

const (
	INTEGER DataType = iota
	STRING
)

func Instance() *sql.DB {
	return postgresIns()
}

// connect for pgql
func postgresIns() *sql.DB {
	if dbCon == nil {
		once.Do(func() {
			// todo: read from a config file
			////////////////////////////////
			uname := "postgres"
			pwd := "Password1"
			dbname := "godb"
			////////////////////////////////
			log.Println("get a data connection")
			db, err := sql.Open("postgres", "user="+uname+" password="+pwd+" dbname="+dbname+" sslmode=disable")
			CheckErr(err)
			dbCon = db
		})
	}
	return dbCon
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
