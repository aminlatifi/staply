package receiver

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// JSONReceiver returns code and error
func JSONReceiver(r *http.Request, saveDir, previewDir string) (code int, err error) {
	whole, err := ioutil.ReadAll(r.Body)
	//decoder := json.NewDecoder(r.Body)
	data := make(map[string]string)
	//err = decoder.Decode(&data)
	err = json.Unmarshal(whole, &data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	for key, value := range data {
		fileContent, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return http.StatusBadRequest, err
		}
		fileSaver(bytes.NewReader(fileContent), key, saveDir, previewDir)
	}

	return
}
