package yttrium

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

const (
	ConfigFile = "./config.json"
)

type Arguments struct {
	Port string

	Key     string
	ShowKey string

	ShowChangedValues string
}

func parseArguments(arguments []string) Arguments {
	parsedArguments := loadArguments()
	parsedArgumentsVal := reflect.ValueOf(&parsedArguments).Elem()
	parsedArgumentsType := parsedArgumentsVal.Type()

	for _, argument := range arguments {
		parts := strings.Split(argument, "=")
		if len(parts) != 2 {
			continue
		}
		name := strings.TrimPrefix(parts[0], "--")
		value := parts[1]

		for i := 0; i < parsedArgumentsVal.NumField(); i++ {
			field := parsedArgumentsType.Field(i)

			if strings.ToLower(field.Name) == name {
				parsedArgumentsVal.Field(i).SetString(value)
				break
			}
		}
	}

	return parsedArguments
}

func printArguments() {
	log.Println("YttriumDB starts with the information: ")
	log.Println(" - port=" + yttriumArguments.Port)

	showKey, _ := strconv.ParseBool(yttriumArguments.ShowKey)
	if showKey {
		log.Println(" - key=" + yttriumArguments.Key)
	} else {
		log.Println(" - key=invisable")
	}
}
