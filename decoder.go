package dotquote


import (
	"github.com/reiver/go-dotquote/detect"

	"fmt"
)


type Decoder struct {
	Bytes []byte
	Logger Logger

	iterationCount int

	index int
	err error

	hasBegun bool

	keyBegin int
	keyEnd   int

	values []struct{
		valueBegin int
		valueEnd   int
	}
}


func (decoder Decoder) Err() error {
	return decoder.err
}


func (decoder *Decoder) Next() bool {
	if nil == decoder {
		panic(errNilReceiver)
	}

	logger := decoder.Logger
	if nil == logger {
		logger = internalDiscardLogger{}
	}

	decoder.hasBegun = true
	if nil != decoder.err {
		return false
	}
	decoder.iterationCount++
	logger.Debugf("[NEXT] ITERATION COUNT #%d", decoder.iterationCount)


	b := decoder.Bytes
	if nil == b {
		decoder.err = errNilBytes
		return false
	}
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Bytes: |||||%s|||||", string(b))
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	decoder.eatWhitespace()
	if nil != decoder.err {
		return false
	}
	//logger.Trace("[NEXT] After eating whitespace.")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	{
		index := decoder.index

		p := b[index:]

		if 0 >= len(p) {
			return false
		}
	}



	{
		index := decoder.index
		if 0 > index {
			decoder.err = errInternalError
			return false
		}

		p := b[index:]

		begin, end, err := dotquotedetect.DetectKey(p)
		if nil != err {
			decoder.err = fmt.Errorf("Problem detecting dotquote key #%d: %q; when looking at index %d of |||||%s||||| ==> |||||%s|||||", decoder.iterationCount, err, index, string(b), string(p))
			return false
		}

		decoder.keyBegin = index + begin
		decoder.keyEnd   = index + end

		decoder.index += end
	}
	//logger.Trace("[NEXT] Key Parsed")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	decoder.eatWhitespace()
	if nil != decoder.err {
		return false
	}
	//logger.Trace("[NEXT] After eating whitespace.")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	decoder.eatEqualsSign()
	if err := decoder.err; nil != err {
		index := decoder.index

		decoder.err = fmt.Errorf("Problem detecting equals sign (=) #%d: %q; when looking at index %d of |||||%s|||||", decoder.iterationCount, err, index, string(b))
		return false
	}
	//logger.Trace("[NEXT] After eating equals sign (=)")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	decoder.eatWhitespace()
	if nil != decoder.err {
		return false
	}
	//logger.Trace("[NEXT] After eating whitespace.")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))



	{
		index := decoder.index

		detectValues := dotquotedetect.DetectValues{
			Bytes: b[index:],
		}
		for detectValues.Next() {

			begin, end, err := detectValues.Detect()
			if nil != err {
				decoder.err = err
				return false
			}

			value := struct{
				valueBegin int
				valueEnd   int
			}{
				valueBegin: index+begin,
				valueEnd:   index+end,
			}

			decoder.values = append(decoder.values, value)

		}
		if err := detectValues.Err(); nil != err {
			decoder.err = err
			return false
		}
		decoder.index += detectValues.EndIndex()
	}
	//logger.Trace("[NEXT] Values Parsed")
	//logger.Tracef("[NEXT] index: %d", decoder.index)
	//logger.Tracef("[NEXT] Substring: |||||%s|||||", string(b[decoder.index:]))


	return true
}
