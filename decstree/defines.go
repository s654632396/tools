package decstree

const MaxAskTimes = 10         // 一次决策中，被询问的最大次数
const NodePositionUnknown = -1 // 未定义位置

type NodeState int

const (
	NodeStatePending NodeState = 0 // 待询问
	NodeStateAsked   NodeState = 1 // 已被询问
	NodeStateWaitAsk NodeState = 2 // 询问过，但是可以被再次询问
)

const (
	MakeStrategyPolling = 0 // 轮询策略
	MakeStrategyRandom  = 1 // 随机策略
)
