package fabrics

import (
	"structures"
	"net/http"
)

func BuildUploadRequest(r *http.Request) structures.UploadRequest {
	var request structures.UploadRequest

	request.Method = r.Method
	request.Name = r.FormValue("name")
	request.File, request.Headers, request.Err = r.FormFile("file")
	
	return request
}