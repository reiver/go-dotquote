package dotquotedetect


import (
	"github.com/reiver/go-inquote"
	"github.com/reiver/go-oi"

	"bytes"
	"io"
	"io/ioutil"
	"unicode/utf8"
)


// MustDetectQuote is like DetectQuote , but panic()s if there is an error.
func MustDetectQuote(b []byte) (int, int) {
	begin, end, err := DetectQuote(b)
	if nil != err {
		panic(err)
	}

	return begin, end
}


// DetectQuote looks for a dotquote quote in the dotquote data in a []byte, and returns
// the beginning index and the ending index.
//
// The returned beginning index and ending index are the values one would need
// to take a slice of the []byte, and get just that quote.
//
// For example:
//
//	b, e, err := dotquotedetect.DetectQuote(p)
//	if nil != err {
//		return err
//	}
//	
//	quote := p[b:e]
func DetectQuote(b []byte) (int, int, error) {
	return DetectQuoteAndUnquote(ioutil.Discard, b)
}

// DetectQuoteAndUnquote is like DetectQuote, expect it also writes the quoted version of the quote
// to the io.Writer.
func DetectQuoteAndUnquote(w io.Writer, b []byte) (int, int, error) {
	if nil == b {
		return -1, -1, errNilBytes
	}

	lenb := len(b)

	if 2 > lenb {
		return -1, -1, errBadRequest
	}

	r0, size := utf8.DecodeRune(b)
	if utf8.RuneError == r0 {
		return -1, -1, errNotUTF8
	}


	if '"' != r0 {
		return -1, -1, newNotQuoteComplainer(string(b))
	}

	const begin = 0
	end := size + begin

	p:= b[size:]

	for {
		if 0 >= len(p) {
			return -1, -1, errBadRequest
		}

		r0, size := utf8.DecodeRune(p)
		if utf8.RuneError == r0 {
			return -1, -1, errNotUTF8
		}

		if '"' == r0 {
			end += size
			break
		}

		r, n, err := inquote.DecodeRune(p)
		if nil != err {
			if io.EOF == err {
				break
			}
			return -1, -1, err
		}

		{
			var buffer bytes.Buffer

			_, err := buffer.WriteRune(r)
			if nil != err {
				return -1, -1, err
			}

			_, err = oi.LongWrite(w, buffer.Bytes())
			if nil != err {
				return -1, -1, err
			}
		}


		end += n

		p = p[n:]
	}


	return begin, end, nil
}
