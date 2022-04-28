package fastjson

import (
	"fmt"
	"reflect"
	"strings"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : plus.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/27 17:14
* 修改历史 : 1. [2022/4/27 17:14] 创建文件 by LongYong
*/

type Unmarshalable interface {
	Unmarshal(value *Value) error
}

func Unmarshal(value *Value, obj Unmarshalable) error {
	return obj.Unmarshal(value)
}

func UnmarshalJson(s string) (*Value, error) {
	var p Parser
	return p.Parse(s)
}

func UnmarshalObject(s string, obj Unmarshalable) error {
	if obj == nil {
		panic("Nil object can not unmarshal")
	}

	var p Parser
	if v, err := p.Parse(s); err != nil {
		return err
	} else {
		return Unmarshal(v, obj)
	}
}

func UnmarshalObjectMap[T Unmarshalable](value *Value, obj T) (map[string]T, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			t := reflect.TypeOf(obj)

			isPtr := t.Kind() == reflect.Pointer

			rtn := make(map[string]T)

			for _, v := range values.kvs {
				if isPtr {
					inst := reflect.New(t.Elem()).Interface().(T)
					inst.Unmarshal((v.v))
					rtn[v.k] = inst
				} else {
					inst := reflect.New(t).Interface().(Unmarshalable)
					inst.Unmarshal((v.v))
					rtn[v.k] = reflect.ValueOf(inst).Elem().Interface().(T)
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalObjectList[T Unmarshalable](value *Value, obj T) ([]T, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			t := reflect.TypeOf(obj)

			isPtr := t.Kind() == reflect.Pointer

			rtn := make([]T, 0, len(values))

			for _, v := range values {
				if isPtr {
					inst := reflect.New(t.Elem()).Interface().(T)
					inst.Unmarshal(v)
					rtn = append(rtn, inst)
				} else {
					inst := reflect.New(t).Interface().(Unmarshalable)
					inst.Unmarshal(v)
					rtn = append(rtn, reflect.ValueOf(inst).Elem().Interface().(T))
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}

		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseBoolList(json string) ([]bool, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalBoolList(v)
	} else {
		return nil, err
	}
}

func ParseBoolMap(json string) (map[string]bool, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalBoolMap(v)
	} else {
		return nil, err
	}
}

func UnmarshalBoolMap(value *Value) (map[string]bool, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]bool)

			for _, v := range values.kvs {
				if o, err := v.v.Bool(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalBoolList(value *Value) ([]bool, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]bool, 0, len(values))

			for _, v := range values {
				if r, err := v.Bool(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}
func ParseInt64List(json string) ([]int64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalInt64List(v)
	} else {
		return nil, err
	}
}

func ParseInt64Map(json string) (map[string]int64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalInt64Map(v)
	} else {
		return nil, err
	}
}

func UnmarshalInt64Map(value *Value) (map[string]int64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]int64)

			for _, v := range values.kvs {
				if o, err := v.v.Int64(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalInt64List(value *Value) ([]int64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]int64, 0, len(values))

			for _, v := range values {
				if r, err := v.Int64(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseIntList(json string) ([]int, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalIntList(v)
	} else {
		return nil, err
	}
}

func ParseIntMap(json string) (map[string]int, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalIntMap(v)
	} else {
		return nil, err
	}
}

func UnmarshalIntMap(value *Value) (map[string]int, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]int)

			for _, v := range values.kvs {
				if o, err := v.v.Int(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalIntList(value *Value) ([]int, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]int, 0, len(values))

			for _, v := range values {
				if r, err := v.Int(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseUintList(json string) ([]uint, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalUintList(v)
	} else {
		return nil, err
	}
}

func ParseUintMap(json string) (map[string]uint, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalUintMap(v)
	} else {
		return nil, err
	}
}

func UnmarshalUintMap(value *Value) (map[string]uint, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]uint)

			for _, v := range values.kvs {
				if o, err := v.v.Uint(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalUintList(value *Value) ([]uint, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]uint, 0, len(values))

			for _, v := range values {
				if r, err := v.Uint(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseUint64List(json string) ([]uint64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalUint64List(v)
	} else {
		return nil, err
	}
}

func ParseUint64Map(json string) (map[string]uint64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalUint64Map(v)
	} else {
		return nil, err
	}
}

func UnmarshalUint64Map(value *Value) (map[string]uint64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]uint64)

			for _, v := range values.kvs {
				if o, err := v.v.Uint64(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalUint64List(value *Value) ([]uint64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]uint64, 0, len(values))

			for _, v := range values {
				if r, err := v.Uint64(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}
func ParseFloat64List(json string) ([]float64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalFloat64List(v)
	} else {
		return nil, err
	}
}

func ParseFloat64Map(json string) (map[string]float64, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalFloat64Map(v)
	} else {
		return nil, err
	}
}

func UnmarshalFloat64Map(value *Value) (map[string]float64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]float64)

			for _, v := range values.kvs {
				if o, err := v.v.Float64(); err == nil {
					rtn[v.k] = o
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}

func UnmarshalFloat64List(value *Value) ([]float64, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]float64, 0, len(values))

			for _, v := range values {
				if r, err := v.Float64(); err == nil {
					rtn = append(rtn, r)
				} else {
					return rtn, err
				}
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseStringList(json string) ([]string, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalStringList(v)
	} else {
		return nil, err
	}
}

func UnmarshalStringList(value *Value) ([]string, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeArray {
		if values, err := value.Array(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values) == 0 {
				return nil, nil
			}

			rtn := make([]string, 0, len(values))

			for _, v := range values {
				rtn = append(rtn, v.s)
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit array", value.s)
	}
}

func ParseStringMap(json string) (map[string]string, error) {
	if v, err := UnmarshalJson(json); err == nil {
		return UnmarshalStringMap(v)
	} else {
		return nil, err
	}
}

func UnmarshalStringMap(value *Value) (map[string]string, error) {
	//func UnmarshalObjectList[T Unmarshalable](value Value, t reflect.Type) ([]T, error) {

	if value == nil {
		return nil, nil
	}

	if value.Type() == TypeObject {
		if values, err := value.Object(); err != nil {
			return nil, err
		} else {

			if values == nil || len(values.kvs) == 0 {
				return nil, nil
			}

			rtn := make(map[string]string)

			for _, v := range values.kvs {
				rtn[v.k] = v.v.s
			}

			return rtn, nil
		}
	} else {
		if len(strings.TrimSpace(value.s)) == 0 {
			return nil, nil
		}
		return nil, fmt.Errorf(" %q doesn't fit object", value.s)
	}
}
