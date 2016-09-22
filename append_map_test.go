package dotquote


import (
	"testing"
)


func TestAppendMap(t *testing.T) {

	tests := []struct{
		P          []byte
		NamePrefix []string
		Map     map[string]interface{}
		Expected string

	}{
		{
			P: []byte(nil),
			NamePrefix: []string(nil),
			Map: map[string]interface{}(nil),
			Expected: ``,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}(nil),
			Expected: ``,
		},
		{
			P: []byte(nil),
			NamePrefix: []string(nil),
			Map: map[string]interface{}{},
			Expected: ``,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{},
			Expected: ``,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
			},
			Expected: `"one"="apple"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
			},
			Expected: `"one"="apple" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
			},
			Expected: `"one"="apple" "three"="cherry" "two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
			},
			Expected: `"hello"."one"="apple"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
			},
			Expected: `"hello"."one"="apple" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
			},
			Expected: `"hello"."one"="apple" "hello"."three"="cherry" "hello"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
			},
			Expected: `"hello"."world"."one"="apple"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
			},
			Expected: `"hello"."world"."one"="apple" "hello"."world"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
			},
			Expected: `"hello"."world"."one"="apple" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
			},
			Expected: `"four"."first"="grape" "four"."second"="kiwi" "four"."third"="watermelon" "one"="apple" "three"="cherry" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
			},
			Expected: `"hello"."four"."first"="grape" "hello"."four"."second"="kiwi" "hello"."four"."third"="watermelon" "hello"."one"="apple" "hello"."three"="cherry" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
			},
			Expected: `"hello"."world"."four"."first"="grape" "hello"."world"."four"."second"="kiwi" "hello"."world"."four"."third"="watermelon" "hello"."world"."one"="apple" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
			},
			Expected: `"five"="true" "four"."first"="grape" "four"."second"="kiwi" "four"."third"="watermelon" "one"="apple" "three"="cherry" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
			},
			Expected: `"hello"."five"="true" "hello"."four"."first"="grape" "hello"."four"."second"="kiwi" "hello"."four"."third"="watermelon" "hello"."one"="apple" "hello"."three"="cherry" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
			},
			Expected: `"hello"."world"."five"="true" "hello"."world"."four"."first"="grape" "hello"."world"."four"."second"="kiwi" "hello"."world"."four"."third"="watermelon" "hello"."world"."one"="apple" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
			},
			Expected: `"five"="true" "four"."first"="grape" "four"."second"="kiwi" "four"."third"="watermelon" "one"="apple" "six"="false" "three"="cherry" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
			},
			Expected: `"hello"."five"="true" "hello"."four"."first"="grape" "hello"."four"."second"="kiwi" "hello"."four"."third"="watermelon" "hello"."one"="apple" "hello"."six"="false" "hello"."three"="cherry" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
			},
			Expected: `"hello"."world"."five"="true" "hello"."world"."four"."first"="grape" "hello"."world"."four"."second"="kiwi" "hello"."world"."four"."third"="watermelon" "hello"."world"."one"="apple" "hello"."world"."six"="false" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
			},
			Expected: `"five"="true" "four"."first"="grape" "four"."second"="kiwi" "four"."third"="watermelon" "one"="apple" "seven"="789" "six"="false" "three"="cherry" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
			},
			Expected: `"hello"."five"="true" "hello"."four"."first"="grape" "hello"."four"."second"="kiwi" "hello"."four"."third"="watermelon" "hello"."one"="apple" "hello"."seven"="789" "hello"."six"="false" "hello"."three"="cherry" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
			},
			Expected: `"hello"."world"."five"="true" "hello"."world"."four"."first"="grape" "hello"."world"."four"."second"="kiwi" "hello"."world"."four"."third"="watermelon" "hello"."world"."one"="apple" "hello"."world"."seven"="789" "hello"."world"."six"="false" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},



		{
			P: []byte(nil),
			NamePrefix: []string{},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
				"eight":8.8,
			},
			Expected: `"eight"="8.800000" "five"="true" "four"."first"="grape" "four"."second"="kiwi" "four"."third"="watermelon" "one"="apple" "seven"="789" "six"="false" "three"="cherry" "two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
				"eight":8.8,
			},
			Expected: `"hello"."eight"="8.800000" "hello"."five"="true" "hello"."four"."first"="grape" "hello"."four"."second"="kiwi" "hello"."four"."third"="watermelon" "hello"."one"="apple" "hello"."seven"="789" "hello"."six"="false" "hello"."three"="cherry" "hello"."two"="banana"`,
		},
		{
			P: []byte(nil),
			NamePrefix: []string{"hello", "world"},
			Map: map[string]interface{}{
				"one":"apple",
				"two":"banana",
				"three":"cherry",
				"four":map[string]interface{}{
					"first":"grape",
					"second":"kiwi",
					"third":"watermelon",
				},
				"five":true,
				"six":false,
				"seven":789,
				"eight":8.8,
			},
			Expected: `"hello"."world"."eight"="8.800000" "hello"."world"."five"="true" "hello"."world"."four"."first"="grape" "hello"."world"."four"."second"="kiwi" "hello"."world"."four"."third"="watermelon" "hello"."world"."one"="apple" "hello"."world"."seven"="789" "hello"."world"."six"="false" "hello"."world"."three"="cherry" "hello"."world"."two"="banana"`,
		},
	}


	for testNumber, test := range tests {

		p := append([]byte(nil), test.P...)

		p = AppendMap(p, test.Map, test.NamePrefix...)

		if expected, actual := test.Expected, string(p); expected != actual {
			t.Errorf("For test #%d, expected & actually got:\n==)>%s<(==\n==)>%s<(==", testNumber, expected, actual)
			continue
		}

	}
}
