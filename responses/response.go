package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OK(responseWriter http.ResponseWriter, statusCode int, data interface{}) {
	responseWriter.WriteHeader(statusCode)
	error := json.NewEncoder(responseWriter).Encode(data)
	if error != nil {
		fmt.Fprintf(responseWriter, "%s", error.Error())
	}
}

func SUCCESS(responseWriter http.ResponseWriter, statusCode int, data interface{}) {
	responseWriter.WriteHeader(statusCode)
	error := json.NewEncoder(responseWriter).Encode(Response{
		Data:    data,
		Message: "Success",
	})
	if error != nil {
		fmt.Fprintf(responseWriter, "%s", error.Error())
	}
}

func ERROR(responseWriter http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		OK(responseWriter, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	OK(responseWriter, http.StatusBadRequest, nil)
}
