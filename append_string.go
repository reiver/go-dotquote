package dotquote


func AppendString(p []byte, value string, name ...string) []byte {

	if 1 > len(name) {
		p = appendQuote(p, []byte(nil))
	} else {
		for i, s := range name {
			if 1 <= i {
				p = appendDot(p)
			}
			p = appendQuoteString(p, s)
		}
	}

	p = appendEqualsSign(p)

	p = appendQuoteString(p, value)

	return p
}
