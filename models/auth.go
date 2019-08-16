package models

import (
	"log"
	"net/http"
	"time"
	"strings"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tawfeeq0/go_ssl_server/services"
)


//AuthHandlers : handler
type AuthHandlers struct {
	logger *log.Logger
	auth *Auth
}
//Auth : asdsa
type Auth struct {
	Token string `json:"token"`
}

//Home : method
func (h *AuthHandlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("AUTH"))
}
//Verify : method
func (h *AuthHandlers) Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if (h.auth.Token == params["token"]) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h.auth)
		h.logger.Println("auth",h.auth)
		return
	}else {
		w.WriteHeader(http.StatusForbidden)
		return
	}
}


// Logger : method
func (h *AuthHandlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if(len(splitToken) < 2){
			h.logger.Println("forb")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h.logger.Println("comp")
		reqToken = splitToken[1]
		allowed := services.Verify(strings.TrimSpace(reqToken))
		if(!allowed){
			w.WriteHeader(http.StatusUnauthorized)
			return
		}else{
			next(w, r)
		}
	}
}

//SetupRoutes : method
func (h *AuthHandlers) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/auth/", h.Logger(h.Home)).Methods("GET")
	router.HandleFunc("/auth/{token}", h.Logger(h.Verify)).Methods("GET")
}

//NewHandler : method
func NewHandler(logger *log.Logger) *AuthHandlers {
	return &AuthHandlers{
		logger: logger,
		auth : &Auth{Token:"2"},
	}
}