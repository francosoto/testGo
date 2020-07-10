package main

import (
	"fmt"
	"net/http"
	"crypto/tls"
	"log"
)

var (
	CertFile 	= os.Getenv("CERT_FILE")
	KeyFile 	= os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

const message = "hello, world\n"

func main() {
	fmt.Printf(message)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := NewServer(mux, ServiceAddr)

	// HTTP/2 is enabled automatically on any Go 1.6+ server if the request is served over TLS/HTTPS
	err := srv.ListenAndServeTLS(CertFile, KeytFile)
	if err != nil {
		log.Fatalf("Server Failed: %", err)
	}
}

func NewServer(mux *http.ServerMux, serverAddress string) *http.Server {
	// https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreference: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305, 
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},

	}

	srv := &http.Server{
		Addr: 			serverAddress,
		ReadTimeout:  	5 * time.Second,
		WriteTimeout: 	10 * time.Second,
		IdleTimeout: 	120 * time.Second,
		TLSConfig: 		tlsConfig,
		Handler: 		mux
	}

	return srv
}