package main

import (
	"gotinyns/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Connect()
	defer db.Close()
}
