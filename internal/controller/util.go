package controller

func buildResx(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"data":   data,
		"status": "success",
	}
}
func buildResErr(data string) map[string]interface{} {
	return map[string]interface{}{
		"error":  data,
		"data":   nil,
		"status": "error",
	}
}
