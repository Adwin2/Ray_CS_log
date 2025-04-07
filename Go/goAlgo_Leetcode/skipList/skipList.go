package main

import (
	"math"
	"math/rand/v2"
)

const (
	maxLevel    = 16
	probability = 0.25
)

type Node struct {
	value   int
	forward []*Node
}

type SkipList struct {
	head   *Node
	level  int
	length int
}

func NewSkipList() *SkipList {
	head := &Node{
		value:   math.MinInt32,
		forward: make([]*Node, maxLevel),
	}
	return &SkipList{
		head:   head,
		level:  1,
		length: 0,
	}
}

// randomLevel for Node
func randomLevel() int {
	level := 1
	for rand.Float64() < probability && level < maxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Insert(val int) {
	update := make([]*Node, maxLevel)
	cur := sl.head

	// traverse from top level
	for i := sl.level - 1; i >= 0; i-- {
		for cur.forward[i].value < val {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur == nil || cur.value != val {
		newLevel := randomLevel()

		if newLevel > sl.level {
			for i := sl.level; i < newLevel; i++ {
				update[i] = sl.head
			}
			sl.level = newLevel
		}
		newNode := &Node{
			value:   value,
			forward: make([]*Node, newLevel),
		}
		for i := 0; i < newLevel; i++ {
			newNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = newNode
		}
		sl.length++
	}
}

// Search
func (sl *SkipList)Search(val int) bool {
	cur := sl.head
	for i := sl.level-1;i>=0;i-- {
		for cur.forward[i] != nil && cur.forward[i].value < val {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	return cur != nil && cur.value == val
}

//
func (sl *SkipList) Delete(val int) bool {
	if sl.Search(val) {
		for i := 0;i< sl.level;i++ {
			if update[i].forward[i] != cur
		}
	}
}


func main() {

}
