package main

import (
	"fileshare/db"
	 "fmt"
	 "github.com/joho/godotenv"
	 "log"
	 "net/http"
	 "fileshare/routes"
)

func loadenv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//here db is initialized and server is started
func main() {
	loadenv()
	fmt.Println("Initializing database..")
	db.Initdb()
	fmt.Println("Database initialized..")

	r := routes.Router()
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}