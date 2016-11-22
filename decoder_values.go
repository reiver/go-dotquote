package dotquote


func (decoder *Decoder) Values() *DecoderValues {
	decoderValues := DecoderValues{
		decoder:decoder,
	}

	return &decoderValues
}


type DecoderValues struct {
	decoder *Decoder

	hasBegun bool
	iterationCount int

	err error

	valueBegin int
	valueEnd   int
}

func (v DecoderValues) Err() error {
	return v.err
}

func (v *DecoderValues) Next() bool {
	if nil == v {
		panic(errNilReceiver)
	}

	logger := v.decoder.Logger
	if nil == logger {
		logger = internalDiscardLogger{}
	}

	v.hasBegun = true
	if nil != v.err {
		return false
	}
	v.iterationCount++
	logger.Debugf("[VALUES][NEXT] ITERATION COUNT #%d", v.iterationCount)

	values := v.decoder.values
	if nil == values {
		return false
	}

	i := v.iterationCount - 1
	if len(values) <= i {
		return false
	}
	currentValue := values[i]

	v.valueBegin = currentValue.valueBegin
	v.valueEnd   = currentValue.valueEnd

	return true
}



func (v DecoderValues) Value() (int, int, error) {
	if !v.hasBegun {
		return -1, -1, errPremature
	}

	return v.valueBegin, v.valueEnd, nil
}

func (v DecoderValues) ValueBytes() ([]byte, error) {
	begin, end, err := v.Value()
	if nil != err {
		return nil, err
	}

	b := v.decoder.Bytes
	if nil == b {
		return nil, errNilBytes
	}
	p := b[begin:end]

	return p, nil
}

func (v DecoderValues) ValueString() (string, error) {
	p, err := v.ValueBytes()
	if nil != err {
		return "", err
	}

	return string(p), nil
}



// MustValue is like Value, expect it panic()s on an error.
func (v DecoderValues) MustValue() (int, int) {
	begin, end, err := v.Value()
	if nil != err {
		panic(err)
	}

	return begin, end
}

// MustValueBytes is like ValueBytes, expect it panic()s on an error.
func (v DecoderValues) MustValueBytes() []byte {
	p, err := v.ValueBytes()
	if nil != err {
		panic(err)
	}

	return p
}

// MustValueString is like ValueString, expect it panic()s on an error.
func (v DecoderValues) MustValueString() string {
	s, err := v.ValueString()
	if nil != err {
		panic(err)
	}

	return s
}
