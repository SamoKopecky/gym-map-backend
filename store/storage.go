package store

import (
	"io"
	"os"
)

type FileType string

const (
	MEDIA FileType = "media"
	MAP   FileType = "map"
)

type StorageObject struct {
	io.ReadSeekCloser
	os.FileInfo
}

type FileStorage interface {
	Read(fileType FileType, name string) (*StorageObject, error)
	Write(fileType FileType, data io.ReadSeeker, name string) error
	Remove(FileType FileType, name string) error
}
