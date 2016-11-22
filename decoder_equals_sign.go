package dotquote


import (
	"unicode/utf8"
)


func (decoder *Decoder) eatEqualsSign() {

	if nil == decoder {
		panic(errNilReceiver)
	}

	b := decoder.Bytes
	if nil == b {
		decoder.err = errNilBytes
		return
	}

	lenb := len(b)

	if 0 >= lenb {
		decoder.err = errTooShort
		return
	}

	p := b[decoder.index:]
	if 0 >= len(p) {
		decoder.err = errTooShort
		return
	}

	r, size := utf8.DecodeRune(p)
	if utf8.RuneError == r {
		decoder.err = errNotUTF8
		return
	}

	if '=' != r {
		decoder.err = errNotEqualsSign
		return
	}

	decoder.index += size
	if lenb < decoder.index {
		decoder.err = errInternalError
		return
	}
}
