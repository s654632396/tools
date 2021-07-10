package seltree

import (
	"math/rand"
	"testing"
	"time"
)

func TestSelTree_Start(t *testing.T) {
	// test
	// 初始化树 & 根节点
	tree := new(SelTree)
	root := NewSampleNode()
	tree.Init(root)
	// 初始化一些节点数据
	nodeGroupA := make([]INode, 3)
	nodeGroupB := make([]INode, 3)
	nodeGroupC := make([]INode, 3)
	nodeGroupD := make([]INode, 3)
	nodeGroupA[0] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupA[1] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupA[2] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupB[0] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupB[1] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupB[2] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupC[0] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupC[1] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupC[2] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupD[0] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupD[1] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	nodeGroupD[2] = NewSampleNode().Register(myJudgement()).Register(myQuest())
	// 构建节点
	tree.Link(root, nodeGroupA ...)
	tree.Link(nodeGroupA[0], nodeGroupB...)
	tree.Link(nodeGroupA[1], nodeGroupC...)
	tree.Link(nodeGroupA[2], nodeGroupD...)
	// 运行树
	tree.Start()
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
