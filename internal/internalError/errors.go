package internalerror

import (
	"errors"

	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

var ErrInternal error = errors.New("internal server error")

func ProcessError(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
