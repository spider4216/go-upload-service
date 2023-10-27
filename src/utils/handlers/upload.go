package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"messages"
	"strings"
	"validators"
	"structures"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var request structures.UploadRequest

	request.Method = r.Method
	request.Name = r.FormValue("name")

	if request.Method != "POST" {
		fmt.Fprintf(w, messages.NOT_PERMITED)
		return
	}

	request.File, request.Headers, request.Err = r.FormFile("file")

	if request.Err != nil {
		fmt.Fprintf(w, messages.ERR_GET_FILE)
		return
	}
	
	defer request.File.Close()

	filename := strings.ToLower(request.Headers.Filename)
	lastIndex := strings.LastIndex(filename, ".")
	extension := filename[lastIndex+1:]

	if ! validators.IsExtensionOk(extension) {
		fmt.Fprintf(w, messages.EXT_NOT_ALLOWED)
		return;
	}

	if request.Name != "" {
		filename = strings.ToLower(request.Name) + "." + extension
	}

	path := "../storage/" + filename
	dst, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(w, messages.ERR_CREATE_FILE)
		return
	}

	io.Copy(dst, request.File)

	fmt.Fprintf(w, messages.UPLOAD_OK_RESPONSE, path)
}
