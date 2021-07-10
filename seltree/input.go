package seltree

import "errors"

type EmptyInput struct {}

func (receiver EmptyInput) IsEmpty() bool {
	return true
}

func (EmptyInput) Resolve(interface{}) (err error)  {
	// don't modify expect
	// 这是空输入参数
	// 单纯形式上继承IInput
	return
}

type NotEmptyInput struct {}

func (receiver NotEmptyInput) IsEmpty() bool {
	return false
}


func NewInput(value interface{}) IInput {
	empty := EmptyInput{}
	switch v := value.(type) {
	case int:
		return  IntInput{ value: v}
	case string:
		return  StringInput{ value: v}
	case nil:
		return  empty
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
	if v ,ok := expect.(*string); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve string get unexpected value`)
	}
	return
}

type IntInput struct {
	value int
	NotEmptyInput
}

func (i IntInput) Resolve(expect interface{}) (err error) {
	if v ,ok := expect.(*int); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}