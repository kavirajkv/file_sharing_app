package routes

import (
	"fileshare/middleware/auth"
	"github.com/gorilla/mux"
	"fileshare/middleware/fileshare"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// route to check server status
	r.HandleFunc("/status", fileshare.Status).Methods("GET")
	
	// routes for user authentication
	r.HandleFunc("/signup", auth.Signup).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	// routes for file sharing
	r.HandleFunc("/upload", auth.Authenticate(fileshare.Uploadfile)).Methods("POST")
	r.HandleFunc("/files", auth.Authenticate(fileshare.GetFiles)).Methods("GET")
	r.HandleFunc("/share", auth.Authenticate(fileshare.ShareFile)).Methods("GET")
	r.HandleFunc("/delete", auth.Authenticate(fileshare.DeleteFile)).Methods("DELETE")

	return r
}
