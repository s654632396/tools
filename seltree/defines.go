package seltree

const MaxAskTimes = 10         // 一次poll中，被询问的最大次数
const NodePositionUnknown = -1 // 未定义位置

type NodeState int
type RegisterFuncType int

const (
	NodeStatePending NodeState = 0 // 待询问
	NodeStateAsked   NodeState = 1 // 已被询问
	NodeStateWaitAsk NodeState = 2 // 询问过，但是可以被再次询问
)

const (
	MakeStrategyPolling = 0 // 轮询策略
	MakeStrategyRandom  = 1 // 随机策略
)

const (
	RegisterFuncTypeQuest RegisterFuncType = 0
	RegisterFuncTypeJudge RegisterFuncType = 1
)