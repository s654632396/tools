package main

import (
	"github.com/s654632396/decstree/seltree"
)

func main() {
	// test
	// 初始化树 & 根节点
	tree := new(seltree.SelTree)
	root := seltree.NewSampleNode()
	tree.Init(root)
	// 初始化一些节点数据
	nodeA1 := seltree.NewSampleNode().Register(myJudgement())
	nodeA2 := seltree.NewSampleNode().Register(myJudgement())
	nodeA3 := seltree.NewSampleNode().Register(myJudgement())
	nodeA4 := seltree.NewSampleNode().Register(myJudgement())
	nodeB1 := seltree.NewSampleNode().Register(myJudgement())
	nodeB2 := seltree.NewSampleNode().Register(myJudgement())
	nodeB3 := seltree.NewSampleNode().Register(myJudgement())
	nodeC1 := seltree.NewSampleNode().Register(myJudgement())
	nodeC2 := seltree.NewSampleNode().Register(myJudgement())
	nodeC3 := seltree.NewSampleNode().Register(myJudgement())

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

func myJudgement() seltree.Judgement {
	return func(self seltree.INode, args []interface{}) bool {

		var (
			arg1 int
			ok   bool
		)
		if arg1, ok = args[0].(int); !ok {
			println("arg1 invalid")
		}
		println(arg1)

		self.SetState(seltree.NodeStateAsked)

		if arg1 <= 10 {
			return false
		} else {
			return true
		}
	}

}
