package utils

import (
	"net/http"
)

func SendJSONError(w http.ResponseWriter, err error, statusCode int) {
	SendJSONResponse(w, map[string]string{"error": err.Error()}, statusCode)
}
