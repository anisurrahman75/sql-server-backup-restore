package main

import (
	"context"
	"fmt"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"log"
	"net/http"
	"time"
)

var ctx = context.Background()

const (
	port         = ":8443"
	responseBody = "Hello, TLS!"
)

func main() {
	certFile := "/home/anisur/go/src/github.com/anisurrahman75/sql-server-backup-to-azurite/certs/server.crt"
	keyFile := "/home/anisur/go/src/github.com/anisurrahman75/sql-server-backup-to-azurite/certs/server.key"
	//cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	//if err != nil {
	//	log.Fatalf("Failed to load X509 key pair: %v", err)
	//}
	//
	//config := &tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//}

	router := http.NewServeMux()
	router.HandleFunc("/", handleRequest)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	timeout := 20 * time.Second
	sctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	url := "https://localhost:8443"
	go func() {
		c := http.Client{Timeout: 100 * time.Millisecond}
		t := time.NewTicker(200 * time.Millisecond)
		for {
			select {
			case <-t.C:
				resp, _ := c.Head(url)
				if resp != nil {
					resp.Body.Close()
				}
			case <-sctx.Done():
				fmt.Errorf("proxy not ready in %s", timeout)
				return
			}
		}
	}()

	log.Printf("Listening on %s...", port)
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseBody))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
