package structures

type UploadResponse struct{
	Status int `json:"status"`
	Text string `json:"message"`
}