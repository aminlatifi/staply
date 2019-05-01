package main

import (
	"log"
	"net/http"
	"strings"

	receiver "./receiver"
)

var tmpDir = "/home/amin/tmp"

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	method := r.Method
	if method == "POST" {
		if strings.HasPrefix(contentType, "multipart/form-data;") {
			code, err := receiver.MultipartReceiver(r, tmpDir)
			if err != nil {
				if err == http.ErrMissingFile {
					http.Error(w, "Request did not contain a file", code)
				} else {
					http.Error(w, err.Error(), code)
				}
			}

			return
		}

	}

	w.WriteHeader(http.StatusOK)

	// fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
