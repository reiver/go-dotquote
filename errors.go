package dotquote


import (
	"errors"
)


var (
	errInternalError = errors.New("Internal Error")
	errNilBytes      = errors.New("Nil Bytes")
	errNilReceiver   = errors.New("Nil Receiver")
	errNotEqualsSign = errors.New("Not Equals Sign")
	errNotUTF8       = errors.New("Not UTF-8")
	errPremature     = errors.New("Premature")
	errTooShort      = errors.New("Too Short")
)
