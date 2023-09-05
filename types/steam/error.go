package steam

import "errors"

var (
	ErrBadRequest   = errors.New("please confirm that all required parameters have been sent")
	ErrUnauthorized = errors.New("forbidden Access is denied, retrying will not help, please verify your key parameter")
	ErrForbidden    = errors.New("forbidden Access is forbidden, retrying will not help, please verify your key parameter")
	ErrNotFound     = errors.New("requested url not found")
	ErrNoMethod     = errors.New("method Not allowed")
	ErrExcessive    = errors.New("you are limited by the rate")
	ErrInternal     = errors.New("an unrecoverable error has occurred, please try again")
	ErrUnavailable  = errors.New("the server is temporarily unavailable or too busy to respond, please wait a moment and try again")
)
