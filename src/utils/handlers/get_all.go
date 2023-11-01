package handlers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("../storage")
	
	if (err != nil) {
		fmt.Fprintf(w, "Something went wrong while reading dir")
		return
	}
	
	content := make([]string, 10)
	
	for _, item := range files {
		content = append(content, item.Name())
	}
	
	fmt.Fprintf(w, strings.Join(content, ","))
}