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
	judgement Judgement
	state     NodeState
	askTries  int
}

func NewSampleNode() *sampleNode {
	return &sampleNode{
		uuid:     uuid.NewString(),
		position: NodePositionUnknown,
		choices:  make([]int, 0),
		judgement: func(self INode, args []IInput) bool {
			// default judgement ,always return false
			return false
		},
	}
}

// Register 注册节点的判断逻辑
func (n *sampleNode) Register(judgement Judgement) INode {
	n.judgement = judgement
	return n
}

// SetState implements
func (n *sampleNode) SetState(state NodeState) INode {
	n.state = state
	return n
}

func (n *sampleNode) ask() IInput {
	// 先返回随机数，之后变更为节点询问的自定义函数体的调用
	time.Sleep(time.Nanosecond * 2500)
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(100)
	return answer
}

func (n *sampleNode) judge(args []IInput) bool {
	if n.judgement == nil {
		panic(errors.New(`this node without judgement`))
	}
	return n.judgement(n, args)
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
			if node.askCount() >= MaxAskTimes {
				remain--
				continue
			}

			if node.judge(args) {
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
func (n sampleNode) pos() int {
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
func (n sampleNode) identify() interface{} {
	return n.uuid
}

func (n *sampleNode) getState() NodeState {
	return n.state
}

func (n *sampleNode) askCount() int {
	n.askTries++
	return n.askTries
}
