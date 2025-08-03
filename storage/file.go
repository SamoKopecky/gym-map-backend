package storage

import (
	"gym-map/config"
	"gym-map/store"
	"os"
	"path/filepath"
)

type FileStorage struct {
	Config config.Config
}

func (fs FileStorage) getPath(fileType store.FileType, name string) string {
	return filepath.Join(fs.Config.StorageLocalPath, string(fileType), name)
}

func (fs FileStorage) Read(fileType store.FileType, name string) (*os.File, error) {
	path := fs.getPath(fileType, name)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (fs FileStorage) Write(fileType store.FileType, data []byte, name string) error {
	path := fs.getPath(fileType, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (fs FileStorage) Remove(fileType store.FileType, name string) error {
	path := fs.getPath(fileType, name)
	return os.Remove(path)
}
