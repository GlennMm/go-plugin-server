package utils

import (
	"encoding/json"
)

type ApiResponse[T any] struct {
	data       T        `json:"data"`
	errors     []string `json:"errors"`
	statusCode int16    `json:"statusCode"`
}

func NewApiResponse[T any](data T, errors []string, statusCode int16) *ApiResponse[T] {
	return &ApiResponse[T]{
		data,
		errors,
		statusCode,
	}
}

func (r ApiResponse[T]) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r.data)
	if err != nil {
		return nil, err
	}

	// errorsBytes, err := json.Marshal(r.errors)
	// if err != nil {
	// 	return nil, err
	// }
	//
	return json.Marshal(struct {
		Data       json.RawMessage `json:"data"`
		Errors     []string        `json:"errors"`
		StatusCode int16           `json:"statusCode"`
	}{
		Data:       data,
		Errors:     r.errors,
		StatusCode: r.statusCode,
	})
}

func (r *ApiResponse[T]) FromJSON(data []byte) error {
	var apiResponse struct {
		Data       json.RawMessage `json:"data"`
		Errors     []string        `json:"errors"`
		StatusCode int16           `json:"statusCode"`
	}

	err := json.Unmarshal(data, &apiResponse)
	if err != nil {
		return err
	}

	err = json.Unmarshal(apiResponse.Data, &r.data)
	if err != nil {
		return err
	}

	r.errors = apiResponse.Errors
	r.statusCode = apiResponse.StatusCode

	return nil
}
