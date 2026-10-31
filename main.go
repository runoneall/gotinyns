package main

import (
	"gotinyns/db"
	"gotinyns/records"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	sigInterrupt := make(chan os.Signal, 1)
	signal.Notify(sigInterrupt, os.Interrupt)

	db.Connect()
	defer db.Close()

	records.Sync()
	go records.SyncJob()

	<-sigInterrupt
}
