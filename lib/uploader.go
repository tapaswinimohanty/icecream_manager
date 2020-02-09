package lib

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func Upload(file *multipart.FileHeader) (*string, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Destination
	uploadFile := fmt.Sprintf(`/files/live/sites/systemsite/files/flavors/products/us/pint/open-closed-pints/%v`, file.Filename)
	dst, err := os.Create("public" + uploadFile)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return &uploadFile, nil

}
