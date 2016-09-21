package dotquote


import (
	"github.com/reiver/go-inquote"

	"unicode/utf8"
)


func appendQuoteString(p []byte, s string) []byte {
	p =  append(p, byte('"'))

	for 0 < len(s) {
		r, size := utf8.DecodeRuneInString(s)

		p = inquote.AppendRune(p, r)

		s = s[size:]
	}

	p =  append(p, byte('"'))

	return p
}
