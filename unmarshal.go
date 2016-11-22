package dotquote


import (
	"fmt"
)


// Unmarshal parses the dotquote-encoded data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {

	switch t := v.(type) {
	case map[string][]string:
		return unmarshalMapStringStringSlice(data, t)

	default:
		return fmt.Errorf("Cannot unmarshal dotquote data into %T", v)
	}
}


func unmarshalMapStringStringSlice(b []byte, m map[string][]string) error {

	decoder := Decoder{
		Bytes: b,
	}

	for decoder.Next() {
		key, err := decoder.KeyString()
		if nil != err {
			return fmt.Errorf("Problem decoding key: (%T) %v", err, err)
		}

		if _, ok := m[key]; !ok {
			m[key] = []string{}
		}

		values := decoder.Values()
		for values.Next() {
			value, err := values.UnquotedValueString()
			if nil != err {
				return fmt.Errorf("Problem decoding value: (%T) %v", err, err)
			}

			m[key] = append(m[key], value)
		}
		if err := values.Err(); nil != err {
			return fmt.Errorf("Received error when decoding values: (%T) %v", err, err)
		}
	}
	if err := decoder.Err(); nil != err {
		return fmt.Errorf("Received error when decoding: (%T) %v", err, err)
	}

	return nil
}
