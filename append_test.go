package dotquote


import (
	"testing"
)


func TestAppend(t *testing.T) {

	tests := []struct{
		P      []byte
		Name [][]byte
		Value  []byte
		Expected string

	}{
		{
			P: []byte(nil),
			Name: [][]byte(nil),
			Expected: `""=""`,
		},
		{
			P: []byte(nil),
			Name: [][]byte{},
			Expected: `""=""`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'a','p','p','l','e'},
			},
			Value:  []byte(nil),
			Expected: `"apple"=""`,
		},
		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'a','p','p','l','e'},
			},
			Value:  []byte{},
			Expected: `"apple"=""`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'a','p','p','l','e'},
			},
			Value:  []byte{'b','a','n','a','n','a'},
			Expected: `"apple"="banana"`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'a','p','p','l','e'},
			},
			Value:  []byte{'b','a','n','a','n','a', ' ', 'c','h','e','r','r','y'},
			Expected: `"apple"="banana cherry"`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'a','p','p','l','e'},
			},
			Value:  []byte{
				'b','a','n','a','n','a',
				' ',
				'c','h','e','r','r','y',
				31,
				'k','i','w','i',
				' ',
				226,156,170,     // UTF-8 encoding of: ✪
				' ',
				226,128,168,     // UTF-8 encoding of: LINE SEPARATOR
				' ',
				240,159,153,130, // UTF-8 encoding of: 🙂
			},
			Expected: `"apple"="banana cherry\x1fkiwi ✪ \u2028 🙂"`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{'o','n','e'},
				[]byte{'t','w','o'},
				[]byte{'t','h','r','e','e'},
			},
			Value:  []byte{
				'a','p','p','l','e',
				'\t',
				'b','a','n','a','n','a',
				' ',
				'c','h','e','r','r','y',
				31,
				'k','i','w','i',
				' ',
				226,156,170,     // UTF-8 encoding of: ✪
				' ',
				226,128,168,     // UTF-8 encoding of: LINE SEPARATOR
				' ',
				240,159,153,130, // UTF-8 encoding of: 🙂
			},
			Expected: `"one"."two"."three"="apple\tbanana cherry\x1fkiwi ✪ \u2028 🙂"`,
		},



		{
			P: []byte(nil),
			Name: [][]byte{
				[]byte{ 0, 1, 2, 3, 4, 5, 6, 7, 8,9},
				[]byte{10,11,12,13,14,15,16,17,18,19},
				[]byte{20,21,22,23,24,25,26,27,28,29},
				[]byte{30,31,32,33,34,35,36,37,38,39},
			},
			Value:  []byte{
				'a','p','p','l','e',
				'\t',
				'b','a','n','a','n','a',
				' ',
				'c','h','e','r','r','y',
				31,
				'k','i','w','i',
				' ',
				226,156,170,     // UTF-8 encoding of: ✪
				' ',
				226,128,168,     // UTF-8 encoding of: LINE SEPARATOR
				' ',
				240,159,153,130, // UTF-8 encoding of: 🙂
			},
			Expected: `"\x00\x01\x02\x03\x04\x05\x06\a\b\t"."\n\v\f\r\x0e\x0f\x10\x11\x12\x13"."\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d"."\x1e\x1f !\"#$%&'"="apple\tbanana cherry\x1fkiwi ✪ \u2028 🙂"`,
		},
	}


	for testNumber, test := range tests {

		p := append([]byte(nil), test.P...)

		p = Append(p, test.Value, test.Name...)

		if expected, actual := test.Expected, string(p); expected != actual {
			t.Errorf("For test #%d, expected ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
			continue
		}

	}
}
