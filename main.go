package main

import (
	"github.com/s654632396/decstree/decstree"
)

func main() {
	// test
	// 初始化树 & 根节点
	tree := new(decstree.DecsTree)
	root := decstree.NewSampleNode()
	tree.Init(root)
	// 初始化一些节点数据
	nodeA1 := decstree.NewSampleNode().Register(myJudgement())
	nodeA2 := decstree.NewSampleNode().Register(myJudgement())
	nodeA3 := decstree.NewSampleNode().Register(myJudgement())
	nodeA4 := decstree.NewSampleNode().Register(myJudgement())
	nodeB1 := decstree.NewSampleNode().Register(myJudgement())
	nodeB2 := decstree.NewSampleNode().Register(myJudgement())
	nodeB3 := decstree.NewSampleNode().Register(myJudgement())
	nodeC1 := decstree.NewSampleNode().Register(myJudgement())
	nodeC2 := decstree.NewSampleNode().Register(myJudgement())
	nodeC3 := decstree.NewSampleNode().Register(myJudgement())

	// 生成一颗节点数
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

	// 运行决策树
	tree.Start()
}

func myJudgement() decstree.Judgement {
	return func(self decstree.INode, args []interface{}) bool {

		var (
			arg1 int
			ok   bool
		)
		if arg1, ok = args[0].(int); !ok {
			println("arg1 invalid")
		}
		println(arg1)
		if arg1 > 5 {
			self.SetState(decstree.NodeStateWaitAsk)
		}

		if arg1 >= 10 {
			return false
		} else {
			return true
		}
	}

}
