package seltree

import (
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

// sampleNode 示例节点
type sampleNode struct {
	uuid      string
	position  int
	choices   []int

	// 提问函数体, 由ask调用
	quest	Quest
	// 判断函数体, 由poll调用
	judge	Judge

	// 节点状态
	state     NodeState
	// 每次poll中的尝试次数
	attempt int
}

func NewSampleNode() *sampleNode {
	return &sampleNode{
		uuid:     uuid.NewString(),
		position: NodePositionUnknown,
		choices:  make([]int, 0),
		judge: func(self INode, args []IInput) bool {
			// default judgement ,always return false
			return false
		},
	}
}
func (n *sampleNode) Register(registerFunc IRegisterFunc) INode {
	switch f := registerFunc.(type) {
	case Quest:
		n.quest = f
	case Judge:
		n.judge = f
	default:
		panic(`unknown register func`)
	}

	return n
}

// SetState implements
func (n *sampleNode) SetState(state NodeState) INode {
	n.state = state
	return n
}

func (n *sampleNode) ask() IInput {
	// 先返回随机数，之后变更为节点询问的自定义函数体的调用
	// TODO : call node.quest here
	time.Sleep(time.Nanosecond * 2500)
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(100)
	return answer
}

func (n *sampleNode) determine(args []IInput) bool {
	if n.judge == nil {
		panic(errors.New(`this node without judge`))
	}
	return n.judge(n, args)
}

func (n *sampleNode) poll(tree *SelTree, args []IInput) INode {
	var total = len(n.choices)
	var remain = total

	var successPos = -1
	for remain > 0 {
		for idx := 0; idx < total; idx++ {
			pos := n.choices[idx]
			var node = tree.index(pos)
			if node == nil {
				panic(`make error: node is nil`)
			}
			if node.getState() == NodeStateAsked {
				remain--
				continue
			}
			if node.attempts() >= MaxAskTimes {
				remain--
				continue
			}

			if node.determine(args) {
				successPos = pos
				break
			}

		}
	}

	if successPos == -1 {
		return nil
	}
	return tree.index(successPos)
}

// 获取当前节点在树上的位置
func (n sampleNode) getPos() int {
	return n.position
}

// position自我标记
// 避免tree反复询问节点的位置
func (n *sampleNode) setPos(pos int) {
	if n.position != NodePositionUnknown {
		panic(errors.New(`cannot set position again`))
	}
	n.position = pos
}

// 添加子节点的 position
func (n *sampleNode) add(pos int) bool {
	var found bool
	for c := range n.choices {
		if c == pos {
			// already
			found = true
			break
		}
	}
	if !found {
		n.choices = append(n.choices, pos)
	}
	return found
}

// 节点身份
func (n sampleNode) getId() interface{} {
	return n.uuid
}

func (n *sampleNode) getState() NodeState {
	return n.state
}

func (n *sampleNode) attempts() int {
	n.attempt++
	return n.attempt
}
