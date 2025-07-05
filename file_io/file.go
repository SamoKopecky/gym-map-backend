package fileio

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, rootFilePath string) (name string, err error) {
	name = fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))

	// Source file
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(filepath.Join(rootFilePath, name))
	if err != nil {
		return
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return
	}

	return
}
