package errors

import err "errors"

var (
	ErrRecordNotFound = err.New("record not found")
)
