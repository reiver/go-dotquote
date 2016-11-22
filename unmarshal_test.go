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
		{
			Bytes: []byte(` "app"."name"="myapi" "app"."deployenv"="PROD" "app"."build"."git"."branch"="master" "app"."build"."git"."revision"="2e3eef1dc1d1" "app"."build"."number"="5432" "app"."build"."when"="Mon Nov 21 10:05:14 PST 2016" "request"."client-address"="123.45.6.7:41325" "request"."method"="GET" "request"."uri"="/my/1.0/apple/banana/cherry" "request"."protocol"="HTTP/1.1" "request"."host"="localhost:8080" "request"."content-length"="0" "request"."transfer-encoding"=[] "response"."status-code"="200" "response"."content-length"="173" "trace"."begin-time"="2016-11-22 23:01:43.065121584 +0000 UTC" "trace"."end-time"="2016-11-22 23:01:43.065378496 +0000 UTC" "request"."header"."Upgrade-Insecure-Requests"=["1"] "request"."header"."User-Agent"=["Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0"] "request"."header"."Accept"=["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"] "request"."header"."Accept-Language"=["en-US,en;q=0.5"] "request"."header"."Accept-Encoding"=["gzip, deflate"] "request"."header"."Dnt"=["1"] "request"."header"."Connection"=["keep-alive"] `),
			Expected: map[string][]string{
				`"app"."name"`:                                   []string{`myapi`},
				`"app"."deployenv"`:                              []string{`PROD`},
				`"app"."build"."git"."branch"`:                   []string{`master`},
				`"app"."build"."git"."revision"`:                 []string{`2e3eef1dc1d1`},
				`"app"."build"."number"`:                         []string{`5432`},
				`"app"."build"."when"`:                           []string{`Mon Nov 21 10:05:14 PST 2016`},
				`"request"."client-address"`:                     []string{`123.45.6.7:41325`},
				`"request"."method"`:                             []string{`GET`},
				`"request"."uri"`:                                []string{`/my/1.0/apple/banana/cherry`},
				`"request"."protocol"`:                           []string{`HTTP/1.1`},
				`"request"."host"`:                               []string{`localhost:8080`},
				`"request"."content-length"`:                     []string{`0`},
				`"request"."transfer-encoding"`:                  []string{},
				`"response"."status-code"`:                       []string{`200`},
				`"response"."content-length"`:                    []string{`173`},
				`"trace"."begin-time"`:                           []string{`2016-11-22 23:01:43.065121584 +0000 UTC`},
				`"trace"."end-time"`:                             []string{`2016-11-22 23:01:43.065378496 +0000 UTC`},
				`"request"."header"."Upgrade-Insecure-Requests"`: []string{`1`},
				`"request"."header"."User-Agent"`:                []string{`Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:49.0) Gecko/20100101 Firefox/49.0`},
				`"request"."header"."Accept"`:                    []string{`text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`},
				`"request"."header"."Accept-Language"`:           []string{`en-US,en;q=0.5`},
				`"request"."header"."Accept-Encoding"`:           []string{`gzip, deflate`},
				`"request"."header"."Dnt"`:                       []string{`1`},
				`"request"."header"."Connection"`:                []string{`keep-alive`},
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
			actualValues, ok := actualMap[key]
			if !ok {
				t.Errorf("For test #%d and key |||||%s|||||, expected value(s) for key, but do not actually have that key.", testNumber, key)
				continue
			}

			if expected, actual := len(expectedValues), len(actualValues); expected != actual {
				t.Errorf("For test #%d and key |||||%s|||||, expected %d, but actually got %d.", testNumber, key, expected, actual)
				continue
			}
			for valueNumber, expectedValue := range expectedValues {
				actualValue := actualValues[valueNumber]

				if expected, actual := expectedValue, actualValue; expected != actual {
					t.Errorf("For test #%d and key |||||%s||||| and value #%d, expected %q, but actually got %q.", testNumber, key, valueNumber, expected, actual)
					continue
				}
			}
		}
	}
}
