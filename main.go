package main

import (
	"github.com/s654632396/seltree/seltree"
	"math/rand"
	"time"
)

func main() {
	testInput()
	// test
	// 初始化树 & 根节点
	tree := new(seltree.SelTree)
	root := seltree.NewSampleNode()
	tree.Init(root)
	// 初始化一些节点数据
	nodeA1 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA2 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA3 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeA4 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB1 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB2 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeB3 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC1 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC2 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeC3 := seltree.NewSampleNode().Register(myJudgement()).Register(myQuest())

	// 构建节点
	tree.Link(root, nodeA1)
	tree.Link(root, nodeA2)
	tree.Link(root, nodeA3)
	tree.Link(root, nodeA4)
	tree.Link(nodeA1, nodeB1)
	tree.Link(nodeA2, nodeB2)
	tree.Link(nodeA3, nodeB3)
	tree.Link(nodeA4, nodeC1)
	tree.Link(nodeA4, nodeC2)
	tree.Link(nodeA4, nodeC3)

	// 运行树
	tree.Start()
}

func testInput()  {
	var myInput seltree.IInput
	myInput = seltree.NewInput("this is my string input argument")

	var expectValue string
	if err := myInput.ResolveValue(&expectValue); err!= nil {
		expectValue = "default values"
	}
	println("test for input:", expectValue)
}

func myJudgement() seltree.Judge {
	return func(self seltree.INode, args []seltree.IInput) bool {

		self.SetState(seltree.NodeStateAsked)

		var  arg1 int
		var  arg2 string
		if args[0].IsEmpty() {
			return  true
		}
		// parsing args
		if err := args[0].ResolveValue(&arg1); err!= nil {
			arg1 = 0
		}
		if err := args[1].ResolveValue(&arg2); err!= nil {
			arg2 = ""
		}
		println("arg1 = " , arg1, ", arg2=", arg2)

		if arg1 < 10 {
			return false
		} else {
			return true
		}
	}

}

// 测试用，询问函数体
func myQuest() seltree.Quest {
	return func(preArgs ...seltree.IInput) (answer seltree.IInput) {
		time.Sleep(time.Nanosecond * 2500)
		rand.Seed(time.Now().UnixNano())
		var  v = rand.Intn(50)
		answer = seltree.NewInput(v)
		return
	}
}
