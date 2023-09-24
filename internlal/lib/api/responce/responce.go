package responce

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Responce struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Responce {
	return Responce{
		Status: StatusOK,
	}
}

func Error(msg string) Responce {
	return Responce{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidatorError(errs validator.ValidationErrors) Responce {

	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("filed %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("filed %s is a required URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("filed %s is a not valides", err.Field()))
		}
	}

	return Responce{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
