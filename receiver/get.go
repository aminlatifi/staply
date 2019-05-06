package receiver

import (
	"errors"
	"net/http"
	"path"
)

// GetReceiver returns code and error
func GetReceiver(r *http.Request, saveDir, previewDir string) (code int, err error) {
	params := r.URL.Query()
	if len(params) != 1 {
		code = http.StatusBadRequest
		err = errors.New("One and onlye one parameter is required")
		return
	}
	var url string
	for _, value := range params {
		url = value[0]
		break
	}
	newResponse, err := http.Get(url)
	if err != nil {
		if newResponse != nil {
			code = newResponse.StatusCode
		} else {
			code = http.StatusNotAcceptable
		}

		return code, err

	}

	if newResponse.StatusCode == http.StatusOK {
		defer newResponse.Body.Close()
		fileSaver(newResponse.Body, path.Base(url), saveDir, previewDir)

	} else {
		return code, errors.New("Status code is not 200")
	}
	return
}
