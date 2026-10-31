package main

import (
	"gotinyns/db"
	"gotinyns/records"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Connect()
	defer db.Close()

	records.Sync()
	go records.SyncJob()

	select {}
}
