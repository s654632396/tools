package seltree

import "testing"

// 声明 IInput 实体
var myInput IInput

func TestStringInput_Resolve(t *testing.T) {
	// 使用string类型的input
	myInput = NewInput("this is my string input argument")

	// 将myInput实体的值解析到一个期望是string类型的变量上
	var expectString string
	if err := myInput.Resolve(&expectString); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestIntInput_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin int = 1024
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue int
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("int value not equal")
		t.Fail()
	}
}

func TestInt32Input_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin int32 = -1024
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue int32
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("int32 value not equal")
		t.Fail()
	}
}

func TestInt64Input_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = int64(1<<32 + 10)
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue int64
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("int64 value not equal")
		t.Fail()
	}
}

func TestUintInput_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = uint(1<<32 + 10)
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue uint
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("uint value not equal")
		t.Fail()
	}
}

func TestUint32Input_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = uint32(1<<32 - 1)
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue uint32
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("bool value not equal")
		t.Fail()
	}
}

func TestUint64Input_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = uint64(1<<32 + 1024)
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue uint64
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("bool value not equal")
		t.Fail()
	}
}

func TestFloat64Input_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = 3.1415936
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue float64
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}

	if origin != expectValue {
		t.Error("bool value not equal")
		t.Fail()
	}
}

func TestBooleanInput_Resolve(t *testing.T) {
	// 使用int类型的input
	var origin = true
	myInput = NewInput(origin)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectValue bool
	if err := myInput.Resolve(&expectValue); err != nil {
		t.Error(err)
		t.Fail()
	}
	if origin != expectValue {
		t.Error("bool value not equal")
		t.Fail()
	}
}

func TestEmptyInput_IsEmpty(t *testing.T) {

	myInput = NewInput(nil)
	if !myInput.IsEmpty() {
		t.Error(`unexpected not empty input`)
		t.Fail()
	}
	myInput = NewInput("ok")
	if myInput.IsEmpty() {
		t.Error(`this must be a not empty input`)
		t.Fail()
	}
}