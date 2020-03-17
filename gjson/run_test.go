package gjson

import (
	"testing"

	"github.com/tidwall/gjson"
)

//
// https://github.com/tidwall/gjson
//

const jsonStr = `
	{
	  "name": {"first": "Tom", "last": "Anderson"},
	  "age":37,
	  "children": ["Sara","Alex","Jack","Cat"],
	  "fav.movie": "Deer Hunter",
	  "friends": [
		{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
		{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
	  ]
	}
`

func TestRun(t *testing.T) {
	val := gjson.Get(jsonStr, "name.first")
	t.Logf("name.first: %+v", val.String())

	val = gjson.Get(jsonStr, "age")
	t.Logf("age: %+v", val.String())

	val = gjson.Get(jsonStr, "friends.#")
	t.Logf("friends.#: %+v", val.String())

	val = gjson.Get(jsonStr, "friends.1.first")
	t.Logf("friends.1.first: %+v", val.String())

	val = gjson.Get(jsonStr, "friends.#.first")
	t.Logf("friends.#.first: %+v", val.Array()[0].String())
}
