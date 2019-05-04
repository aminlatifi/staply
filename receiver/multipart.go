package receiver

import (
	"log"
	"net/http"
)

// MultipartReceiver returns code and error
func MultipartReceiver(r *http.Request, saveFolder string) (code int, err error) {
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

		err = fileSaver(f, saveFolder, fh.Filename)
		if err != nil {
			return http.StatusInternalServerError, err
		}

	}
	return
}
