package db

import(
	"time"
	"fmt"
)

//need delete files metadata records after expiry time
//reference from go by example tickers are used to run something at regular intervals

func CleanupScheduler() {
    ticker := time.NewTicker(24 * time.Hour) 
    go func() {
        for range ticker.C {
            db := ConnectDB()
			defer db.Close()
            
            _, err := db.Exec(`
                DELETE FROM userfiles 
                WHERE expiry_at < NOW()
            `)
            
            if err != nil {
                fmt.Println("Failed to delete expired files")
            } else {
                fmt.Println("Successfully cleaned up expired files")
            }
            
            
        }
    }()
}


