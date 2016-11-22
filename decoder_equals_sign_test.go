package dotquote


import (
	"testing"
)


func TestEatEqualsSignFailNilBytes(t *testing.T) {

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

		decoder.eatEqualsSign()
		if expected, actual := errNilBytes, decoder.Err(); expected != actual {
			t.Errorf("For test #%d, expected error (%T) %q, but actually got error (%T) %q; for %q.", testNumber, expected, expected, actual, actual, string(test.Bytes))
			continue
		}

	}
}

func TestEatEqualsSignFailNotEqualsSign(t *testing.T) {

	tests := []struct{
		Bytes []byte
	}{
		{
			Bytes: []byte(` =`),
		},
		{
			Bytes: []byte(` ="one two three"`),
		},
		{
			Bytes: []byte(` = "one two three"`),
		},

		{
			Bytes: []byte(`"apple"="one"`),
		},
		{
			Bytes: []byte(`"apple"."banana"="one two"`),
		},
		{
			Bytes: []byte(`"apple"."banana"."cherry"="one two three"`),
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

		decoder.eatEqualsSign()
		if expected, actual := errNotEqualsSign, decoder.Err(); expected != actual {
			t.Errorf("For test #%d, expected error (%T) %q, but actually got error (%T) %q; for %q.", testNumber, expected, expected, actual, actual, string(test.Bytes))
			continue
		}

	}

}


func TestEatEqualsSign(t *testing.T) {

	tests := []struct{
		Bytes []byte
		InitialIndex  int
		ExpectedIndex int
	}{
		{
			Bytes: []byte(`=`),
			InitialIndex:  0,
			ExpectedIndex: 1,
		},
		{
			Bytes: []byte(`="something"`),
			InitialIndex:  0,
			ExpectedIndex: 1,
		},
		{
			Bytes: []byte(`= "something"`),
			InitialIndex:  0,
			ExpectedIndex: 1,
		},
		{
			Bytes: []byte(`=	"something"`),
			InitialIndex:  0,
			ExpectedIndex: 1,
		},



		{
			Bytes: []byte(` =`),
			InitialIndex:  1,
			ExpectedIndex: 2,
		},
		{
			Bytes: []byte(` ="something"`),
			InitialIndex:  1,
			ExpectedIndex: 2,
		},
		{
			Bytes: []byte(` = "something"`),
			InitialIndex:  1,
			ExpectedIndex: 2,
		},
		{
			Bytes: []byte(` =	"something"`),
			InitialIndex:  1,
			ExpectedIndex: 2,
		},
	}


	for testNumber, test := range tests {

		decoder := Decoder{
			Bytes: test.Bytes,
		}
		decoder.index = test.InitialIndex

		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		decoder.eatEqualsSign()
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expected an error, but actually got one: (%T) %v; for |||||%s|||||", testNumber, err, err, string(test.Bytes))
			continue
		}

		if expected, actual := test.ExpectedIndex, decoder.index; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for len() = %d and %q.", testNumber, expected, actual, len(string(test.Bytes)), string(test.Bytes))
			continue
		}

	}
}
