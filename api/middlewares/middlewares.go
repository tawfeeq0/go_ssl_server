package middlewares

import (
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s%s %s", r.Method, r.Host, r.RequestURI, r.Proto)
		next(w, r)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applciation/json")
		next(w, r)
	}
}
