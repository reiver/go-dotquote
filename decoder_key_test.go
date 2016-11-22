package dotquote


import (
	"testing"
)


func TestDecoderMustKeyString(t *testing.T) {

	tests := []struct{
		Bytes []byte
                ExpectedKeys []string
	}{
		{
			Bytes: []byte{},
			ExpectedKeys: []string{},
		},



		{
			Bytes: []byte(`"apple"="one"`),
			ExpectedKeys: []string{
				`"apple"`,
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two"`),
			ExpectedKeys: []string{
				`"apple"`,
				`"banana"`,
			},
		},
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three"`),
			ExpectedKeys: []string{
				`"apple"`,
				`"banana"`,
				`"cherry"`,
			},
		},



		{
			Bytes: []byte(`"one"="1" "two"."three"="2 3" "four"."five"."six"="4 5 6"`),
			ExpectedKeys: []string{
				`"one"`,
				`"two"."three"`,
				`"four"."five"."six"`,
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


		actualKeys := []string{}
		for decoder.Next() {
			key := decoder.MustKeyString()

			actualKeys = append(actualKeys, key)
		}
		if err := decoder.Err(); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.ExpectedKeys), len(actualKeys); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for keyNumber, expectedKey := range test.ExpectedKeys {
			actualKey := actualKeys[keyNumber]
			if expected, actual := expectedKey, actualKey; expected != actual {
				t.Errorf("For test #%d and key #%d, expected |||||%s|||||, but actually got |||||%s|||||.", testNumber, keyNumber, expected, actual)
				for i,v := range actualKeys {
					t.Errorf("\t key %d => |||||%s|||||", i, v)
				}
				t.Errorf("%s", string(test.Bytes))
				continue
			}
		}
	}
}
