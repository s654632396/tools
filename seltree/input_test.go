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
		// expectString = "default values"
		t.Error(err)
		t.Fail()
	}
	t.Log("test for string input:", expectString)

}

func TestIntInput_Resolve(t *testing.T) {
	// 使用int类型的input
	myInput = NewInput(1024)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectInt int
	if err := myInput.Resolve(&expectInt); err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("test for integer input:", expectInt)
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