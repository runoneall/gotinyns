package records

import (
	"fmt"
	"gotinyns/db"
	"sync"
)

var (
	records []Record
	mu      sync.RWMutex
)

func Sync() {
	dbconn := db.Conn()
	rows, err := dbconn.Query("SELECT name, type, value, ttl FROM records")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var records_new []Record
	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.Name, &r.Type, &r.Value, &r.TTL); err != nil {
			panic(err)
		}

		records_new = append(records_new, r)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	mu.Lock()
	records = records_new
	mu.Unlock()

	fmt.Println("synced", len(records), "record(s)")
}
