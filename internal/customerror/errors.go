package customerror

import "errors"

var (
	ErrEmptyResult         = errors.New("empty result")
	ErrNotFound            = errors.New("not found")
	ErrDatabase            = errors.New("database error")
	ErrParse               = errors.New("parse error")
	ErrInternalServerError = errors.New("internal server error")
)
