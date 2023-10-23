package errorhandler

import (
	"errors"
	"machship-go/util/logger"
	"net/http"
)

type ResponseError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

var (
	ErrGithubNotFound = errors.New("Github User not found")
)

func HandleError(log *logger.Logger, err error) *ResponseError {
	log.Error(err)

	switch err {
	case ErrGithubNotFound:
		return &ResponseError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	default:
		return &ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error occurred",
		}
	}
}
