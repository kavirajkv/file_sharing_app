package routes

import (
	"fileshare/middleware/auth"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", auth.Signup).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	return r
}
