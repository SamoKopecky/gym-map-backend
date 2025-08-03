package fileio

import (
	"gym-map/config"
	"gym-map/model"
	"os"
	"path/filepath"
)

const MAP_NAME = "map.svg"

type FloorMap struct {
	Config config.Config
}

func (fm FloorMap) getMapPath() string {
	return filepath.Join(fm.Config.MapFileRepository, MAP_NAME)
}

func (fm FloorMap) GetMap() (floorMap model.FloorMap, err error) {
	content, err := os.ReadFile(fm.getMapPath())
	if err != nil {
		return
	}
	floorMap = model.FloorMap(content)
	return

}

func (fm FloorMap) SaveMap(floorMap model.FloorMap) (err error) {
	err = os.WriteFile(filepath.Join(fm.getMapPath()), []byte(floorMap), 0644)
	if err != nil {
		return
	}

	return
}
