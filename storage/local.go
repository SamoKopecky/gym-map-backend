package storage

import (
	"gym-map/config"
	"gym-map/store"
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Config config.Config
}

func (ls LocalStorage) getPath(fileType store.FileType, name string) string {
	return filepath.Join(ls.Config.StorageLocalPath, string(fileType), name)
}

func (ls LocalStorage) Read(fileType store.FileType, name string) (*store.StorageObject, error) {
	path := ls.getPath(fileType, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &store.StorageObject{
		ReadSeekCloser: file,
		FileInfo:       stat,
	}, nil
}

func (ls LocalStorage) Write(fileType store.FileType, data io.ReadSeeker, name string) error {
	path := ls.getPath(fileType, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, data)
	if err != nil {
		return err
	}

	return nil
}

func (ls LocalStorage) Remove(fileType store.FileType, name string) error {
	path := ls.getPath(fileType, name)
	return os.Remove(path)
}
