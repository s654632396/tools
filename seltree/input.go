package seltree

import "errors"

type EmptyInput struct {}

func (receiver EmptyInput) IsEmpty() bool {
	return true
}

func (EmptyInput) ResolveValue(expect interface{}) (err error)  {
	// don't modify expect
	// 这是空输入参数
	// 单纯形式上继承IInput
	return
}

type NotEmptyInput struct {}

func (receiver NotEmptyInput) IsEmpty() bool {
	return false
}

/*
对Input参数的解析：
```go
	var myInput seltree.IInput
	myInput = seltree.NewInput("this is my string input argument")

	var expectValue string
	if err := myInput.ResolveValue(&expectValue); err!= nil {
		expectValue = "default values"
	}
	println("test for input:", expectValue)
```

从这里开始补充添加想要基本类型的Input参数结构体
...
 */

type StringInput struct {
	value string
	NotEmptyInput
}

type IntInput struct {
	value int
	NotEmptyInput
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



func (i StringInput) ResolveValue(expect interface{}) (err error) {
	if v ,ok := expect.(*string); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve string get unexpected value`)
	}
	return
}
func (i IntInput) ResolveValue(expect interface{}) (err error) {
	if v ,ok := expect.(*int); ok {
		*v = i.value
	} else {
		err = errors.New(`resolve int get unexpected value`)
	}
	return
}
