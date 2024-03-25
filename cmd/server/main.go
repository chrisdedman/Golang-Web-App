package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world!\n")
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Request handled successfully")
	})
}

func setupHandlers(mux *http.ServeMux) {
	mux.Handle("/", loggingMiddleware(http.FileServer(http.Dir("static"))))
	mux.Handle("/api", loggingMiddleware(http.HandlerFunc(apiHandler)))
	mux.Handle("/healthz", loggingMiddleware(http.HandlerFunc(healthCheckHandler)))
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	mux := http.NewServeMux()
	fmt.Printf("Starting server on http://localhost%s\n", listenAddr)

	setupHandlers(mux)

	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
