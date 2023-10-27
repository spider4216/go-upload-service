package structures

import (
	"mime/multipart"
)

type UploadRequest struct {
	File multipart.File
	Headers *multipart.FileHeader
	Name string
	Err error
	Method string
}