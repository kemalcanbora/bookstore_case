package helper

import (
	"bookstore_case/models"
	"encoding/json"
	"net/http"
)

func HTTPErrorHandler(response http.ResponseWriter, body string, status int) {
	response.WriteHeader(status)
	errorResponse := models.HttpErrorResponse{
		Body: body,
	}
	jData, _ := json.Marshal(errorResponse)
	response.Write(jData)
}
