package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Id int64 `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func NewDbMap(dsn string) *gorp.DbMap {
	db, err := sql.Open("mysql", dsn)
	
	if err != nil {
		log.Fatalln("sql.Open failed")
	}

	return &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
}

func CreateUserTable(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	err := dbmap.CreateTables()

	if err != nil {
		panic(err)
	}
}
