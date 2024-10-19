package yttrium_web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	version string = "v1"
)

var (
	ShowChangedValues bool
)

func Start(port string, showChangedValues string) {
	path := "/api/" + version

	ShowChangedValues, _ = strconv.ParseBool(showChangedValues)

	http.HandleFunc(path+"/pop", pop)
	http.HandleFunc(path+"/push", push)

	//about
	http.HandleFunc(path+"/about", about)
	http.HandleFunc(path+"/", about)

	http.ListenAndServe(":"+port, nil)
}

func PrintInformation(port string) {
	log.Println(fmt.Sprintf("The web server starts at the URL: http://localhost:%s/api/%s", port, version))
}
