package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

const ContentTypeJSON = "application/json"

type ErrorResponse struct {
	Error string `json:"error"`
}

func newErrorResponse(errorDetail string) *ErrorResponse {
	return &ErrorResponse{
		Error: errorDetail,
	}
}

type HTTPError struct {
	StatusCode int
	Err        error
}

func (he HTTPError) Error() string {
	return he.Err.Error()
}

func NewHTTPError(statusCode int, err error) HTTPError {
	return HTTPError{
		StatusCode: statusCode,
		Err:        err,
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", ContentTypeJSON)
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func MakeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			var httpErr HTTPError
			if errors.As(err, &httpErr) {
				WriteJson(w, httpErr.StatusCode, newErrorResponse(httpErr.Error()))
			} else {
				WriteJson(w, http.StatusInternalServerError, newErrorResponse(err.Error()))
			}
		}
	}
}
