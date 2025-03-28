package fileshare

import (
	"net/http"
	// "fileshare/db"
	"encoding/json"
)

//function to check server status
func Status(w http.ResponseWriter, r *http.Request) {
	msg:=Response{Message: "Server is up and running"}
	json.NewEncoder(w).Encode(msg)
}

