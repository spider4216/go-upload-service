package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Operation not permited. Use POST")
		return
	}

	file, headers, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintf(w, "Something went wrong while getting file")
		return
	}

	defer file.Close()

	filename := headers.Filename
	path := "../storage/" + filename
	dst, err := os.Create(path)

	if err != nil {
		fmt.Fprintf(w, "Something went wrong while creating file: $v")
		return
	}

	io.Copy(dst, file)

	fmt.Fprintf(w, "upload ok: $q", path)
}
