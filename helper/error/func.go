package errorConv

import (
	"errors"
	"strings"
)

func Conversion(err error) error {
	var errNew error
	if strings.Contains(err.Error(), "not found") {
		errNew = errors.New(ErrDBNotFound)
	}
	if strings.Contains(err.Error(), "invalid") ||
		strings.Contains(err.Error(), "missmatch") {
		errNew = errors.New(ErrInvalid)
	}
	if strings.Contains(err.Error(), "jwt") ||
		strings.Contains(err.Error(), "jwt token not found") {
		errNew = errors.New(ErrInvalid)
	}
	return errNew
}
