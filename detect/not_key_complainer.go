package dotquotedetect


import (
	"fmt"
)


type notKeyComplainer interface {
	error
	NotKeyComplainer()
}

type internalNotKeyComplainer struct {
	s string
}


func newNotKeyComplainer(format string, args ...interface{}) error {
	err := internalNotKeyComplainer{
		s: fmt.Sprintf("Not Key: "+format, args...),
	}

	return err
}


func (err internalNotKeyComplainer) Error() string {
	return err.s
}

func (internalNotKeyComplainer) NotKeyComplainer() {
	// Nothing here.
}
