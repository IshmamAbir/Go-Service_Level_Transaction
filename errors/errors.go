package errors

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")

	ErrNotFound = errors.New("Your requested Item is not found")

	ErrConflict = errors.New("Your Item already exist")

	ErrBadParamInput = errors.New("Given Param is not valid")

	ErrBadReusePassword = errors.New("Your Password already used")

	ErrPasswordMismatched = errors.New("Password matching failed")

	ErrFileSizeLimit = errors.New("Limit File Size capacity")

	ErrSingnin = errors.New("username or password is not correct")

	ErrNotFoundCookie = errors.New("session not found")

	ErrAcountLocked = errors.New("your account locked")
)
