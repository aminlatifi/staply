package receiver

import (
	"log"
	"net/http"
)

// MultipartReceiver returns code and error
func MultipartReceiver(r *http.Request, saveDir, previewDir string) (code int, err error) {
	r.ParseMultipartForm(1024)

	if len(r.MultipartForm.File) == 0 {
		log.Fatal("MultipartForm does not have file attached!")
		code = http.StatusBadRequest
		err = http.ErrMissingFile
		return
	}

	for key := range r.MultipartForm.File {
		f, fh, err := r.FormFile(key)

		if err != nil {
			return http.StatusBadRequest, err
		}
		defer f.Close()

		err = fileSaver(f, fh.Filename, saveDir, previewDir)
		if err != nil {
			return http.StatusInternalServerError, err
		}

	}
	return
}
