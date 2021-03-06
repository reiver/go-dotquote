package dotquotedetect


import (
	"github.com/reiver/go-whitespace"

	"unicode/utf8"
)


var (
	errNotKeyLenLessThanTwo = newNotKeyComplainer("length less that 2")
)


// MustDetectKey is like DetectKey, but panic()s if there is an error.
func MustDetectKey(b []byte) (int, int) {
	begin, end, err := DetectKey(b)
	if nil != err {
		panic(err)
	}

	return begin, end
}


// DetectKey looks for a dotquote key in the dotquote data in a []byte, and returns
// the beginning index and the ending index.
//
// The returned beginning index and ending index are the values one would need
// to take a slice of the []byte, and get just that key.
//
// For example:
//
//	b, e, err := dotquotedetect.DetectKey(p)
//	if nil != err {
//		return err
//	}
//	
//	key := p[b:e]
func DetectKey(b []byte) (int, int, error) {
	if nil == b {
		return -1, -1, errNilBytes
	}

	lenb := len(b)

	if 0 >= lenb {
		return -1, -1, errEmptyRequest
	}

	if 2 > lenb {
		return -1, -1, errNotKeyLenLessThanTwo
	}

	{
		r0, _ := utf8.DecodeRune(b)
		if utf8.RuneError == r0 {
			return -1, -1, errNotUTF8
		}

		if '"' != r0 {
			return -1, -1, newNotKeyComplainer("first rune not a double quote (i.e,. \"): %q %d", string(r0), r0)
		}
	}


	p := b

	begin, end, err := DetectQuote(p)
	if nil != err {
		return -1, -1, newNotKeyComplainer("could not detect quoted part of key: %s", err)
	}
	p = p[end:]

	for {
		if 0 >= len(p) {
			return -1, -1, errEmptyRequest
		}

		r0, size := utf8.DecodeRune(p)
		if utf8.RuneError == r0 {
			return -1, -1, errNotUTF8
		}
		if '=' == r0 {
			break
		}
		if whitespace.IsWhitespace(r0) {
			break
		}

		if '.' != r0 {
			return -1, -1, newNotKeyComplainer("not a dot at index %d of |||||%s|||||: %q (%d)", end, string(b), string(r0), r0)
		}

		p = p[size:]
		end += size

		_, finish, err := DetectQuote(p)
		if nil != err {
			return -1, -1, newNotKeyComplainer("could not detect quoted part of key: %s", err)
		}

		end += finish

                p = p[finish:]
	}

	return begin, end, nil
}
