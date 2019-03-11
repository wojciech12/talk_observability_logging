package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	 log "github.com/sirupsen/logrus"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello!")

	log.WithFields(log.Fields{
    	"method": r.Method,
    	"handler": "hello",
	}).Info("hello!")
}

func WorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "World!")

	log.WithFields(log.Fields{
    	"method": r.Method,
    	"handler": "world",
	 }).Info("world!")
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bye!")

	log.WithFields(log.Fields{
		"method": r.Method,
		"handler": "error",
	}).Error("What does 'bye' mean?!")
}

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	r := mux.NewRouter()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/world", WorldHandler)
	r.HandleFunc("/error", ErrorHandler)
	http.ListenAndServe(":8080", r)
}
