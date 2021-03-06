/*
Package dotquote serializes data into the "dotquote" format.

The "dotquote" format is a more human-readable for a list of namespaced key-valued pairs.

The "dotquote" format is also useful for fitting your entire data on a single line of text.

Example

	"app"."name"="myapi" "app"."build"."number"="23" "apple"="one" "banana"="two" "cherry"="three" "trace"."id"="DtehCQqBnw93Tw4h"

The example dotquote line could have been generated by the following code:

	m := map[string]interface{}{
		"app":map[string]interface{}{
			"name":"myapi",
			"build":map[string]interface{
				"number":23,
			},
		},
		"apple":"one",
		"banana":"two",
		"cherry":"three",
		"trace":map[string]interface{}{
			"id":"DtehCQqBnw93Tw4h",
		},
	}
	
	var p []byte
	p = dotquote.AppendMap(p, m)

Alternatively, it that example dotquote line could have been generated by the following code:

	var p []byte
	
	p = dotquote.AppendString(p, "myapi", "app", "name")
	p = dotquote.AppendString(p, "23",    "app", "build", "number")
	p = dotquote.AppendString(p, "one",   "apple")
	p = dotquote.AppendString(p, "two",   "banana")
	p = dotquote.AppendString(p, "three", "cherry")
	p = dotquote.AppendString(p, "DtehCQqBnw93Tw4h", "trace", "id")
*/
package dotquote
