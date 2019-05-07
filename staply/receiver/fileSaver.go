package receiver

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"staply/preview"
	"sync"
)

func createFile(dir, fileName string) (file *os.File, err error) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	// TODO: check a file exists with same name
	file, err = os.Create(path.Join(dir, fileName))
	return
}

func fileSaver(src io.Reader, fileName, saveDir, previewDir string) error {
	destFile, err := createFile(saveDir, fileName)
	if err != nil {
		return err
	}
	defer destFile.Close()

	var previewFile *os.File
	if previewDir != "" {
		previewFile, err = createFile(previewDir, fileName+"-preview.jpg")
		if err != nil {
			return err
		}
		defer previewFile.Close()
	}

	if previewDir != "" {
		data, _ := ioutil.ReadAll(src)
		var wg sync.WaitGroup
		wg.Add(1)
		var errSave error

		go func() {
			// Write contents of uploaded file to destFile
			src1 := bytes.NewReader(data)
			_, errSave = io.Copy(destFile, src1)
			wg.Done()
		}()

		src2 := bytes.NewReader(data)
		err = preview.MakePreview(src2, previewFile, 100, 100)

		wg.Wait()

		if errSave != nil {
			return errSave
		}

	} else {
		_, err = io.Copy(destFile, src)
	}

	return err
}
