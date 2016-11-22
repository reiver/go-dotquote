package dotquote

import (
	"github.com/reiver/go-whitespace"

	"unicode/utf8"
)


func (decoder *Decoder) eatWhitespace() {

	if nil == decoder {
		panic(errNilReceiver)
	}

	logger := decoder.Logger
	if nil == logger {
		logger = internalDiscardLogger{}
	}

	b := decoder.Bytes
	if nil == b {
		decoder.err = errNilBytes
		return
	}

	lenb := len(b)

	if 0 >= lenb {
		return
	}

	index := decoder.index
	if lenb < index {
		return
	}
	p := b[index:]

	if 0 >= len(p) {
		return
	}

	//logger.Tracef("[EAT WHITESPACE] Will eat whitespace from: |||||%s|||||", string(p))

	for {
		r, size := utf8.DecodeRune(p)
		if utf8.RuneError == r {
			decoder.err = errNotUTF8
			return
		}
		//logger.Tracef("[EAT WHITESPACE] Have %d %q character.", r, string(r))

		if ! whitespace.IsWhitespace(r) {
			//logger.Trace("[EAT WHITESPACE] That character is NOT whitespace.")
			break
		}
		//logger.Trace("[EAT WHITESPACE] That character is whitespace.")

		//logger.Tracef("[EAT WHITESPACE] Old index is: %d", decoder.index)
		decoder.index += size
		index := decoder.index
		//logger.Tracef("[EAT WHITESPACE] New index is: %d", decoder.index)
		if lenb < index {
			decoder.err = errInternalError
			return
		}
		p = b[index:]
		if 0 >= len(p) {
			break
		}
	}

	//logger.Tracef("[EAT WHITESPACE] After eating whitespace: |||||%s|||||", string(p))
}
