package controller

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	Name     string
	User     string
	Password string
	IP       string
}

var database Database

func SetupDatabase(db Database) {
	database = db
	fmt.Printf("%#v", database)

}

func NewDB() *sqlx.DB {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", database.User, database.Password, database.IP, database.Name)

	//dd, errr := sqlx.Connect("mysql", connectString)
	db, err := sqlx.Connect("mysql", connectString)
	if err != nil {
		log.Println(err)
	}

	return db
}
