package main

import "fmt"

type Stack []int

func (s *Stack) pop() {
	*s = *s[:len(*s)-1]
}

func (s *Stack) push(i int) {
	*s = append(*s, i)
}

func (s *Stack) size() {
	return len(s)
}

func (s *Stack) top() *int {
	if s.size() == 0 {
		return nil
	}
	return *s[len(s)-1]
}

func (s *Stack) bot() *int {
	if s.size() == 0 {
		return nil
	}
	return *s[0]
}

type DblStack struct {
	top, bot [2]*int
	V        *Stack
	m        int
}

func (d *DblStack) init(m int) {
	d.m = m
	s := make(*Stack, m)
	s.bot() = top[0], bot[0]
	s.top() = top[1], bot[1]
}

func (d *DblStack) isEmpty() {
	return top[0] == bot[0] && top[1] == bot[1]
}

func (d *DblStack) isFull() {
	return top[0] == top[1]
}

// the second param: 0 for left , 1 for right
func (d *DblStack) push(i, dir int) {
	if dir > 1 || dir < 0 {
		fmt.Errorf("wrong dir value")
		return
	}
	if d.isFull() {
		fmt.Errorf("DblSatck is full")
		return
	}

	if dir == 0 {
		*top[0] = i
		top[0]++
	} else {
		*top[1] = i
		top[1]--
	}
}

// 0 for left, 1 for right
func (d *DblSatck) pop(i, dir int) {
	if dir > 1 || dir < 0 {
		fmt.Errorf("wrong dir value")
		return
	}
	if d.isEmpty() {
		fmt.Errorf("DblSatck is empty")
		return
	}

	if dir == 0 {
		if top[0] != bot[0] {
			*top[0]--
		} else {
			fmt.Errorf("left stack is empty")
			return
		}
	} else {
		if top[1] != bot[1] {
			*top[1]++
		} else {
			fmt.Errorf("right stack is empty")
			return
		}
	}
}
