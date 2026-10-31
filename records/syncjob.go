package records

import (
	"fmt"
	"time"
)

func SyncJob() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("sync failed:", err)
				}
			}()

			time.Sleep(5 * time.Minute)
			Sync()
		}()
	}
}
