package receiver

import (
	"io"
	"os"
	"path"
)

func fileSaver(src io.Reader, saveFolder string, dstFileName string) error {
	// TODO: check a file exists with same name
	destFile, err := os.Create(path.Join(saveFolder, dstFileName))
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Write contents of uploaded file to destFile
	if _, err = io.Copy(destFile, src); err != nil {
		return err
	}

	return nil
}
