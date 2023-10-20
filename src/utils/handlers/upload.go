package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"messages"
	"strings"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, messages.NOT_PERMITED)
		return
	}

	file, headers, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintf(w, messages.ERR_GET_FILE)
		return
	}

	defer file.Close()

	filename := headers.Filename

	if inputName := r.FormValue("name"); inputName != "" {
		lastIndex := strings.LastIndex(filename, ".")
		filename = inputName + "." + filename[lastIndex+1:]
	}

	path := "../storage/" + filename
	dst, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(w, messages.ERR_CREATE_FILE)
		return
	}

	io.Copy(dst, file)

	fmt.Fprintf(w, messages.UPLOAD_OK_RESPONSE, path)
}
