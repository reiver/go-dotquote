package dotquote


import (
	"strconv"

	"testing"
)


func TestAppendQuoteStringWithJustASingleRune(t *testing.T) {

	tests := []struct{
		P      []byte
		Rune     rune
		Expected string
	}{
		{
			P: []byte(nil),
			Rune:     rune(0), // NUL
			Expected: `"\x00"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(1), // SOH
			Expected: `"\x01"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(2), // STX
			Expected: `"\x02"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(3), // ETX
			Expected: `"\x03"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(4), // EOT
			Expected: `"\x04"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(5), // ENQ
			Expected: `"\x05"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(6), // ACK
			Expected: `"\x06"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(7), // BEL
			Expected: `"\a"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(8), // BS
			Expected: `"\b"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(9), //HT
			Expected: `"\t"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(10), // LF
			Expected: `"\n"`,
		},

		{
			P: []byte(nil),
			Rune:     rune(11), // VT
			Expected: `"\v"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(12), // FF
			Expected: `"\f"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(13), // CR
			Expected: `"\r"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(14), // SO
			Expected: `"\x0e"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(15), // SI
			Expected: `"\x0f"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(16), // DLE
			Expected: `"\x10"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(17), // DC1
			Expected: `"\x11"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(18), // DC2
			Expected: `"\x12"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(19), // DC3
			Expected: `"\x13"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(20), // DC4
			Expected: `"\x14"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(21), // NAK
			Expected: `"\x15"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(22), // SYN
			Expected: `"\x16"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(23), // ETB
			Expected: `"\x17"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(24), // CAN
			Expected: `"\x18"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(25), // EM
			Expected: `"\x19"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(26), // SUB
			Expected: `"\x1a"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(27), // ESC
			Expected: `"\x1b"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(28), // FS
			Expected: `"\x1c"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(29), // GS
			Expected: `"\x1d"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(30), // RS
			Expected: `"\x1e"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(31), // US
			Expected: `"\x1f"`,
		},
		{
			P: []byte(nil),
			Rune:     ' ',
			Expected: `" "`,
		},
		{
			P: []byte(nil),
			Rune:     '!',
			Expected: `"!"`,
		},
		{
			P: []byte(nil),
			Rune:     '"',
			Expected: `"\""`,
		},
		{
			P: []byte(nil),
			Rune:     '#',
			Expected: `"#"`,
		},
		{
			P: []byte(nil),
			Rune:     '$',
			Expected: `"$"`,
		},
		{
			P: []byte(nil),
			Rune:     '%',
			Expected: `"%"`,
		},
		{
			P: []byte(nil),
			Rune:     '&',
			Expected: `"&"`,
		},
		{
			P: []byte(nil),
			Rune:     '\'',
			Expected: `"'"`,
		},
		{
			P: []byte(nil),
			Rune:     '(',
			Expected: `"("`,
		},
		{
			P: []byte(nil),
			Rune:     ')',
			Expected: `")"`,
		},
		{
			P: []byte(nil),
			Rune:     '*',
			Expected: `"*"`,
		},
		{
			P: []byte(nil),
			Rune:     '+',
			Expected: `"+"`,
		},
		{
			P: []byte(nil),
			Rune:     ',',
			Expected: `","`,
		},
		{
			P: []byte(nil),
			Rune:     '-',
			Expected: `"-"`,
		},
		{
			P: []byte(nil),
			Rune:     '.',
			Expected: `"."`,
		},
		{
			P: []byte(nil),
			Rune:     '/',
			Expected: `"/"`,
		},
		{
			P: []byte(nil),
			Rune:     '0',
			Expected: `"0"`,
		},
		{
			P: []byte(nil),
			Rune:     '1',
			Expected: `"1"`,
		},
		{
			P: []byte(nil),
			Rune:     '2',
			Expected: `"2"`,
		},
		{
			P: []byte(nil),
			Rune:     '3',
			Expected: `"3"`,
		},
		{
			P: []byte(nil),
			Rune:     '4',
			Expected: `"4"`,
		},
		{
			P: []byte(nil),
			Rune:     '5',
			Expected: `"5"`,
		},
		{
			P: []byte(nil),
			Rune:     '6',
			Expected: `"6"`,
		},
		{
			P: []byte(nil),
			Rune:     '7',
			Expected: `"7"`,
		},
		{
			P: []byte(nil),
			Rune:     '8',
			Expected: `"8"`,
		},
		{
			P: []byte(nil),
			Rune:     '9',
			Expected: `"9"`,
		},
		{
			P: []byte(nil),
			Rune:     ':',
			Expected: `":"`,
		},
		{
			P: []byte(nil),
			Rune:     ';',
			Expected: `";"`,
		},
		{
			P: []byte(nil),
			Rune:     '<',
			Expected: `"<"`,
		},
		{
			P: []byte(nil),
			Rune:     '=',
			Expected: `"="`,
		},
		{
			P: []byte(nil),
			Rune:     '>',
			Expected: `">"`,
		},
		{
			P: []byte(nil),
			Rune:     '?',
			Expected: `"?"`,
		},
		{
			P: []byte(nil),
			Rune:     '@',
			Expected: `"@"`,
		},
		{
			P: []byte(nil),
			Rune:     'A',
			Expected: `"A"`,
		},
		{
			P: []byte(nil),
			Rune:     'B',
			Expected: `"B"`,
		},
		{
			P: []byte(nil),
			Rune:     'C',
			Expected: `"C"`,
		},
		{
			P: []byte(nil),
			Rune:     'D',
			Expected: `"D"`,
		},
		{
			P: []byte(nil),
			Rune:     'E',
			Expected: `"E"`,
		},
		{
			P: []byte(nil),
			Rune:     'F',
			Expected: `"F"`,
		},
		{
			P: []byte(nil),
			Rune:     'G',
			Expected: `"G"`,
		},
		{
			P: []byte(nil),
			Rune:     'H',
			Expected: `"H"`,
		},
		{
			P: []byte(nil),
			Rune:     'I',
			Expected: `"I"`,
		},
		{
			P: []byte(nil),
			Rune:     'J',
			Expected: `"J"`,
		},
		{
			P: []byte(nil),
			Rune:     'K',
			Expected: `"K"`,
		},
		{
			P: []byte(nil),
			Rune:     'L',
			Expected: `"L"`,
		},
		{
			P: []byte(nil),
			Rune:     'M',
			Expected: `"M"`,
		},
		{
			P: []byte(nil),
			Rune:     'N',
			Expected: `"N"`,
		},
		{
			P: []byte(nil),
			Rune:     'O',
			Expected: `"O"`,
		},
		{
			P: []byte(nil),
			Rune:     'P',
			Expected: `"P"`,
		},
		{
			P: []byte(nil),
			Rune:     'Q',
			Expected: `"Q"`,
		},
		{
			P: []byte(nil),
			Rune:     'R',
			Expected: `"R"`,
		},
		{
			P: []byte(nil),
			Rune:     'S',
			Expected: `"S"`,
		},
		{
			P: []byte(nil),
			Rune:     'T',
			Expected: `"T"`,
		},
		{
			P: []byte(nil),
			Rune:     'U',
			Expected: `"U"`,
		},
		{
			P: []byte(nil),
			Rune:     'V',
			Expected: `"V"`,
		},
		{
			P: []byte(nil),
			Rune:     'W',
			Expected: `"W"`,
		},
		{
			P: []byte(nil),
			Rune:     'X',
			Expected: `"X"`,
		},
		{
			P: []byte(nil),
			Rune:     'Y',
			Expected: `"Y"`,
		},
		{
			P: []byte(nil),
			Rune:     'Z',
			Expected: `"Z"`,
		},
		{
			P: []byte(nil),
			Rune:     '[',
			Expected: `"["`,
		},
		{
			P: []byte(nil),
			Rune:     '\\',
			Expected: `"\\"`,
		},
		{
			P: []byte(nil),
			Rune:     ']',
			Expected: `"]"`,
		},
		{
			P: []byte(nil),
			Rune:     '^',
			Expected: `"^"`,
		},
		{
			P: []byte(nil),
			Rune:     '_',
			Expected: `"_"`,
		},
		{
			P: []byte(nil),
			Rune:     '`',
			Expected: "\"`\"",
		},
		{
			P: []byte(nil),
			Rune:     'a',
			Expected: `"a"`,
		},
		{
			P: []byte(nil),
			Rune:     'b',
			Expected: `"b"`,
		},
		{
			P: []byte(nil),
			Rune:     'c',
			Expected: `"c"`,
		},
		{
			P: []byte(nil),
			Rune:     'd',
			Expected: `"d"`,
		},
		{
			P: []byte(nil),
			Rune:     'e',
			Expected: `"e"`,
		},
		{
			P: []byte(nil),
			Rune:     'f',
			Expected: `"f"`,
		},
		{
			P: []byte(nil),
			Rune:     'g',
			Expected: `"g"`,
		},
		{
			P: []byte(nil),
			Rune:     'h',
			Expected: `"h"`,
		},
		{
			P: []byte(nil),
			Rune:     'i',
			Expected: `"i"`,
		},
		{
			P: []byte(nil),
			Rune:     'j',
			Expected: `"j"`,
		},
		{
			P: []byte(nil),
			Rune:     'k',
			Expected: `"k"`,
		},
		{
			P: []byte(nil),
			Rune:     'l',
			Expected: `"l"`,
		},
		{
			P: []byte(nil),
			Rune:     'm',
			Expected: `"m"`,
		},
		{
			P: []byte(nil),
			Rune:     'n',
			Expected: `"n"`,
		},
		{
			P: []byte(nil),
			Rune:     'o',
			Expected: `"o"`,
		},
		{
			P: []byte(nil),
			Rune:     'p',
			Expected: `"p"`,
		},
		{
			P: []byte(nil),
			Rune:     'q',
			Expected: `"q"`,
		},
		{
			P: []byte(nil),
			Rune:     'r',
			Expected: `"r"`,
		},
		{
			P: []byte(nil),
			Rune:     's',
			Expected: `"s"`,
		},
		{
			P: []byte(nil),
			Rune:     't',
			Expected: `"t"`,
		},
		{
			P: []byte(nil),
			Rune:     'u',
			Expected: `"u"`,
		},
		{
			P: []byte(nil),
			Rune:     'v',
			Expected: `"v"`,
		},
		{
			P: []byte(nil),
			Rune:     'w',
			Expected: `"w"`,
		},
		{
			P: []byte(nil),
			Rune:     'x',
			Expected: `"x"`,
		},
		{
			P: []byte(nil),
			Rune:     'y',
			Expected: `"y"`,
		},
		{
			P: []byte(nil),
			Rune:     'z',
			Expected: `"z"`,
		},

		{
			P: []byte(nil),
			Rune:     '{',
			Expected: `"{"`,
		},
		{
			P: []byte(nil),
			Rune:     '|',
			Expected: `"|"`,
		},
		{
			P: []byte(nil),
			Rune:     '}',
			Expected: `"}"`,
		},
		{
			P: []byte(nil),
			Rune:     '~',
			Expected: `"~"`,
		},
		{
			P: []byte(nil),
			Rune:     rune(127),
			Expected: `"\x7f"`,
		},



		{
			P: []byte(nil),
			Rune:     '✪', // CIRCLED WHITE STAR
			Expected: "\"\u272A\"",
		},



		{
			P: []byte(nil),
			Rune:     rune(8232), // LINE SEPARATOR
			Expected: `"\u2028"`,
		},



		{
			P: []byte(nil),
			Rune:     '🙂', // SLIGHTLY SMILING FACE
			Expected: "\"\U0001F642\"",
		},
	}


	for testNumber, test := range tests {
		p := append([]byte(nil), test.P...)

		p2 := appendQuoteString(p, string(test.Rune))

		if actual := p2; nil == actual {
			t.Errorf("For test #%d, did NOT expected nil, but actually got %v.", testNumber, actual)
			continue
		}

		if expected, actual := test.Expected, string(p2); expected != actual {
			t.Errorf("For test #%d, expected ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
			continue
		}

		// For some reason Golang encodes character 127 in \u format. So skipping that.
		if rune(127) != test.Rune {
			should := strconv.Quote( string(test.Rune) )

			if expected, actual := should, string(p2); expected != actual {
				t.Errorf("For test #%d, expected that it should be ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
				continue
			}
		}
	}
}


func TestAppendQuoteString(t *testing.T) {

	tests := []struct{
		P     []byte
		Slice []byte
		Expected string
	}{
		{
			P: []byte(nil),
			Slice: []byte{'a','p','p','l','e'},
			Expected: `"apple"`,
		},
		{
			P: []byte(nil),
			Slice: []byte{'b','a','n','a','n','a'},
			Expected: `"banana"`,
		},
		{
			P: []byte(nil),
			Slice: []byte{'c','h','e','r','r','y'},
			Expected: `"cherry"`,
		},



		{
			P: []byte(nil),
			Slice: []byte{'a','p','p','l','e', ' ', 'b','a','n','a','n','a',  ' ', 'c','h','e','r','r','y'},
			Expected: `"apple banana cherry"`,
		},



		{
			P: []byte(nil),
			Slice: []byte{'H','e','l','l','o',' ','w','o','r','l','d','!'},
			Expected: `"Hello world!"`,
		},



		{
			P: []byte(nil),
			Slice: []byte{'W','h','a','t',' ','d','o','e','s',' ','"','b','a','n','a','l','"',' ','m','e','a','n','?'},
			Expected: `"What does \"banal\" mean?"`,
		},



		{
			P: []byte(nil),
			Slice: []byte{'o','n','e','\n','t','w','o','\t','t','h','r','e','e','\r','\n'},
			Expected: `"one\ntwo\tthree\r\n"`,
		},



		{
			P: []byte(nil),
			Slice: []byte{'l','i','n','e',' ','1', byte(6), 'l','i','n','e',' ','2'},
			Expected: `"line 1\x06line 2"`,
		},
	}


	for testNumber, test := range tests {
		p := append([]byte(nil), test.P...)
		s := string(test.Slice)

		p2 := appendQuoteString(p, s)

		if actual := p2; nil == actual {
			t.Errorf("For test #%d, did NOT expected nil, but actually got %v.", testNumber, actual)
			continue
		}

		if expected, actual := test.Expected, string(p2); expected != actual {
			t.Errorf("For test #%d, expected ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
			continue
		}

		{
			should := strconv.Quote( string(test.Slice) )

			if expected, actual := should, string(p2); expected != actual {
				t.Errorf("For test #%d, expected that it should be ==)>%s<(==, but actually got ==)>%s<(==", testNumber, expected, actual)
				continue
			}
		}
	}
}
