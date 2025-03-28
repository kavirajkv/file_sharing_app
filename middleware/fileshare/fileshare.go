package fileshare

import (
	"context"
	"encoding/json"
	"fileshare/db"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// function to check server status
func Status(w http.ResponseWriter, r *http.Request) {
	msg := Response{Message: "Server is up and running"}
	json.NewEncoder(w).Encode(msg)
}

// function to upload file
func Uploadfile(w http.ResponseWriter, r *http.Request) {
	db := db.ConnectDB()
	defer db.Close()

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//username fetched from context send from authenticaion middleware
	username, ok := r.Context().Value("username").(string)
	if !ok {
		http.Error(w, "Username not found in context", http.StatusInternalServerError)
		return
	}

	userchan := make(chan int)
	filechan := make(chan *manager.UploadOutput)
	errchan := make(chan error, 2)


	go func() {

		row := db.QueryRow(`SELECT userid FROM users WHERE username=$1`, username)
		var userid int
		err = row.Scan(&userid)
		if err != nil {
			errchan <- err
			return
		}
		userchan <- userid
	}()

	go func() {
		uploader := manager.NewUploader(S3Client)

		result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
			Key:    aws.String(header.Filename),
			Body:   file,
		})
		if err != nil {
			errchan <- err
			return
		}
		filechan <- result
	}()
	
	var userid int
	var result *manager.UploadOutput

	//Ref:- go by example (handling multiple channels)
	for i := 0; i < 2; i++ {
        select {
		case user := <-userchan:
			userid = user
        case file := <-filechan:
			result = file
		case er:= <-errchan:
			err = er
        }
    }

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	

	//expire after 2 days (s3 lifecycle policy created to delete files after 2 days)
	uploadedat := time.Now()
	expiresat := time.Now().Add(time.Hour * 48)

	//inserting file meta data into database
	_, err = db.Exec(`INSERT INTO userfiles (filename, url, uploaded_at, expiry_at, filesize, userid) VALUES ($1,$2,$3,$4,$5,$6)`, header.Filename, result.Location, uploadedat, expiresat, header.Size, userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := FileUploadResponse{Message: "File uploaded successfully", Url: result.Location, Filename: header.Filename, UploadedAt: uploadedat, Expiresat: expiresat, Size: header.Size, Userid: userid}

	json.NewEncoder(w).Encode(msg)
}

