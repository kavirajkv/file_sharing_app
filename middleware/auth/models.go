package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

////////////
// Models //

// to store user details
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

// to store login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// to  store jwt token claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
