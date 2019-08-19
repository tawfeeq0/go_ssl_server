package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	//	"runtime"
	"time"
	//	"github.com/tawfeeq0/go_ssl_server/pages"
	//	"github.com/tawfeeq0/go_ssl_server/models"
	//	"github.com/tawfeeq0/go_ssl_server/services"
	"github.com/gorilla/mux"
	"github.com/tawfeeq0/go_ssl_server/models/auto"
	"github.com/tawfeeq0/go_ssl_server/server/config"
)

// Run : method
func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("Listening [::]:%d..\n", config.PORT)
	listen(config.PORT)

}

func listen(port int) {
	router := NewRouter()
	srv := NewServer(router, fmt.Sprintf(":%d", port))
//	fmt.Println(config.CERTFILE, config.KEYFILE)
	log.Fatal(srv.ListenAndServeTLS(config.CERTFILE, config.KEYFILE))
}

/*
func Run(){
	logger := log.New(os.Stdout, "go_ssl_server ", log.LstdFlags|log.Lshortfile)
	fmt.Println(runtime.GOMAXPROCS(-1))
	services.Signin("77995")
	fmt.Println(services.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0MjMwODksInVzZXJuYW1lIjoiNzc5OTUifQ.xt_jdjfPg2XULVNMg6ioY0eBwv_izlR534OrcmSfLxo"))
	fmt.Println(services.GetUserInfo("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU0MjMwODksInVzZXJuYW1lIjoiNzc5OTUifQ.xt_jdjfPg2XULVNMg6ioY0eBwv_izlR534OrcmSfLxo"))
	router := NewRouter()
	h := pages.NewHandler(logger)
	hAuth := models.NewHandler(logger)
	h.SetupRoutes(router)
	hAuth.SetupRoutes(router)
	srv := NewServer(router, SericeAddress)
	logger.Println("server starting")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start : %v", err)
	}
}
*/
// New : method
func NewServer(router *mux.Router, serverAddress string) *http.Server {
	//See : https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}
	return srv
}
