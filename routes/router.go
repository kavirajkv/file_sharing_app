package routes

import (
	"fileshare/middleware/auth"
	"fileshare/middleware/fileshare"
	"golang.org/x/time/rate" 
	"net/http"
	"time"
)

func Router() *http.ServeMux {
	r:=http.NewServeMux()

	

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
	r.HandleFunc("GET /status",fileshare.Status)
	
	// // routes for user authentication
	r.Handle("POST /signup", rateLimitMiddleware(http.HandlerFunc(auth.Signup)))
	r.Handle("POST /login", rateLimitMiddleware(http.HandlerFunc(auth.Login)))

	// // routes for file sharing
	r.Handle("POST /upload", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.Uploadfile))))
	r.Handle("GET /files", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.GetFiles))))
	r.Handle("GET /share", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.ShareFile))))
	r.Handle("DELETE /delete", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.DeleteFile))))
	r.Handle("GET /search", rateLimitMiddleware(http.HandlerFunc(auth.Authenticate(fileshare.SearchFile))))

	return r
}
