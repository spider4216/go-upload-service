package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/v1/upload", handlerUpload)
	http.ListenAndServe("localhost:8000", nil)
}

func handlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Println("Operation not permited. Use POST")
		return
	}

	file, headers, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Something went wrong while getting file")
		return
	}

	defer file.Close()

	filename := headers.Filename
	path := "../storage/" + filename
	dst, err := os.Create(path)

	if err != nil {
		fmt.Printf("Something went wrong while creating file: $v\n", err)
		return
	}

	io.Copy(dst, file)

	fmt.Println("Api Upload Response")
}
