package main

import (
	"fileshare/db"
	 "fmt"
	 "github.com/joho/godotenv"
	 "log"
)

func loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadenv()
	fmt.Println("Initializing database..")
	db.Initdb()
	fmt.Println("Database initialized..")
}