package yttrium

import (
	"log"

	yttrium_web "github.com/aredblock/yttriumDB/yttrium/web"
	yttrium_wrapper "github.com/aredblock/yttriumDB/yttrium/wrapper"
)

var (
	yttriumArguments Arguments
)

func Launch(arguments []string) {
	//Arguments
	yttriumArguments = parseArguments(arguments)
	printArguments()

	//Wrapper
	yttrium_wrapper.Setup()

	//Web
	yttrium_web.Key = yttriumArguments.Key
	yttrium_web.PrintInformation(yttriumArguments.Port)
	yttrium_web.Start(yttriumArguments.Port, yttriumArguments.ShowChangedValues)

	log.Println("Done")
}
