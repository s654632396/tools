package decstree

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
		judgement: func(self INode, args []interface{}) bool {
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

func (n *sampleNode) ask() interface{} {
	// println(n.uuid + " @(" + strconv.Itoa(n.position) + ") " + "ask for something")
	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(11)
	time.Sleep(time.Second)
	// println(n.uuid + " @(" + strconv.Itoa(n.position) + ") " + "ans=" + strconv.Itoa(answer))
	return answer
}

func (n *sampleNode) judge(args []interface{}) bool {
	if n.judgement == nil {
		panic(errors.New(`this node without judgement`))
	}
	// println(n.uuid + " @(" + strconv.Itoa(n.position) + ") " + "judge something by args")
	return n.judgement(n, args)
}

func (n *sampleNode) make(tree *DecsTree, args []interface{}) INode {
	// println(n.uuid + " @(" + strconv.Itoa(n.position) + ") " + "make ... ")

	var total = len(n.choices)
	var remain = total

	var decsionPos = -1
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
				decsionPos = pos
				break
			}

		}
	}

	if decsionPos == -1 {
		return nil
	}
	return tree.index(decsionPos)
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

// 节点身份证
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
