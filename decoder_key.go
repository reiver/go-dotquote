package dotquote


// MustKey is like the Key method, expect it panic()s on an error.
func (decoder Decoder) MustKey() (int, int) {
	b, e, err := decoder.Key()
	if nil != err {
		panic(err)
	}

	return b, e
}

// Key returns the "begin index" and "end index" for the current dotquote key.
func (decoder Decoder) Key() (int, int, error) {
	if !decoder.hasBegun {
		return -1, -1, errPremature
	}

	return decoder.keyBegin, decoder.keyEnd, nil
}

// MustKeyBytes is like the KeyBytes method, expect it panic()s on an error.
func (decoder Decoder) MustKeyBytes() []byte {
	p, err := decoder.KeyBytes()
	if nil != err {
		panic(err)
	}

	return p
}

// KeyBytes returns a []byte to the current dotquote key.
//
// Note that this is a slice on the []byte originally given to this
// decoder, in its "Bytes" field.
//
// It is NOT a copy.
func (decoder Decoder) KeyBytes() ([]byte, error) {
	begin, end, err := decoder.Key()
	if nil != err {
		return nil, err
	}

	b := decoder.Bytes
	if nil == b {
		return nil, errNilBytes
	}
	p := b[begin:end]

	return p, nil
}

// MustKeyString is like the KeyString method, expect it panic()s on an error.
func (decoder Decoder) MustKeyString() string {
	s, err := decoder.KeyString()
	if nil != err {
		panic(err)
	}

	return s
}

// KeyString returns the current dotquote key, as a string.
func (decoder Decoder) KeyString() (string, error) {
	p, err := decoder.KeyBytes()
	if nil != err {
		return "", err
	}

	return string(p), nil
}
