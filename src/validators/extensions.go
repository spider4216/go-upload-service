package validators

func IsExtensionOk(ext string) bool {
	extensions := [2]string{"jpg", "png"}
	
	for _, value := range extensions {
		if value == ext {
			return true
		}
	}
	
	return false
}