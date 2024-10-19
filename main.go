package main

import (
	"log"
	"os"
	"strings"

	"github.com/aredblock/yttriumDB/yttrium"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	arguments := os.Args

	if strings.Contains(arguments[0], "yttriumDB") {
		arguments = arguments[1:]
	}

	printInformations()

	yttrium.Launch(arguments)
}

func printInformations() {
	myFigure := figure.NewFigure("YttriumDB", "", true)
	myFigure.Print()
	log.Println("  ")
}
