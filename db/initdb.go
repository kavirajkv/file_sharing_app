package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)



// db connection based on database name
func ConnectDB() *sql.DB {
	db_pass := os.Getenv("PG_PASSWORD")
	url:=os.Getenv("PG_HOST")
	port:=os.Getenv("PG_PORT")
	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=fileshare port=%s sslmode=disable",url,db_pass,port)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("Database connected..")
	return db

}

//initialize db with tables
func Initdb(){
	db := ConnectDB()
	defer db.Close()

	// table for users 
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		userid SERIAL PRIMARY KEY,
		username VARCHAR(40) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("failed to create table", err)
	}
	
	// table to store files meta data
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS userfiles (
        fileid SERIAL PRIMARY KEY,
        filename VARCHAR(100) NOT NULL,
        url TEXT NOT NULL,
        uploaded_at TIMESTAMP NOT NULL,
        expiry_at TIMESTAMP NOT NULL,
        filesize INT NOT NULL,
        userid INT NOT NULL,
        CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES users(userid)
        ON DELETE CASCADE

    );
`)

	if err != nil {
		log.Fatal("failed to create table", err)
	}

	fmt.Println("Table created..")
}