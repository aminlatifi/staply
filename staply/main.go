package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	receiver "staply/receiver"
)

var saveDir = "/downloads"
var previewDir = "/downloads/preview"

func httpHandler(w http.ResponseWriter, r *http.Request) {

	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	method := r.Method

	var f func(*http.Request, string, string) (int, error)

	if method == http.MethodPost {

		if strings.HasPrefix(contentType, "multipart/form-data") {
			f = receiver.MultipartReceiver
		} else if strings.HasPrefix(contentType, "application/json") {
			f = receiver.JSONReceiver
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

	} else if method == http.MethodGet {
		f = receiver.GetReceiver
	} else {
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		return
	}

	code, err := f(r, saveDir, previewDir)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", httpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
