package dotquote


func AppendStrings(p []byte, value []string, name ...string) []byte {

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

	p = append(p, '[')
	for ii, datum := range value {
		if 0 < ii {
			p = append(p, ',')
		}
		p = appendQuoteString(p, datum)
	}
	p = append(p, ']')


	return p
}
