package main

import (
	"crypto/rand"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"reflect"
	"strconv"
	"testing"
)

var testFileSize = 4 * 1024

func TestGet(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code == http.StatusOK {
		t.Error("Should not accept request without parameter")
	}

	return
}

// Create #numberOfFiles number of random files with names "0", "1", ....
// Returns temporary directory path and error
func createRandomOrigFiles(t *testing.T, numberOfFiles int) (origFilesContent [][]byte, err error) {

	saveDir, err = ioutil.TempDir("/tmp", "staply_")

	if err != nil {
		t.Error("Failed to create temporary directory")
		return
	}

	if nil != os.Mkdir(path.Join(saveDir, "orig"), os.ModePerm) {
		t.Error("Failed to create orig directory")
		return
	}

	origFilesContent = make([][]byte, numberOfFiles)

	token := make([]byte, testFileSize)
	for i := 0; i < numberOfFiles; i++ {
		rand.Read(token)
		origFilesContent[i] = token
		err = ioutil.WriteFile(path.Join(saveDir, "orig", strconv.Itoa(i)), token, 0644)
		if err != nil {
			t.Error("Failed to create orig file")
			return
		}
	}

	return
}

func checkFilesWithOrigs(t *testing.T, numberOfFiles int, origFilesContent [][]byte) {
	for i := 0; i < numberOfFiles; i++ {
		func() {
			file, err := os.Open(path.Join(saveDir, strconv.Itoa(i)))
			if err != nil {
				t.Errorf("Error opening file %d", i)
				return
			}
			defer file.Close()
			content, err := ioutil.ReadAll(file)

			if err != nil {
				t.Errorf("Error reading file %d", i)
				return
			}

			if !reflect.DeepEqual(origFilesContent[i], content) {
				t.Errorf("Content of %d files does not match to orig one", i)
			}
		}()
	}
}
