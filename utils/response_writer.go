package utils

import (
	"net/http"
)

func Respond[T interface{}](w http.ResponseWriter, data T, errors []string, code int16) {
	response := NewApiResponse[interface{}](data, errors, code)
	w.WriteHeader(http.StatusNotFound)
	jsonData, err := response.ToJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}
