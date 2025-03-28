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

// to get files details
type Files struct {
	File_id    int       `json:"file_id"`
	Filename   string    `json:"filename"`
	Url        string    `json:"url"`
	UploadedAt time.Time `json:"uploadedat"`
	Expiresat  time.Time `json:"expiresat"`
	Size       int64     `json:"size"`
}

// request struct to share file
type ShareFileRequest struct {
	File_id int `json:"file_id"`
}

type ShareFileResponse struct {
	File_id   int       `json:"file_id"`
	Url       string    `json:"url"`
	Expiresat time.Time `json:"expires_at"`
}

type SearchFileRequest struct {
	Filename string `json:"filename"`
}