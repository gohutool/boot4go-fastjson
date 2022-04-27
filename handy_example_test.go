package fastjson

import (
	"fmt"
)

func ExampleGetString() {
	data := []byte(`{"foo":{"bar":[123,"baz"]}}`)

	s := GetString(data, "foo", "bar", "1")
	fmt.Printf("data.foo.bar[1] = %s", s)

	// Output:
	// data.foo.bar[1] = baz
}

func ExampleGetInt() {
	data := []byte(`{"foo": [233,true, {"bar": [2343]} ]}`)

	n1 := GetInt(data, "foo", "0")
	fmt.Printf("data.foo[0] = %d\n", n1)

	n2 := GetInt(data, "foo", "2", "bar", "0")
	fmt.Printf("data.foo[2].bar[0] = %d\n", n2)

	// Output:
	// data.foo[0] = 233
	// data.foo[2].bar[0] = 2343
}

func ExampleExists() {
	data := []byte(`{"foo": [1.23,{"bar":33,"baz":null}]}`)

	fmt.Printf("exists(data.foo) = %v\n", Exists(data, "foo"))
	fmt.Printf("exists(data.foo[0]) = %v\n", Exists(data, "foo", "0"))
	fmt.Printf("exists(data.foo[1].baz) = %v\n", Exists(data, "foo", "1", "baz"))
	fmt.Printf("exists(data.foobar) = %v\n", Exists(data, "foobar"))
	fmt.Printf("exists(data.foo.bar) = %v\n", Exists(data, "foo", "bar"))

	// Output:
	// exists(data.foo) = true
	// exists(data.foo[0]) = true
	// exists(data.foo[1].baz) = true
	// exists(data.foobar) = false
	// exists(data.foo.bar) = false
}
