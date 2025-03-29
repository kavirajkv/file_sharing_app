package routes

import (
	"fileshare/middleware/auth"
	"github.com/gorilla/mux"
	"fileshare/middleware/fileshare"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	limiter := rate.NewLimiter(rate.Every(time.Minute), 100)
	//rate limiting middleware to allow only 100 request per minute as required 
	rateLimitMiddleware := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                http.Error(w, "Too many requests", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
	}

	// route to check server status
	r.HandleFunc("/status", fileshare.Status).Methods("GET")
	
	// routes for user authentication
	r.Handle("/signup", rateLimitMiddleware(http.HandlerFunc(auth.Signup))).Methods("POST")
	r.Handle("/login", rateLimitMiddleware(http.HandlerFunc(auth.Login))).Methods("POST")

	// routes for file sharing
	r.Handle("/upload", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.Uploadfile)))).Methods("POST")
	r.Handle("/files", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.GetFiles)))).Methods("GET")
	r.Handle("/share", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.ShareFile)))).Methods("GET")
	r.Handle("/delete", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.DeleteFile)))).Methods("DELETE")
	r.Handle("/search", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.SearchFile)))).Methods("GET")

	return r
}
