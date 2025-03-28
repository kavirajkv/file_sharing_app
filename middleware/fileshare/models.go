package fileshare

import (
	"time"
)

type Response struct {
	Message string `json:"message"`
}

type FileUploadResponse struct {
	Message    string    `json:"message"`
	Url        string    `json:"url"`
	Filename   string    `json:"filename"`
	UploadedAt time.Time `json:"uploadedat"`
	Expiresat  time.Time `json:"expiresat"`
	Size       int64     `json:"size"`
	Userid     int 	
}
