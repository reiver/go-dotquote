package dotquotedetect


import (
	"github.com/reiver/go-whitespace"

	"bytes"
	"unicode/utf8"
)


// DetectValues looks for dotquote values in the dotquote data in a []byte.
//
// Example usage:
//
//	iterator := dotquotedetect.DetectValues{
//		Bytes: p,
//	}
//	
//	for iterator.Next() {
//		b, e, err := iterator.Detect()
//		if nil != err {
//			return err
//		}
//		
//		value := p[b:e]
//
//		//@TODO: Do something with `value`.
//	}
//	if err := iterator.Err(); nil != err {
//		return err
//	}
type DetectValues struct {
	Bytes []byte

	err error

	detectedBegin int
	detectedEnd   int
	detectedErr   error
	unquoted      bytes.Buffer

	hasBegun bool
	isSingular bool

	p []byte
	index int
	r0 rune
}

func (v *DetectValues) pslice(n int) error {
	if nil == v {
		return errNilReceiver
	}

	v.p = v.p[n:]
	v.index += n

	if 0 >= len(v.p) {
		return nil
	}

	v.r0, _ = utf8.DecodeRune(v.p)
	if utf8.RuneError == v.r0 {
		return errNotUTF8
	}

	return nil
}

func (v *DetectValues) detect() bool {
	if nil == v {
		v.err = errNilReceiver
		return false
	}

	v.unquoted.Reset()
	v.detectedBegin, v.detectedEnd, v.detectedErr = DetectQuoteAndUnquote(&v.unquoted, v.p)
	if nil != v.detectedErr {
		v.err = v.detectedErr
		return false
	}

	n := v.detectedEnd

	v.detectedBegin += v.index
	v.detectedEnd   += v.index

	if err := v.pslice(n); nil != err {
		v.err = errInternalError
		return false
	}

	return true
}


// MustDetect is like Detect, but panic()s if there is an error.
func (v *DetectValues) MustDetect() (int, int) {
	begin, end, err := v.Detect()
	if nil != err {
		panic(err)
	}

	return begin, end
}


func (v *DetectValues) Detect() (int, int, error) {
	if nil == v {
		err := errNilReceiver
		v.err = err
		return -1, -1, err
	}

	return v.detectedBegin, v.detectedEnd, v.detectedErr
}


// MustUnquoteBytes is like UnquoteBytes, expect it panic()s if there is an error.
func (v DetectValues) MustUnquoteBytes() []byte {
	p, err := v.UnquoteBytes()
	if nil != err {
		panic(err)
	}

	return p
}

func (v DetectValues) UnquoteBytes() ([]byte, error) {
	if err := v.detectedErr; nil != err {
		return nil, err
	}

	p := append([]byte(nil), v.unquoted.Bytes()...)

	return p, nil
}


// MustUnquoteString is like UnquoteString, expect it panic()s if there is an error.
func (v DetectValues) MustUnquoteString() string {
	s, err := v.UnquoteString()
	if nil != err {
		panic(err)
	}

	return s
}

func (v DetectValues) UnquoteString() (string, error) {
	if err := v.detectedErr; nil != err {
		return "", err
	}

	s := v.unquoted.String()

	return s, nil
}

func (v DetectValues) Err() error {
	return v.err
}


func (v *DetectValues) Next() bool {
	if nil == v {
		v.err = errNilReceiver
		return false
	}

	if nil != v.err {
		return false
	}

	if !v.hasBegun {
		return v.init()
	}

	if v.isSingular {
		return false
	}


	if err := v.eatWhitespace(); nil != err {
		v.err = err
		return false
	}

	if ',' == v.r0 {
		if err := v.pslice(1); nil != err {
			v.err = errInternalError
			return false
		}

		if err := v.eatWhitespace(); nil != err {
			v.err = err
			return false
		}
	}


	if ']' == v.r0 {
		if err := v.pslice(1); nil != err {
			v.err = errInternalError
			return false
		}
		return false
	}


	if ! v.detect() {
		return false
	}
	if nil != v.Err() {
		return false
	}

	return true
}


func (v *DetectValues) init() bool {

	if nil == v.Bytes {
		v.err = errNilBytes
		return false
	}

	v.hasBegun = true

	v.p = v.Bytes

	v.r0, _ = utf8.DecodeRune(v.p)
	if utf8.RuneError == v.r0 {
		v.err = errNotUTF8
		return false
	}


	if err := v.eatWhitespace(); nil != err {
		v.err = err
		return false
	}

	switch v.r0 {
	case '"':
		v.isSingular = true

		if ! v.detect() {
			return false
		}
		if nil != v.Err() {
			return false
		}
		return true
	case '[':
		v.isSingular = false
		if err := v.pslice(1); nil != err {
			v.err = errInternalError
			return false
		}

		if err := v.eatWhitespace(); nil != err {
			v.err = err
			return false
		}

		if ']' == v.r0 {
			if err := v.pslice(1); nil != err {
				v.err = errInternalError
				return false
			}
			return false
		}


		if ! v.detect() {
			return false
		}
		if nil != v.Err() {
			return false
		}
		return true
	default:
		v.err = errNotValues
		return false
	}

	return true
}


func (v DetectValues) EndIndex() int {
	return v.index
}

func (v *DetectValues) eatWhitespace() error {

	for whitespace.IsWhitespace(v.r0) {
		if err := v.pslice(1); nil != err { // <--- we expect to be in the ASCII range, so 1 byte is correct.
			v.err = errInternalError
			return v.err
		}
	}

	return nil
}
