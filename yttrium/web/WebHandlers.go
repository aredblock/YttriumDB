package yttrium_web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	yttrium_utils "github.com/aredblock/yttriumDB/yttrium/utils"
	yttrium_wrapper "github.com/aredblock/yttriumDB/yttrium/wrapper"
)

// Add securety!!!
func pop(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" || !KeyExist(req) || !CheckKey(req) {
		return
	}

	loadedRow := yttrium_wrapper.LoadColum(name)
	bytes, _ := json.MarshalIndent(loadedRow, "", "\t")
	fmt.Fprintf(w, string(bytes))
}

// Add securety!!!
func push(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	column := yttrium_utils.Column{}

	if name == "" || !KeyExist(req) || !CheckKey(req) {
		return
	}

	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, &column)

	//Print changed values
	if ShowChangedValues {
		for _, row := range column.Rows {
			log.Println(fmt.Sprintf("%s -> '%s':'%s'", name, row.Key, row.Value))
		}
	}

	yttrium_wrapper.UpdateColum(name, column)
}

func about(w http.ResponseWriter, req *http.Request) {
	about := yttrium_utils.About{
		"YttriumDB",
		version,
		"amarena",
	}

	bytes, _ := json.MarshalIndent(about, "", "\t")
	fmt.Fprintf(w, string(bytes))
}
