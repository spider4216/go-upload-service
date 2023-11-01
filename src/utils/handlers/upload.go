package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"messages"
	"strings"
	"validators"
	"fabrics"
	"encoding/json"
	"structures"
)

const (
	HTTP_POST_METHOD = "POST"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	request := fabrics.BuildUploadRequest(r)

	if request.Method != HTTP_POST_METHOD {
		fmt.Fprintf(w, messages.NOT_PERMITED)
		return
	}

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
	
	response, err := json.Marshal(structures.UploadResponse{Status: 200,Text: messages.UPLOAD_OK_RESPONSE})
	
	if err != nil {
		fmt.Fprintf(w, "something went wrong")
		return
	}

	fmt.Fprintf(w, string(response))
}
