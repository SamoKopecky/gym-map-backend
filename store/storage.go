package store

import (
	"os"
)

type FileType string

const (
	MEDIA FileType = "media"
	MAP   FileType = "map"
)

type FileStorage interface {
	Write(fileType FileType, data []byte, name string) error
	Read(fileType FileType, name string) (*os.File, error)
	Remove(FileType FileType, name string) error
}
