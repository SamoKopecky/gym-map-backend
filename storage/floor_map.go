package storage

import (
	"gym-map/config"
	"gym-map/model"
	"gym-map/store"
	"io"
)

const MAP_NAME = "map.svg"

type FloorMap struct {
	Config  config.Config
	Storage store.FileStorage
}

func (fm FloorMap) GetMap() (floorMap model.FloorMap, err error) {
	file, err := fm.Storage.Read(store.MAP, MAP_NAME)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	floorMap = model.FloorMap(data)
	return

}

func (fm FloorMap) SaveMap(floorMap model.FloorMap) error {
	err := fm.Storage.Write(store.MAP, []byte(floorMap), MAP_NAME)
	if err != nil {
		return err
	}

	return nil
}
