package pages

import (
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

const message string = "Hello KFUPM 2020!\n"

//Handlers : handler
type Handlers struct {
	logger *log.Logger
}

//Home : method
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

//About : method
func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text-plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("About"))
}

// Logger : method
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

//SetupRoutes : method
func (h *Handlers) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/", h.Logger(h.Home)).Methods("POST")
	router.HandleFunc("/about", h.Logger(h.About)).Methods("GET")
}

//NewHandler : method
func NewHandler(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
