package yttrium

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/google/uuid"
)

func loadArguments() Arguments {
	arguments := Arguments{
		"46676",
		strings.Split(uuid.New().String(), "-")[0],
		"false",
		"true",
	}

	if _, err := os.Stat(ConfigFile); errors.Is(err, os.ErrNotExist) {
		//Create new
		bytes, _ := json.MarshalIndent(arguments, "", "\t")
		err := os.WriteFile(ConfigFile, bytes, os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		//Load existing
		bytes, _ := os.ReadFile(ConfigFile)
		err := json.Unmarshal(bytes, &arguments)
		if err != nil {
			panic(err)
		}
	}

	return arguments
}
