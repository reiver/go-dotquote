package dotquote


import (
	"testing"
)


func TestDecoderValues(t *testing.T) {

	tests := []struct{
		Bytes []byte
		ExpectedValues [][]string
	}{
		{
			Bytes: []byte(``),
			ExpectedValues: [][]string{},
		},



		{
			Bytes: []byte(`"apple"="one"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three" "grape"="four"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three" "grape"="four" "kiwi"="five"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
				[]string{
					`"five"`,
				},
			},
		},



		{
			Bytes: []byte(` "apple"="one"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three" "grape"="four"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"="one" "banana"="two" "cherry"="three" "grape"="four" "kiwi"="five"`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
				[]string{
					`"five"`,
				},
			},
		},



		{
			Bytes: []byte(`"apple"=["one"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
			},
		},
		{
			Bytes: []byte(`"apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"] "kiwi"=["five"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
				[]string{
					`"five"`,
				},
			},
		},



		{
			Bytes: []byte(` "apple"=["one"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"=["one"] "banana"=["two"] "cherry"=["three"] "grape"=["four"] "kiwi"=["five"]`),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
				[]string{
					`"two"`,
				},
				[]string{
					`"three"`,
				},
				[]string{
					`"four"`,
				},
				[]string{
					`"five"`,
				},
			},
		},



		{
			Bytes: []byte(`"apple"=[]`),
			ExpectedValues: [][]string{
				[]string{},
			},
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[] "grape"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(`"apple"=[] "banana"=[] "cherry"=[] "grape"=[] "kiwi"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{},
			},
		},



		{
			Bytes: []byte(` "apple"=[]`),
			ExpectedValues: [][]string{
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[] "grape"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"=[] "banana"=[] "cherry"=[] "grape"=[] "kiwi"=[]`),
			ExpectedValues: [][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{},
			},
		},



		{
			Bytes: []byte(` "apple" = "one" `),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana" = "one two" `),
			ExpectedValues: [][]string{
				[]string{
					`"one two"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = "one two three" `),
			ExpectedValues: [][]string{
				[]string{
					`"one two three"`,
				},
			},
		},



		{
			Bytes: []byte(` "apple" = [ "one" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana" = [ "one two" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one two"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ "one two three" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one two three"`,
				},
			},
		},



		{
			Bytes: []byte(` "apple" = [ ] `),
			ExpectedValues: [][]string{
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"."banana" = [ ] `),
			ExpectedValues: [][]string{
				[]string{},
			},
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ ] `),
			ExpectedValues: [][]string{
				[]string{},
			},
		},



		{
			Bytes: []byte(` "apple" = [ "one" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana" = [ "one" , "two" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
					`"two"`,
				},
			},
		},
		{
			Bytes: []byte(` "apple"."banana"."cherry" = [ "one" , "two" , "three" ] `),
			ExpectedValues: [][]string{
				[]string{
					`"one"`,
					`"two"`,
					`"three"`,
				},
			},
		},



		{
			Bytes: []byte(` "apple"="1"  "banana"."cherry"=["2", "3"] "zero"=[] "zero"."one"=["ONE"] "zero"."one"."two"=["ONE","TWO"] "apple"."banana"."cherry" = [ "orange 1", "orange 2" ,"orange 3","orange 4"]`),
			ExpectedValues: [][]string{
				[]string{
					`"1"`,
				},
				[]string{
					`"2"`,
					`"3"`,
				},
				[]string{},
				[]string{
					`"ONE"`,
				},
				[]string{
					`"ONE"`,
					`"TWO"`,
				},
				[]string{
					`"orange 1"`,
					`"orange 2"`,
					`"orange 3"`,
					`"orange 4"`,
				},
			},
		},


		{
			Bytes: []byte(` "apple"."banana" = "this is it!" `),
			ExpectedValues: [][]string{
				[]string{
					`"this is it!"`,
				},
			},
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

		actualValues := [][]string{}
		for decoder.Next() {
			values := decoder.Values()

			v := []string{}
			for values.Next() {
				value := values.MustValueString()

				v = append(v, value)
			}

			actualValues = append(actualValues, v)
		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v; for |||||%s|||||", testNumber, err, err, string(test.Bytes))
			continue
		}

		if expected, actual := len(test.ExpectedValues), len(actualValues); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for: |||||%s|||||", testNumber, expected, actual, string(test.Bytes))
			continue
		}

		for valuesNumber, expectedSubValues := range test.ExpectedValues {
			actualSubValues := actualValues[valuesNumber]
			if expected, actual := len(expectedSubValues), len(actualSubValues); expected != actual {
				t.Errorf("For test #%d and values #%d, expected %d, but actually got %d.", testNumber, valuesNumber, expected, actual)
				t.Errorf("BYTES: |||||%s|||||", string(test.Bytes))
				t.Errorf("EXPECTED VALUES: %#v", test.ExpectedValues)
				t.Errorf("ACTUAL VALUES:   %#v", actualValues)
				continue
			}

			for valueNumber, expectedValue := range expectedSubValues {
				actualValue := actualSubValues[valueNumber]
				if expected, actual := expectedValue, actualValue; expected != actual {
					t.Errorf("For test #%d and values #%d and value #%d, expected %q, but actually got %q.", testNumber, valuesNumber, valueNumber, expected, actual)
					t.Errorf("BYTES: |||||%s|||||", string(test.Bytes))
					t.Errorf("EXPECTED VALUES: %#v", test.ExpectedValues)
					t.Errorf("ACTUAL VALUES:   %#v", actualValues)
					continue
				}
			}
		}
	}
}
