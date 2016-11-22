package dotquote


import (
	"sort"

	"testing"
)


func TestUnmarshalMapStringString(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Expected map[string][]string
	}{
		{
			Bytes: []byte(`"apple"="one" "banana"="two" "cherry"="three" "kiwi"."watermelon"="45" "one"."two"."three"=[] "zero"=["alpha", "beta"] "do"."it"="ONE" "do"."it"="TWO" "do"."it"=["THREE", "FOUR"]`),
			Expected: map[string][]string{
				`"apple"`:             []string{`one`},
				`"banana"`:            []string{`two`},
				`"cherry"`:            []string{`three`},
				`"kiwi"."watermelon"`: []string{`45`},
				`"one"."two"."three"`: []string{},
				`"zero"`:              []string{`alpha`, `beta`},
				`"do"."it"`:           []string{`ONE`, `TWO`, `THREE`, `FOUR`},
			},
		},
	}


	for testNumber, test := range tests {

		actualMap := map[string][]string{}

		if err := Unmarshal(test.Bytes, actualMap); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %v", testNumber, err, err)
			continue
		}

		if expected, actual := len(test.Expected), len(actualMap); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)

			expectedKeys := []string{}
			for k, _ := range test.Expected {
				expectedKeys = append(expectedKeys, k)
			}
			sort.Strings(expectedKeys)
			for i,k := range expectedKeys {
				t.Errorf("EXPECTED [%d] %s => %q", i, k, test.Expected[k])
			}

			actualKeys := []string{}
			for k, _ := range actualMap {
				actualKeys = append(actualKeys, k)
			}
			sort.Strings(actualKeys)
			for i,k := range actualKeys {
				t.Errorf("ACTUAL   [%d] %s => %q", i, k, actualMap[k])
			}

			continue
		}
		for key, expectedValues := range test.Expected {
			actualValues := actualMap[key]

			for valueNumber, expectedValue := range expectedValues {
				actualValue := actualValues[valueNumber]

				if expected, actual := expectedValue, actualValue; expected != actual {
					t.Errorf("For test #%d and key %q and value #%d, expected %q, but actually got %q.", testNumber, key, valueNumber, expected, actual)
					continue
				}
			}
		}
	}
}
