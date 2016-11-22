package dotquote


import (
	"testing"
)


func TestDecoderNext(t *testing.T) {

	tests := []struct{
		Bytes []byte
		ExpectedCount int
	}{
		{
			Bytes: []byte(``),
			ExpectedCount: 0,
		},



		{
			Bytes: []byte(`"apple"="one"`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two"`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three" "grape"="four"`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three" "grape"="four" "kiwi"="five"`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(` "apple"="one"`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two"`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three" "grape"="four"`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three" "grape"="four" "kiwi"="five"`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(`"apple"=["one"]`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"]`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"]`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"]`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"] "kiwi"=["five"]`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(` "apple"=["one"]`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"]`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"]`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"]`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"] "kiwi"=["five"]`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(`"apple"=[]`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[]`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[]`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[] "grape"=[]`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[] "grape"=[] "kiwi"=[]`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(` "apple"=[]`),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[]`),
			ExpectedCount: 2,
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[]`),
			ExpectedCount: 3,
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[] "grape"=[]`),
			ExpectedCount: 4,
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[] "grape"=[] "kiwi"=[]`),
			ExpectedCount: 5,
		},



		{
			Bytes: []byte(` "apple" = "one" `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana" = "one two" `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = "one two three" `),
			ExpectedCount: 1,
		},



		{
			Bytes: []byte(` "apple" = [ "one" ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana" = [ "one two" ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ "one two three" ] `),
			ExpectedCount: 1,
		},



		{
			Bytes: []byte(` "apple" = [ ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana" = [ ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ ] `),
			ExpectedCount: 1,
		},



		{
			Bytes: []byte(` "apple" = [ "one" ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana" = [ "one" , "two" ] `),
			ExpectedCount: 1,
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ "one" , "two" , "three" ] `),
			ExpectedCount: 1,
		},



		{
			Bytes: []byte(` "apple"="1"  "banana"."cherry"=["2", "3"] "zero"=[] "zero"."one"=["ONE"] "zero"."one"."two"=["ONE","TWO"] "apple"."banana"."cherry" = [ "orange 1", "organe 2" ,"orange 3","orange 4"]`),
			ExpectedCount: 6,
		},


		{
			Bytes: []byte(` "apple"."banana" = "this is it!" `),
			ExpectedCount: 1,
		},
	}


	for testNumber, test := range tests {

		decoder := Decoder{
			Bytes: test.Bytes,
			Logger: otherLoggerAdaptor{OtherLogger: t},
		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		actualCount := 0
		for decoder.Next() {
			actualCount++
		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v; for |||||%s|||||", testNumber, err, err, string(test.Bytes))
			continue
		}

		if expected, actual := test.ExpectedCount, actualCount; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for: |||||%s|||||", testNumber, expected, actual, string(test.Bytes))
			continue
		}
	}
}
