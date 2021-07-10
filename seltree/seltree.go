package seltree

import (
	"context"
	"errors"
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

type IRegisterFunc interface {
	RfType()  RegisterFuncType
}

// Quest Ask提问函数结构体
type Quest func(self INode, args []IInput)
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

func (dt *SelTree) Link(parent, child INode) {
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
		println(">>>> prepare asking ...")
		// ask for arguments
		answer := node.ask()
		// make decisions and get next node
		node = node.poll(dt, []IInput{answer, "yet_another_argument"})
		if node == nil {
			break
		}
	}

	println("done.")
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

//INode 节点接口
type INode interface {
	// Register 提供注册节点判断流程的函数体
	Register(registerFunc IRegisterFunc) INode
	// SetState 提供变更节点的状态
	SetState(state NodeState) INode

	// ask 询问，获取poll前的入参
	ask() IInput
	// poll 开始询问子节点, 满足条件则发动跳转，选出要移动目标的子节点
	poll(tree *SelTree, args []IInput) INode
	// determine 用来判断条件是否成立
	determine(args []IInput) bool
	// add 用position 来添加一个 choice
	add(pos int) bool
	// getPos 获取当前节点在树上的位置
	getPos() int
	// setPos 获取当前节点在树上的位置
	setPos(pos int)
	// getId 获取当前节点的唯一身份
	getId() interface{}
	// getState 获取节点状态
	getState() NodeState
	// attempts 节点询问数自增, 并返回节点尝试次数
	attempts() int
}

//IInput 输入接口
// TODO
type IInput interface{}
