package dotquote


import (
	"fmt"
	"sort"
)


func AppendMap(p []byte, mapData map[string]interface{}, namePrefix ...string) []byte {

	sortedKeys := make([]string, len(mapData))
	{
		i := 0
		for key, _ := range mapData {
			sortedKeys[i] = key
			i++
		}
	}
	sort.Strings(sortedKeys)


	for i, key := range sortedKeys {

		value := mapData[key]

		name := append([]string(nil), namePrefix...)
		name  = append(name, key)

		switch x := value.(type) {
		case string:
			if 0 < i {
				p = append(p, ' ')
			}
			p = AppendString(p, x, name...)

		case int, int8, int16, int32, int64,
		    uint,uint8,uint16,uint32,uint64:
			if 0 < i {
				p = append(p, ' ')
			}
			p = AppendString(p, fmt.Sprintf("%d", x), name...)

		case float32, float64:
			if 0 < i {
				p = append(p, ' ')
			}
			p = AppendString(p, fmt.Sprintf("%f", x), name...)

		case bool:
			if 0 < i {
				p = append(p, ' ')
			}
			p = AppendString(p, fmt.Sprintf("%t", x), name...)

		case map[string]interface{}:
			if 0 < i {
				p = append(p, ' ')
			}
			p = AppendMap(p, x, name...)
		default:
			// Nothing here.
		}
	}

	return p
}
