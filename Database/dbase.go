package Database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "raja.db.elephantsql.com"
	port     = 5432
	user     = "stzrurfj"
	password = "mHLqxoPKfj2P0R5XD2AImPSr8Ozu7rWr"
	dbname   = "stzrurfj"
)

var DB *sql.DB

func DBInit() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = db

}
