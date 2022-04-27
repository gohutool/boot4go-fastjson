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

type JsonObj2 struct {
	Name string `json:"name"`
	No   int    `json:"no"`
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

type Test3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Ok   bool   `json:"ok"`
	Map  map[string]*JsonObj2
	Json []*JsonObj
}

func (t *Test3) Unmarshall(value Value) error {
	t.Name = string(value.GetStringBytes("name"))
	t.Age = value.GetInt("age")
	t.Ok = value.GetBool("ok")

	if list, err := UnmarshallObjectList(*value.Get("json"), &JsonObj{}); err == nil {
		t.Json = list
	} else {
		return err
	}

	if m, err := UnmarshallObjectMap(*value.Get("map"), &JsonObj2{}); err == nil {
		t.Map = m
		return nil
	} else {
		return err
	}
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

func (t *JsonObj2) Unmarshall(value Value) error {
	t.Name = value.GetString("name")
	t.No = value.GetInt("no")

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

func TestJsonMap(t *testing.T) {
	str := `
		{
			"AAA":{"name":"AAA-name","age":"22","ok":"on","json":[{"name":"ddddd111","no":27}]},
			"BBB":{"name":"BBB-name","age":"12","ok":"off","json":[{"name":"ddddd222","no":17}]}
		}
	`

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test2{}

	if v, err := UnmarshallJson(str); err == nil {
		fmt.Printf("JsonValue %+v\n", v)

		if m, err := UnmarshallObjectMap(*v, test); err == nil {
			fmt.Printf("List %+v\n", m)
		} else {
			fmt.Printf("Error %+v\n", err)
		}

	} else {
		fmt.Printf("%+v\n", err)
	}
}

func TestJsonMap2(t *testing.T) {
	str := `
		{
			"AAA":{"name":"AAA-name","age":"22","ok":"on","json":[{"name":"ddddd111","no":27}], 
						"map":{"A-1":{"name":"A1-ddddd111","no":"27"}, "A-2":{"name":"A2-ddddd111","no":227}} },
			"BBB":{"name":"BBB-name","age":"12","ok":"off","json":[{"name":"ddddd222","no":17}], 
						"map":{"B-1":{"name":"B1-ddddd111","no":27}, "B-2":{"name":"B2-ddddd111","no":"127"}} }
		}
	`

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test3{}

	if v, err := UnmarshallJson(str); err == nil {
		fmt.Printf("JsonValue %+v\n", v)

		if m, err := UnmarshallObjectMap(*v, test); err == nil {
			fmt.Printf("List %+v\n", m)
		} else {
			fmt.Printf("Error %+v\n", err)
		}

	} else {
		fmt.Printf("%+v\n", err)
	}
}
