package utils

import (
	"bytes"
	"io"
	"net/http"
)

// ReadBody reads the request body and returns it as a string.
// It also resets the request body so it can be read again later.
func ReadBody(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the body for further processing
	return string(body), nil
}

func ReadBodyAndUnmarshal(r *http.Request, v any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return UnmarshalJson(body, v)
}
