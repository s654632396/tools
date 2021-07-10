package seltree


// IRegisterFunc 注册函数体接口
type IRegisterFunc interface {
	RfType()  RegisterFuncType
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
type IInput interface{
	ResolveValue(interface{}) error
	IsEmpty() bool
}
