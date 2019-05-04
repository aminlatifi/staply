package main

import (
	"log"
	"net/http"
	"strings"

	receiver "./receiver"
)

var tmpDir = "/home/amin/tmp"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	method := r.Method

	if method == "POST" {
		var f func(*http.Request, string) (int, error)

		if strings.HasPrefix(contentType, "multipart/form-data") {
			f = receiver.MultipartReceiver
		} else if strings.HasPrefix(contentType, "application/json") {
			f = receiver.JSONReceiver
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		code, err := f(r, tmpDir)
		if err != nil {
			http.Error(w, err.Error(), code)
		}

		w.WriteHeader(http.StatusOK)
	} else if method == "GET" {
		w.WriteHeader(http.StatusOK)
	}

	// fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
