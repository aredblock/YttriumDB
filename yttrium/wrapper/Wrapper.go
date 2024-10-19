package yttrium_wrapper

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	yttrium_utils "github.com/aredblock/yttriumDB/yttrium/utils"
)

const (
	DataFolder = "./data/"
)

func LoadColum(name string) yttrium_utils.LoadedColumn {
	loadedColumn := yttrium_utils.LoadedColumn{}
	targetFile := DataFolder + name

	if _, err := os.Stat(targetFile); !errors.Is(err, os.ErrNotExist) {
		bytes, _ := os.ReadFile(targetFile)
		err := json.Unmarshal(bytes, &loadedColumn.Column)
		if err != nil {
			panic(err)
		}

		loadedColumn.Exists = true
		return loadedColumn
	}
	loadedColumn.Exists = false
	return loadedColumn
}

func UpdateColum(name string, column yttrium_utils.Column) {
	targetFile := DataFolder + name

	bytes, _ := json.MarshalIndent(column, "", "\t")
	err := os.WriteFile(targetFile, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Setup() {
	if _, err := os.Stat(DataFolder); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(DataFolder, os.ModePerm)
		log.Println("The database folder did not exist and was therefore recreated")
	}
}
