package fileio

import (
	"fmt"
	"gym-map/config"
	"gym-map/model"
	"os"
	"path/filepath"
)

const MAP_NAME = "map.svg"

type FloorMap struct {
	Config config.Config
}

func (fm FloorMap) GetMap() (floorMap model.FloorMap, err error) {
	path := filepath.Join(fm.Config.MapFileRepository, MAP_NAME)
	fmt.Println(path)
	fmt.Println(fm.Config.MapFileRepository)
	content, err := os.ReadFile(filepath.Join(fm.Config.MapFileRepository, MAP_NAME))
	if err != nil {
		return
	}
	floorMap = model.FloorMap(content)
	return

}
