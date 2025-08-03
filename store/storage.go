package store

import (
	"os"
)

type FileType string

const (
	VIDEO FileType = "video"
	IMAGE FileType = "image"
	MAP   FileType = "map"
)

type FileStorage interface {
	Write(fileType FileType, data []byte, name string) error
	Read(fileType FileType, name string) (*os.File, error)
}
