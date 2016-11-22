package dotquotedetect


import (
	"github.com/reiver/go-inquote"

	"io"
)


// DetectQuote looks for a dotquote quote in the dotquote data in a []byte, and returns
// the beginning index and the ending index.
//
// The returned beginning index and ending index are the values one would need
// to take a slice of the []byte, and get just that quote.
//
// For example:
//
//	b, e, err :+ dotquotedetect.DetectQuote(p)
//	if nil != err {
//		return err
//	}
//	
//	quote := p[b:e]
func DetectQuote(b []byte) (int, int, error) {
	if nil == b {
		return -1, -1, errNilBytes
	}

	lenb := len(b)

	if 2 > lenb {
		return -1, -1, errBadRequest
	}

	b0 := b[0]

	if '"' != b0 {
		return -1, -1, newNotQuoteComplainer(string(b))
	}

	const begin = 0
	end := 1 + begin

	p:= b[1:]

	for {
		if 0 >= len(p) {
			return -1, -1, errBadRequest
		}

		p0 := p[0]
		if '"' == p0 {
			end++
			break
		}

		_, n, err := inquote.DecodeRune(p)
		if nil != err {
			if io.EOF == err {
				break
			}
			return -1, -1, err
		}


		end += n

		p = p[n:]
	}


	return begin, end, nil
}
