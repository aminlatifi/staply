package preview

import (
	"image"
	_ "image/gif" // Support gif
	"image/jpeg"
	_ "image/png" // Support png
	"io"

	_ "github.com/spakin/netpbm" // Support pgm
	_ "golang.org/x/image/bmp"   // Support bpm
	_ "golang.org/x/image/tiff"  // Support tiff

	_ "github.com/chai2010/webp" // Support webp
	_ "github.com/lmittmann/ppm" // Support ppm
	"github.com/nfnt/resize"
)

// MakePreview return resized image
func MakePreview(src io.Reader, w io.Writer, width, height uint) (err error) {
	orginalImage, _, err := image.Decode(src)
	if err != nil {
		return err
	}
	newImage := resize.Resize(width, height, orginalImage, resize.Lanczos3)
	err = jpeg.Encode(w, newImage, nil)
	return
}
