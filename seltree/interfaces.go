package seltree

// IInput Generic Input Type
//	// 泛型参数的接口
//
//	// 对Input参数的解析：
//	// 声明 IInput 实体
//	var myInput IInput
//	// 使用string类型的input
//	myInput = NewInput("this is my string input argument")
//
//	// 将myInput实体的值解析到一个期望是string类型的变量上
//	var expectString string
//	if err := myInput.Resolve(&expectString); err!= nil {
//		expectString = "default values"
//	}
//
//	// 使用int类型的input
//	myInput = NewInput(1024)
//	// 将myInput实体的值解析到一个期望是int类型的变量上
//	var  expectInt int
//	if err := myInput.Resolve(&expectInt); err!= nil {
//		expectInt = 0
//	}
//
//	// 判断一个IInput是否是EmptyInput
//	var myInput IInput
//	myInput = NewInput("this is my string input argument")
//
//	if myInput.IsEmpty() {
//		// empty here ..
//	} else {
//		// not empty here
//	}
//
type IInput interface{

	// Resolve 将receiver的value解析到expect上
	// 参数 interface{} 必须是 receiver.value 类型的指针类型
	Resolve(interface{}) error

	// IsEmpty 判断entity是否是一个 EmptyInput
	IsEmpty() bool
}

// IRegisterFunc 注册函数体接口
type IRegisterFunc interface {
	// RfType 获取这个注册函数体的类型
	RfType()  RegisterFuncType
}

//INode 节点接口
type INode interface {

	// Register 提供注册节点判断流程的函数体
	Register(registerFunc IRegisterFunc) INode

	// SetState 提供变更节点的状态
	SetState(state NodeState) INode

	// --------- private access -------------

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
