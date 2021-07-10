package seltree

import (
	"errors"
)

type EmptyInput struct{}

func (receiver EmptyInput) IsEmpty() bool {
	return true
}

func (EmptyInput) Resolve(interface{}) (err error) {
	// don't modify expect
	// 这是空输入参数
	// 单纯形式上继承IInput
	return
}

type NotEmptyInput struct{}

func (receiver NotEmptyInput) IsEmpty() bool {
	return false
}

// NewInput 创建一个 IInput 类型实体
//
// 目前支持基础的类型：string,int,int32,int64,uint,uint32,uint64,float,bool
//
// TODO: 传递复杂的数据类型, 例如 slice,chan,map,struct
func NewInput(value interface{}) IInput {
	empty := EmptyInput{}
	// force cast
	switch v := value.(type) {
	case int32:
		value = int64(v)
	case int:
		value = int64(v)
	case uint:
		value = uint64(v)
	case uint32:
		value = uint64(v)
	case float32:
		value = float64(v)
	}

	switch v := value.(type) {
	case int64:
		return Int64Input{value: v}
	case uint64:
		return Uint64Input{value: v}
	case string:
		return StringInput{value: v}
	case float64:
		return Float64Input{value: v}
	case bool:
		return BooleanInput{value: v}
	case nil:
		return empty
	default:
		panic(`value type not support now`)
	}
}

// StringInput wrapper of string value Input argument
type StringInput struct {
	value string
	NotEmptyInput
}

func (i StringInput) Resolve(expect interface{}) (err error) {
	if v, ok := expect.(*string); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve string get unexpected value`)
	}
	return
}

// Uint64Input wrapper of string value Input argument
type Uint64Input struct {
	value uint64
	NotEmptyInput
}

func (i Uint64Input) Resolve(expect interface{}) (err error) {
	if v, ok := expect.(*uint64); ok {
		*v = i.value
	} else if v, ok := expect.(*uint); ok {
		*v = uint(i.value)
	} else if v, ok := expect.(*uint32); ok {
		*v = uint32(i.value)
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}

// Int64Input wrapper of string value Input argument
type Int64Input struct {
	value int64
	NotEmptyInput
}

func (i Int64Input) Resolve(expect interface{}) (err error) {
	if v, ok := expect.(*int64); ok {
		*v = i.value
	} else if v, ok := expect.(*int); ok {
		*v = int(i.value)
	} else if v, ok := expect.(*int32); ok {
		*v = int32(i.value)
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}

// Float64Input wrapper of string value Input argument
type Float64Input struct {
	value float64
	NotEmptyInput
}

func (i Float64Input) Resolve(expect interface{}) (err error) {
	if v, ok := expect.(*float64); ok {
		*v = i.value
	} else if v, ok := expect.(*float32); ok {
		*v = float32(i.value)
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}

// BooleanInput wrapper of string value Input argument
type BooleanInput struct {
	value bool
	NotEmptyInput
}

func (i BooleanInput) Resolve(expect interface{}) (err error) {
	if v, ok := expect.(*bool); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}
