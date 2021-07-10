package main

import (
	. "github.com/s654632396/seltree/seltree"
	"math/rand"
	"time"
)

func main() {
	testInput()
	// test
	// 初始化树 & 根节点
	tree := new(SelTree)
	root := NewSampleNode()
	tree.Init(root)
	// 初始化一些节点数据
	nodeA1 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA2 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA3 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA4 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB1 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB2 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB3 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC1 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC2 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC3 := NewSampleNode().Register(myJudgement()).Register(myQuest())
	// 构建节点
	tree.Link(root, nodeA1, nodeA2, nodeA3, nodeA4)
	tree.Link(nodeA1, nodeB1, nodeB2, nodeB3)
	tree.Link(nodeA4, nodeC1, nodeC2, nodeC3)
	// 运行树
	tree.Start()
}

// 测试输入类型
func testInput() {
	// 声明 IInput 实体
	var myInput IInput

	// 使用string类型的input
	myInput = NewInput("this is my string input argument")

	// 将myInput实体的值解析到一个期望是string类型的变量上
	var expectString string
	if err := myInput.Resolve(&expectString); err != nil {
		expectString = "default values"
	}
	println("test for string input:", expectString)

	// 使用int类型的input
	myInput = NewInput(1024)

	// 将myInput实体的值解析到一个期望是int类型的变量上
	var expectInt int
	if err := myInput.Resolve(&expectInt); err != nil {
		expectInt = 0
	}
	println("test for integer input:", expectInt)
	println()

}

func myJudgement() Judge {
	return func(self INode, args []IInput) bool {

		self.SetState(NodeStateAsked)

		var arg1 int
		var arg2 string
		if args[0].IsEmpty() {
			return true
		}
		// parsing args
		if err := args[0].Resolve(&arg1); err != nil {
			arg1 = 0
		}
		if err := args[1].Resolve(&arg2); err != nil {
			arg2 = ""
		}
		println("arg1 = ", arg1, ", arg2 =", arg2)

		if arg1 < 10 {
			return false
		} else {
			return true
		}
	}

}

// 测试用，询问函数体
func myQuest() Quest {
	return func(preArgs ...IInput) (answer IInput) {
		time.Sleep(time.Nanosecond * 2500)
		rand.Seed(time.Now().UnixNano())
		var v = rand.Intn(50)
		answer = NewInput(v)
		return
	}
}
