package helpers

import (
	"encoding/json"
	"structures"
)

func Response(status int, message string) string {
	response, _ := json.Marshal(structures.UploadResponse{Status: status,Text: message})
	
	return string(response)
}

func ResponseArr(status int, data []structures.File) string {
	response, _ := json.Marshal(data)
	
	return string(response)
}