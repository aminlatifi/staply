package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMultipart(t *testing.T) {

	numberOfFiles := 4

	origFilesContent, err := createRandomOrigFiles(t, numberOfFiles)
	if err != nil {
		return
	}

	// Prepare body
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for i := 0; i < numberOfFiles; i++ {
		fileContent := origFilesContent[i]

		part, err := writer.CreateFormFile(strconv.Itoa(i), strconv.Itoa(i))
		if err != nil {
			t.Errorf("Error creating form file %d", i)
		}

		part.Write(fileContent)
	}
	err = writer.Close()

	// Create POST request
	req, err := http.NewRequest("POST", "/", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
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
