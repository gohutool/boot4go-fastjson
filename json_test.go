package fastjson

import (
	"fmt"
	"testing"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : json_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/27 09:58
* 修改历史 : 1. [2022/4/27 09:58] 创建文件 by LongYong
*/

type JsonObj struct {
	Name string `json:"name"`
	No   string `json:"no"`
}

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Ok   bool   `json:"ok"`
	Json JsonObj
}

type Test2 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Ok   bool   `json:"ok"`
	Json []*JsonObj
}

func (t *Test) Unmarshall(value Value) error {
	t.Name = string(value.GetStringBytes("name"))
	t.Age = value.GetInt("age")
	t.Ok = value.GetBool("ok")
	t.Json = JsonObj{}
	t.Json.Unmarshall(*value.Get("json"))
	return nil
}

func (t *Test2) Unmarshall(value Value) error {
	t.Name = string(value.GetStringBytes("name"))
	t.Age = value.GetInt("age")
	t.Ok = value.GetBool("ok")

	if list, err := UnmarshallObjectList(*value.Get("json"), &JsonObj{}); err == nil {
		t.Json = list
		return nil
	} else {
		return err
	}
}

func (t *JsonObj) Unmarshall(value Value) error {
	t.Name = value.GetString("name")
	t.No = value.GetString("no")

	return nil
}

func TestJsonOne(t *testing.T) {
	str := "{\"name\":120, \"age\":\"2\", \"ok\":\"on\",\"json\":{\"name\":\"ddddd\", \"no\":7}}"

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test{}

	if err := UnmarshallObject(str, test); err == nil {
		fmt.Printf("%+v", test)
	} else {
		fmt.Printf("%+v", err)
	}

}

func TestJsonList(t *testing.T) {
	str := "[{\"name\":120, \"age\":\"2\", \"ok\":\"on\",\"json\":{\"name\":\"ddddd\", \"no\":7}}]"

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test{}

	if v, err := UnmarshallJson(str); err == nil {
		fmt.Printf("JsonValue %+v\n", v)

		if list, err := UnmarshallObjectList(*v, test); err == nil {
			fmt.Printf("List %+v\n", list)
		} else {
			fmt.Printf("Error %+v\n", err)
		}

	} else {
		fmt.Printf("%+v\n", err)
	}
}

func TestJsonList2(t *testing.T) {
	str := "[{\"name\":120, \"age\":\"2\", \"ok\":\"on\",\"json\":[{\"name\":\"ddddd\", \"no\":7}]}]"

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test2{}

	if v, err := UnmarshallJson(str); err == nil {
		fmt.Printf("JsonValue %+v\n", v)

		if list, err := UnmarshallObjectList(*v, test); err == nil {
			fmt.Printf("List %+v\n", list)
		} else {
			fmt.Printf("Error %+v\n", err)
		}

	} else {
		fmt.Printf("%+v\n", err)
	}
}
