package seltree

import (
	"context"
	"errors"
	"log"
)

// SelTree 选择树
type SelTree struct {
	// global variable
	ctx *context.Context

	// use for search node
	pointer int

	// all nodes
	nodes []INode
}


// Quest Ask提问函数结构体
type Quest func(...IInput) (answer IInput)
// Judge 节点判断函数结构体
type Judge func(self INode, args []IInput) bool

func (Quest) RfType() RegisterFuncType  {
	return RegisterFuncTypeQuest
}

func (Judge) RfType() RegisterFuncType  {
	return RegisterFuncTypeJudge
}

// Init 设置root节点
func (dt *SelTree) Init(root INode) *SelTree {
	if dt.nodes != nil || len(dt.nodes) != 0 {
		panic(errors.New(`root node already exists`))
	}
	if root == nil {
		panic(errors.New(`root node must not be nil`))
	}
	dt.pointer = 0
	dt.nodes = make([]INode, 0)
	// root 永远都是第0位
	root.setPos(0)
	dt.nodes = append(dt.nodes, root)
	return dt
}

// Link 链接 parentNode 和 childNodes
// children 为复数不确定节点
func (dt *SelTree) Link(parent INode, children ...INode)  {
	for _, child := range children {
		dt.link(parent, child)
	}
}

// link 链接 parentNode 和 1个childNode
func (dt *SelTree) link(parent, child INode) {
	var found bool
	for _, find := range dt.nodes {
		if find.getId() == child {
			found = true
			break
		}
	}
	if !found {
		dt.nodes = append(dt.nodes, child)
	}
	dt.lookup(child)
	parent.add(child.getPos())
	return
}

// 查询节点在树上的位置
// 如果没查到，则返回 NodePositionUnknown
func (dt *SelTree) lookup(target INode) int {
	for idx, node := range dt.nodes {
		if node.getId() == target.getId() {
			target.setPos(idx)
			break
		}
	}
	return target.getPos()
}

func (dt *SelTree) Start() {
	var node = dt.current()

	for {
		// ask for arguments
		answer := node.ask()

		log.Print(">>>> asked: ")
		log.Printf("%#v\n", answer)

		if answer == nil {
			panic(`answer invalid: nil`)
		}
		// make decisions and get next node
		node = node.poll(dt, []IInput{answer, StringInput{value: "yet another argument"}})
		if node == nil {
			break
		}
	}

	log.Println("seltree is done.")
}

// current 获取当前节点
func (dt *SelTree) current() INode {
	var pointer = dt.pointer
	if len(dt.nodes) < pointer {
		panic(errors.New(`unexpected pointer called`))
	}
	return dt.nodes[pointer]
}

func (dt *SelTree) index(pos int) INode {
	if len(dt.nodes) <= pos {
		return nil
	}
	return dt.nodes[pos]
}
