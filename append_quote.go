package dotquote


import (
	"github.com/reiver/go-inquote"

	"unicode/utf8"
)


func appendQuote(p []byte, b []byte) []byte {
	p =  append(p, byte('"'))

	for 0 < len(b) {
		r, size := utf8.DecodeRune(b)

		p = inquote.AppendRune(p, r)

		b = b[size:]
	}

	p =  append(p, byte('"'))

	return p
}
