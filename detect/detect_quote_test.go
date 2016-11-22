package dotquotedetect


import (
	"testing"
)


func TestDetectQuoteFailNilBytes(t *testing.T) {

	tests := []struct{
		Bytes []byte
	}{
		{
			Bytes: nil,
		},
		{
			Bytes: []byte(nil),
		},
	}


	for testNumber, test := range tests {

		_, _, err := DetectQuote( []byte(test.Bytes) )
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %v", testNumber, err)
			continue
		}
		if expected, actual := errNilBytes, err; expected != actual {
			t.Errorf("For test #%d, expected (%T) %q, but actually got (%T) %q.", testNumber, expected, expected, actual, actual)
			continue
		}

	}
}

func TestDetectQuoteFailBadRequest(t *testing.T) {

	tests := []struct{
		Bytes []byte
	}{
		{
			Bytes: []byte{},
		},



		{
			Bytes: []byte(`"`),
		},
		{
			Bytes: []byte(`"a`),
		},
		{
			Bytes: []byte(`"ab`),
		},
		{
			Bytes: []byte(`"abc`),
		},
	}


	for testNumber, test := range tests {

		_, _, err := DetectQuote( []byte(test.Bytes) )
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %v", testNumber, err)
			continue
		}
		if expected, actual := errBadRequest, err; expected != actual {
			t.Errorf("For test #%d, expected (%T) %q, but actually got (%T) %q.", testNumber, expected, expected, actual, actual)
			continue
		}

	}
}

func TestDetectQuoteFailNotQuote(t *testing.T) {

	tests := []struct{
		Bytes []byte
	}{
		{
			Bytes: []byte(` "apple"`),
		},
		{
			Bytes: []byte(`  "banana"`),
		},
		{
			Bytes: []byte(`   "cherry"`),
		},



		{
			Bytes: []byte("\t\"apple\""),
		},
		{
			Bytes: []byte("\t\t\"banana\""),
		},
		{
			Bytes: []byte("\t\t\t\"cherry\""),
		},



		{
			Bytes: []byte("\v\"apple\""),
		},
		{
			Bytes: []byte("\v\v\"banana\""),
		},
		{
			Bytes: []byte("\v\v\v\"cherry\""),
		},
	}


	for testNumber, test := range tests {

		_, _, err := DetectQuote( []byte(test.Bytes) )
		if nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually get one: %v", testNumber, err)
			continue
		}
		if _, ok := err.(notQuoteComplainer); !ok {
			t.Errorf("For test #%d, expected NotQuoteComplainer error, but actually got (%T) %q.", testNumber, err, err)
			continue
		}
	}
}

func TestDetectQuote(t *testing.T) {

	tests := []struct{
		Bytes       []byte
		ExpectedBegin int
		ExpectedEnd   int
		Expected      string
	}{

		{
			Bytes: []byte(`"apple"="one"`),
			ExpectedBegin: 0,
			ExpectedEnd:          7,
			Expected:     `"apple"`,
		},
		{
			Bytes: []byte(`"banana"="two"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"banana"`,
		},
		{
			Bytes: []byte(`"cherry"="three"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"cherry"`,
		},


		{
			Bytes: []byte(`"apple"="one" "banana"="two"`),
			ExpectedBegin: 0,
			ExpectedEnd:          7,
			Expected:     `"apple"`,
		},
		{
			Bytes: []byte(`"banana"="two" "cherry"="three"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"banana"`,
		},
		{
			Bytes: []byte(`"cherry"="three" "apple"="one"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"cherry"`,
		},


		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedBegin: 0,
			ExpectedEnd:          7,
			Expected:     `"apple"`,
		},
		{
			Bytes: []byte(`"banana"="two" "cherry"="three" "apple"="one"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"banana"`,
		},
		{
			Bytes: []byte(`"cherry"="three" "apple"="one" "banana"="two"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"cherry"`,
		},



		{
			Bytes: []byte(`"apple"."banana"."cherry"="one two three"`),
			ExpectedBegin: 0,
			ExpectedEnd:          7,
			Expected:     `"apple"`,
		},



		{
			Bytes: []byte(`""="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:     2,
			Expected:     `""`,
		},
		{
			Bytes: []byte(`"1"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:      3,
			Expected:     `"1"`,
		},
		{
			Bytes: []byte(`"12"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:       4,
			Expected:     `"12"`,
		},
		{
			Bytes: []byte(`"123"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:        5,
			Expected:     `"123"`,
		},
		{
			Bytes: []byte(`"1234"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:         6,
			Expected:     `"1234"`,
		},
		{
			Bytes: []byte(`"12345"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:          7,
			Expected:     `"12345"`,
		},
		{
			Bytes: []byte(`"123456"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:           8,
			Expected:     `"123456"`,
		},
		{
			Bytes: []byte(`"1234567"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:            9,
			Expected:     `"1234567"`,
		},
		{
			Bytes: []byte(`"12345678"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:             10,
			Expected:     `"12345678"`,
		},
		{
			Bytes: []byte(`"123456789"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:              11,
			Expected:     `"123456789"`,
		},
		{
			Bytes: []byte(`"123456789a"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:               12,
			Expected:     `"123456789a"`,
		},
		{
			Bytes: []byte(`"123456789ab"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                13,
			Expected:     `"123456789ab"`,
		},
		{
			Bytes: []byte(`"123456789abc"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                 14,
			Expected:     `"123456789abc"`,
		},
		{
			Bytes: []byte(`"123456789abcd"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                  15,
			Expected:     `"123456789abcd"`,
		},
		{
			Bytes: []byte(`"123456789abcde"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                   16,
			Expected:     `"123456789abcde"`,
		},
		{
			Bytes: []byte(`"123456789abcdef"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                    17,
			Expected:     `"123456789abcdef"`,
		},
		{
			Bytes: []byte(`"123456789abcdefg"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                     18,
			Expected:     `"123456789abcdefg"`,
		},
		{
			Bytes: []byte(`"123456789abcdefgh"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                      19,
			Expected:     `"123456789abcdefgh"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghi"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                       20,
			Expected:     `"123456789abcdefghi"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghij"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                        21,
			Expected:     `"123456789abcdefghij"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijk"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                         22,
			Expected:     `"123456789abcdefghijk"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijkl"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                          23,
			Expected:     `"123456789abcdefghijkl"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklm"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                           24,
			Expected:     `"123456789abcdefghijklm"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmn"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                            25,
			Expected:     `"123456789abcdefghijklmn"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmno"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                             26,
			Expected:     `"123456789abcdefghijklmno"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnop"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                              27,
			Expected:     `"123456789abcdefghijklmnop"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopq"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                               28,
			Expected:     `"123456789abcdefghijklmnopq"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqr"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                29,
			Expected:     `"123456789abcdefghijklmnopqr"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrs"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                 30,
			Expected:     `"123456789abcdefghijklmnopqrs"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrst"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                  31,
			Expected:     `"123456789abcdefghijklmnopqrst"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstu"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                   32,
			Expected:     `"123456789abcdefghijklmnopqrstu"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstuv"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                    33,
			Expected:     `"123456789abcdefghijklmnopqrstuv"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstuvw"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                     34,
			Expected:     `"123456789abcdefghijklmnopqrstuvw"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstuvwx"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                      35,
			Expected:     `"123456789abcdefghijklmnopqrstuvwx"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstuvwxy"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                       36,
			Expected:     `"123456789abcdefghijklmnopqrstuvwxy"`,
		},
		{
			Bytes: []byte(`"123456789abcdefghijklmnopqrstuvwxyz"="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                                        37,
			Expected:     `"123456789abcdefghijklmnopqrstuvwxyz"`,
		},



		{
			Bytes: []byte(`"\""="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:       4,
			Expected:     `"\""`,
		},



		{
			Bytes: []byte(`"She said, \"hello\"."="something is here"`),
			ExpectedBegin: 0,
			ExpectedEnd:                         22,
			Expected:     `"She said, \"hello\"."`,
		},
	}


	for testNumber, test := range tests {


		actualBegin, actualEnd, err := DetectQuote( []byte(test.Bytes) )
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := test.ExpectedBegin, actualBegin; expected != actual {
			t.Errorf("For test #%d, expected \"begin\" index to be %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
		if expected, actual := test.ExpectedEnd, actualEnd; expected != actual {
			t.Errorf("For test #%d, expected \"end\" index to be %d, but actually got %d.", testNumber, expected, actual)
			continue
		}
		if expected, actual := test.Expected, string( test.Bytes[actualBegin:actualEnd] ); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
