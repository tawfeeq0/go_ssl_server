package middlewares

import (
	"github.com/tawfeeq0/go_ssl_server/utils/console"
	"log"
	"net/http"
	"github.com/tawfeeq0/go_ssl_server/models/auth"
	"github.com/tawfeeq0/go_ssl_server/api/responses"
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

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValidate(r)
		if err != nil{
			log.Println(err)
			responses.ERROR(w,http.StatusUnauthorized,err)
			return
		}
		uid,err := auth.ExtractTokenID(r)
		if err != nil {
			responses.ERROR(w,http.StatusUnauthorized,err)
			return
		}
		console.Pretty(uid)
		next(w, r)
	}
}