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
}

func (t *Test) Unmarshall(value Value) error {
	t.Name = string(value.GetStringBytes("name"))
	t.Age = value.GetInt("age")
	t.Ok = value.GetBool("ok")
	return nil
}

func TestJsonOne(t *testing.T) {
	str := "{\"name\":120, \"age\":\"2\", \"ok\":\"on\"}"

	//test := &Test{Name: "AAA", Age: 10}
	test := &Test{}

	UnmarshallObject(str, test)

	fmt.Printf("%+v", test)
}
