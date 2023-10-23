package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"messages"
	"strings"
	"validators"
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

	filename := strings.ToLower(headers.Filename)
	lastIndex := strings.LastIndex(filename, ".")
	extension := filename[lastIndex+1:]

	if ! validators.IsExtensionOk(extension) {
		fmt.Fprintf(w, messages.EXT_NOT_ALLOWED)
		return;
	}

	if inputName := r.FormValue("name"); inputName != "" {
		filename = strings.ToLower(inputName) + "." + extension
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
