package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(database.dev)/game_currency?parseTime=true")
	checkErr(err)

	fmt.Println("Database connected")
}

// checkErr for represent handling error to database connection
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
