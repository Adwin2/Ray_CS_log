// 递归转迭代
package main

import (
	"fmt"

	"github.com/Adwin2/personal-go-modules/stack"
)

type Frame struct {
	Stage int
	M     int
	N     int
}

func ack_ii(m, n int) int {
	stk := stack.NewStack[Frame]()
	stk.Push(Frame{Stage: 0, M: m, N: n})
	cur := 0

	for !stk.IsEmpty() {
		top, _ := stk.Pop()

		if top.Stage == 0 {
			if top.M == 0 {
				cur = top.N + 1
			} else if top.N == 0 {
				stk.Push(Frame{Stage: 0, M: top.M - 1, N: 1})
			} else {
				stk.Push(Frame{Stage: 1, M: top.M})
				stk.Push(Frame{Stage: 0, M: top.M, N: top.N - 1})
			}
		} else if top.Stage == 1 {
			stk.Push(Frame{Stage: 0, M: top.M - 1, N: cur})
		}
	}
	return cur
}

func main() {
	fmt.Println(ack_ii(2, 1))
}
