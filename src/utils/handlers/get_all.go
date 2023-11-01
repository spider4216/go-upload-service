package handlers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"messages"
	"utils/helpers"
	"constants"
	"structures"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("../storage")
	
	if (err != nil) {
		fmt.Fprintf(w, helpers.Response(constants.HTTP_UNPROCESSABLE, messages.ERR_READING_DIR))
		return
	}
	
	content := make([]structures.File, 0)
	
	for _, item := range files {
		content = append(content, structures.File{Filename: item.Name()})
	}
	
	fmt.Fprintf(w, helpers.ResponseArr(constants.HTTP_OK, content))
}