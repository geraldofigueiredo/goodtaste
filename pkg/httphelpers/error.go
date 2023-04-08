package httphelpers

import (
	"encoding/json"
)

type HttpError struct {
	message string   `json:"message"`
	errors  []string `json:"errors"`
}

func NewHttpError(message string, errors ...error) HttpError {
	he := HttpError{
		message: message,
	}

	for _, err := range errors {
		he.errors = append(he.errors, err.Error())
	}

	return he
}

func (he *HttpError) Error() string {
	b, err := json.Marshal(he)
	if err != nil {
		return ""
	}

	return string(b)
}
