package dotquote


import (
	"bytes"
	"math/rand"
	"time"

	"testing"
)


var (
	randomness = rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
)


func TestEatWhitespaceFailNilBytes(t *testing.T) {

	tests := []struct{
		Bytes []byte
	}{
		{
			Bytes: []byte(nil),
		},
	}


	for testNumber, test := range tests {

		decoder := Decoder{
			Bytes: test.Bytes,
		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		decoder.eatWhitespace()
		if expected, actual := decoder.Err(), errNilBytes; expected != actual {
			t.Errorf("For test #%d, expected (%T) %v, but actually got (%T) %v; for %q.", testNumber, expected, expected, actual, actual, string(test.Bytes))
			continue
		}

	}
}

func TestEatWhitespace(t *testing.T) {

	type TestType struct {
		Bytes []byte
		ExpectedIndex int
		ExpectedSlice []byte
	}

	appendTests := func(a []TestType, s string) []TestType {
		const whitespaceLimit = 30

		for i:=1; i < whitespaceLimit; i++ {
			p := bytes.Repeat([]byte{' '}, i)

			p = append(p, []byte(s)...)

			datum := struct{
				Bytes []byte
				ExpectedIndex int
				ExpectedSlice []byte
			}{
				Bytes: p,
				ExpectedIndex: i,
				ExpectedSlice: append([]byte(nil), s...),
			}

			a = append(a, datum)
		}

		whitespaceRunes := []rune{
			'\u0009', // horizontal tab
			'\u000A', // line feed
			'\u000B', // vertical tab
			'\u000C', // form feed
			'\u000D', // carriage return
			'\u0020', // space
			'\u0085', // next line
			'\u00A0', // no-break space
			'\u1680', // ogham space mark
			'\u180E', // mongolian vowel separator
			'\u2000', // en quad
			'\u2001', // em quad
			'\u2002', // en space
			'\u2003', // em space
			'\u2004', // three-per-em space
			'\u2005', // four-per-em space
			'\u2006', // six-per-em space
			'\u2007', // figure space
			'\u2008', // punctuation space
			'\u2009', // thin space
			'\u200A', // hair space
			'\u2028', // line separator
			'\u2029', // paragraph separator
			'\u202F', // narrow no-break space
			'\u205F', // medium mathematical space
			'\u3000', // ideographic space
		}


		repeatWhitespace := func(n int) []byte {
			p := []byte(nil)

			for i:=0; i<n; i++ {
				r := whitespaceRunes[randomness.Intn(len(whitespaceRunes))]

				p = append(p, string(r)...)
			}

			return p
		}

		for i:=1; i < whitespaceLimit; i++ {
			p := repeatWhitespace(i)

			expectedIndex := len(p)

			p = append(p, []byte(s)...)

			datum := struct{
				Bytes []byte
				ExpectedIndex int
				ExpectedSlice []byte
			}{
				Bytes: p,
				ExpectedIndex: expectedIndex,
				ExpectedSlice: append([]byte(nil), s...),
			}

			a = append(a, datum)
		}


		return a
	}

	tests := []TestType{
		{
			Bytes: []byte{},
			ExpectedIndex: 0,
			ExpectedSlice: []byte(``),
		},
		{
			Bytes: []byte(``),
			ExpectedIndex: 0,
			ExpectedSlice: []byte(``),
		},



		{
			Bytes: []byte(`apple`),
			ExpectedIndex: 0,
			ExpectedSlice: []byte(`apple`),
		},
		{
			Bytes: []byte(` apple`),
			ExpectedIndex: 1,
			ExpectedSlice: []byte(`apple`),
		},
		{
			Bytes: []byte(`  apple`),
			ExpectedIndex: 2,
			ExpectedSlice: []byte(`apple`),
		},



		{
			Bytes: []byte(`banana`),
			ExpectedIndex: 0,
			ExpectedSlice: []byte(`banana`),
		},
		{
			Bytes: []byte(` banana`),
			ExpectedIndex: 1,
			ExpectedSlice: []byte(`banana`),
		},
		{
			Bytes: []byte(`  banana`),
			ExpectedIndex: 2,
			ExpectedSlice: []byte(`banana`),
		},



		{
			Bytes: []byte(`cherry`),
			ExpectedIndex: 0,
			ExpectedSlice: []byte(`cherry`),
		},
		{
			Bytes: []byte(` cherry`),
			ExpectedIndex: 1,
			ExpectedSlice: []byte(`cherry`),
		},
		{
			Bytes: []byte(`  cherry`),
			ExpectedIndex: 2,
			ExpectedSlice: []byte(`cherry`),
		},



		{
			Bytes: []byte(`12345`),
			ExpectedIndex: 0,
			ExpectedSlice: []byte(`12345`),
		},
		{
			Bytes: []byte(` 12345`),
			ExpectedIndex: 1,
			ExpectedSlice: []byte(`12345`),
		},
		{
			Bytes: []byte(`  12345`),
			ExpectedIndex: 2,
			ExpectedSlice: []byte(`12345`),
		},
	}

	tests = appendTests(tests, "")
	tests = appendTests(tests, "apple")
	tests = appendTests(tests, "banana")
	tests = appendTests(tests, "cherry")
	tests = appendTests(tests, "apple banana cherry")



	for testNumber, test := range tests {

		decoder := Decoder{
			Bytes: test.Bytes,
		}

		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}
		if expected, actual := 0, decoder.index; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		decoder.eatWhitespace()
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v; for %q.", testNumber, err, err, string(test.Bytes))
			continue
		}

		if expected, actual := test.ExpectedIndex, decoder.index; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for len() = %d and %q.", testNumber, expected, actual, len(string(test.Bytes)), string(test.Bytes))
			continue
		}

		if expected, actual := test.ExpectedSlice, decoder.Bytes[decoder.index:]; !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, string(expected), string(actual))
			continue
		}

	}
}
