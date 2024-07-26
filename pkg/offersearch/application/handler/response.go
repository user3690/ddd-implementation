package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonErrorResponse struct {
	Errors []JsonError `json:"errors"`
}

type JsonError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func writeError(
	writer http.ResponseWriter,
	status int,
	title string,
	responseErr error,
) {
	newJsonError := JsonErrorResponse{
		Errors: []JsonError{
			{
				Status: status,
				Title:  title,
				Detail: responseErr.Error(),
			},
		},
	}

	data, err := json.Marshal(newJsonError)
	if err != nil {
		fmt.Println(err)

		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	_, err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	return
}
