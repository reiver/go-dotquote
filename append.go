package dotquote


func Append(p []byte, value []byte, name ...[]byte) []byte {

	if 1 > len(name) {
		p = appendQuote(p, []byte(nil))
	} else {
		for i, b := range name {
			if 1 <= i {
				p = appendDot(p)
			}
			p = appendQuote(p, b)
		}
	}

	p = appendEqualsSign(p)

	p = appendQuote(p, value)

	return p
}
