package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var stdLog, errLog *log.Logger

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errLog.Fatal("Error 500!")
}

func toJSON(model Profile) string {
	jsonProfile, err := json.Marshal(model)
	if err != nil {
		errLog.Fatal("error!!!!!")
		return "nil"
	}
	return string(jsonProfile)
}

func customHandler(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Vladimir", []string{"basketball", "guitar"}}
	fmt.Fprintln(w, toJSON(profile))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func main() {

	stdLog = log.New(os.Stdout, "INFO", 0)
	stdLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	errLog = log.New(os.Stderr, "ERROR", 1)
	errLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/123/321", customHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/", rootHandler)
	// http.ListenAndServe(":4000", nil)
	err := http.ListenAndServe(":8080", logRequest(http.DefaultServeMux))

	if err != nil {
		log.Fatal(err)
	}
}
