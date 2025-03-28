package auth

import (
	"net/http"
	"fileshare/db"
	"encoding/json"
	"github.com/kavirajkv/security/digest"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
)


// user handlers 

//signup
func Signup(w http.ResponseWriter, r *http.Request) {
	db:=db.ConnectDB()
	defer db.Close()


	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username, email and password are required", http.StatusBadRequest)
		return
	}

	//hashing the password using sha256 from my(kavirajkv) security package
	pass:=digest.ShaDigest(user.Password)

	_, err = db.Exec(`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`, user.Username, user.Email, pass)
	if err != nil {
		http.Error(w, "Internal error while execting insertion", http.StatusInternalServerError)
		return
	}
	msg:=Response{Message: "User created successfully"}
	json.NewEncoder(w).Encode(msg)
}


//login
func Login(w http.ResponseWriter, r *http.Request) {
	db:=db.ConnectDB()
	defer db.Close()

	var user LoginRequest
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	row:= db.QueryRow(`SELECT username, password FROM users WHERE username=$1`, user.Username)

	var usercheck LoginRequest
	err = row.Scan(&usercheck.Username, &usercheck.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	pass:=digest.ShaCheck(usercheck.Password,user.Password,)
	if !pass{
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//here after valid credentials, generate the jwt token and send it to the user (token is set to valid for 10 minutes)
	claim:=Claims{Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10))},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Value: tokenString,
		Expires: time.Now().Add(time.Minute * 10),
	})

	msg:=Response{Message: "User logged in successfully"}

	json.NewEncoder(w).Encode(msg)
		
}
