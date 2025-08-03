package storage

import (
	"gym-map/config"
	"gym-map/model"
	"gym-map/store"
)

const MAP_NAME = "map.svg"

type FloorMap struct {
	Config  config.Config
	Storage store.FileStorage
}

func (fm FloorMap) GetMap() (floorMap model.FloorMap, err error) {
	file, err := fm.Storage.Read(store.MAP, MAP_NAME)

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stats.Size()
	data := make([]byte, size)

	_, err = file.Read(data)
	if err != nil {
		return
	}
	defer file.Close()

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
