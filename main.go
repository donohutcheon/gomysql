package main

import (
	"database/sql"
	"fmt"
	"gomysql/db"
)

func main() {
	var con *sql.DB
	fmt.Println("phpflow.com - Go MySQL Tutorial")
	con := db.CreateCon()
}
