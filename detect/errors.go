package dotquotedetect


import (
	"errors"
)


var (
	errBadRequest    = errors.New("Bad Request")
	errEmptyRequest  = errors.New("Empty Request")
	errInternalError = errors.New("Internal Error")
	errNilBytes      = errors.New("Nil Bytes")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotUTF8       = errors.New("Not UTF-8")
	errNotValues     = errors.New("Not Values")
)
