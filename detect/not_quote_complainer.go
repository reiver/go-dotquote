package dotquotedetect


import (
	"fmt"
)


type notQuoteComplainer interface {
	error
	NotQuoteComplainer()
}

type internalNotQuoteComplainer struct {
	s string
}


func newNotQuoteComplainer(s string) error {
	err := internalNotQuoteComplainer{
		s: fmt.Sprintf("Not Quote: %q", s),
	}

	return err
}


func (err internalNotQuoteComplainer) Error() string {
	return err.s
}

func (internalNotQuoteComplainer) NotQuoteComplainer() {
	// Nothing here.
}
