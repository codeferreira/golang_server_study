package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(begin))
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintln(w, "OK: "+id)
	})

	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      Log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  30 * time.Second,
		WriteTimeout:                 30 * time.Second,
		IdleTimeout:                  60 * time.Second,
		MaxHeaderBytes:               0,
	}

	log.Fatal(srv.ListenAndServe())
}
