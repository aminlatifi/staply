package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {

	numberOfFiles := 4

	origFilesContent, err := createRandomOrigFiles(t, numberOfFiles)
	if err != nil {
		return
	}

	// Prepare body
	body := new(bytes.Buffer)

	body.WriteByte('{')
	for i := 0; i < numberOfFiles; i++ {
		fileContent := origFilesContent[i]
		// Write to file decoded value to json
		if i > 0 {
			body.WriteByte(',')
		}

		fmt.Fprintf(body, "\"%d\":\"%s\"", i, base64.StdEncoding.EncodeToString(fileContent))
	}
	body.WriteByte('}')

	// Create POST request
	req, err := http.NewRequest(http.MethodPost, "/", body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Get request does not return 200")
	}

	checkFilesWithOrigs(t, numberOfFiles, origFilesContent)
}
