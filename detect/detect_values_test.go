package dotquotedetect


import (
	"testing"
)


func TestDetectValues(t *testing.T) {

	tests := []struct{
		Bytes          []byte
		ExpectedEndIndex int
		ExpectedValues []string
	}{
		{
			Bytes:  []byte(`[]`),
			ExpectedEndIndex: 2,
			ExpectedValues: []string{
				// Nothing here.
			},
		},



		{
			Bytes:  []byte(` []`),
			ExpectedEndIndex:  3,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(`[ ]`),
			ExpectedEndIndex:  3,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(`[] `),
			ExpectedEndIndex: 2,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(` [ ]`),
			ExpectedEndIndex:   4,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(` [] `),
			ExpectedEndIndex:  3,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(`[ ] `),
			ExpectedEndIndex:  3,
			ExpectedValues: []string{
				// Nothing here.
			},
		},
		{
			Bytes:  []byte(` [ ] `),
			ExpectedEndIndex:   4,
			ExpectedValues: []string{
				// Nothing here.
			},
		},



		{
			Bytes:  []byte(`   [   ]   `),
			ExpectedEndIndex:       8,
			ExpectedValues: []string{
				// Nothing here.
			},
		},



		{
			Bytes:  []byte(`"apple"`),
			ExpectedEndIndex:      7,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(` "apple"`),
			ExpectedEndIndex:       8,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`"apple" `),
			ExpectedEndIndex:      7,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(` "apple" `),
			ExpectedEndIndex:       8,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(`["apple"]`),
			ExpectedEndIndex:        9,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(` ["apple"]`),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`[ "apple"]`),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ]`),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`["apple"] `),
			ExpectedEndIndex:        9,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(` [ "apple"]`),
			ExpectedEndIndex:          11,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(` ["apple" ]`),
			ExpectedEndIndex:          11,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(` ["apple"] `),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(`[ "apple" ]`),
			ExpectedEndIndex:          11,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`[ "apple"] `),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(`["apple" ] `),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(` [ "apple" ]`),
			ExpectedEndIndex:           12,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(` [ "apple"] `),
			ExpectedEndIndex:          11,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(` ["apple" ] `),
			ExpectedEndIndex:          11,
			ExpectedValues: []string{
				`"apple"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ] `),
			ExpectedEndIndex:         10,
			ExpectedValues: []string{
				`"apple"`,
			},
		},



		{
			Bytes:  []byte(`["apple","banana"]`),
			ExpectedEndIndex:                 18,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
			},
		},
		{
			Bytes:  []byte(`["apple", "banana"]`),
			ExpectedEndIndex:                  19,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ,"banana"]`),
			ExpectedEndIndex:                  19,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
			},
		},



		{
			Bytes:  []byte(`["apple","banana","cherry"]`),
			ExpectedEndIndex:                          27,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple", "banana","cherry"]`),
			ExpectedEndIndex:                           28,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ,"banana","cherry"]`),
			ExpectedEndIndex:                           28,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple","banana", "cherry"]`),
			ExpectedEndIndex:                           28,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple","banana" ,"cherry"]`),
			ExpectedEndIndex:                           28,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple", "banana", "cherry"]`),
			ExpectedEndIndex:                            29,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ,"banana" ,"cherry"]`),
			ExpectedEndIndex:                            29,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple", "banana" ,"cherry"]`),
			ExpectedEndIndex:                            29,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},
		{
			Bytes:  []byte(`["apple" ,"banana", "cherry"]`),
			ExpectedEndIndex:                            29,
			ExpectedValues: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},



		{
			Bytes:  []byte(`["one","two","three"] "apple"."banana"."cherry"="something"`),
			ExpectedEndIndex:                    21,
			ExpectedValues: []string{
				`"one"`,
				`"two"`,
				`"three"`,
			},
		},



		{
			Bytes:  []byte(`  [  "one"  ,  "two"  ,  "three"  ]  "apple"  .  "banana"  .  "cherry"  =  "something"  `),
			ExpectedEndIndex:                                  35,
			ExpectedValues: []string{
				`"one"`,
				`"two"`,
				`"three"`,
			},
		},
	}


	TestLoop: for testNumber, test := range tests {

		iterator := DetectValues{
			Bytes: test.Bytes,
		}


		actualValues := []string{}
		for iterator.Next() {
			b, e, err := iterator.Detect()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
				continue TestLoop
			}

			value := test.Bytes[b:e]
			s := string(value)

			actualValues = append(actualValues, s)
		}
		if err := iterator.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			t.Errorf("EXPECTED: %#v", test.ExpectedValues)
			t.Errorf("ACTUAL:   %#v", actualValues)
			t.Errorf("ORIGINAL: %s", string(test.Bytes))
			continue
		}

		actualEndIndex := iterator.EndIndex()
		if expected, actual := test.ExpectedEndIndex, actualEndIndex; expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d; for |||||%s|||||", testNumber, expected, actual, (test.Bytes))
			continue
		}

		if expected, actual := len(test.ExpectedValues), len(actualValues); expected != actual {
			t.Errorf("For test #%d, expected %d values, but actually got %d values.", testNumber, expected, actual)
			t.Errorf("EXPECTED: %#v", test.ExpectedValues)
			t.Errorf("ACTUAL:   %#v", actualValues)
			t.Errorf("ORIGINAL: %s", string(test.Bytes))
			continue
		}
		for valueNumber, expectedValue := range test.ExpectedValues {
			actualValue := actualValues[valueNumber]

			if expected, actual := expectedValue, actualValue; expected != actual {
				t.Errorf("For test #%d and value #%d, expected {%s}, but actually got {%s}.", testNumber, valueNumber, expected, actual)
				continue
			}
		}
	}
}
