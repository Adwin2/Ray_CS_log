/**
	合并有序链表 I 、II、遍历找出链表最大值、O(1)反转链表
**/

package main

type ListNode struct {
	val  int
	next *ListNode
}

// 2.1 合并有序链表
func MergeList(pa *ListNode, pb *ListNode) *ListNode {
	//空节点作为起点
	ept := &ListNode{}
	cur := ept

	for pa != nil && pb != nil {
		if pa.val <= pb.val {
			if pa.val == pb.val {
				pb = pb.next
			}
			cur.next = pa
			pa = pa.next
		} else {
			cur.next = pb
			pb = pb.next
		}
		cur = cur.next
	}

	if pa != nil {
		cur.next = pa
	} else {
		cur.next = pb
	}
	return ept.next
}

// 2.2 合并有序链表II
func MergeListII(pa *ListNode, pb *ListNode) *ListNode {
	//空节点作为起点
	ept := &ListNode{}
	cur := ept

	for pa != nil && pb != nil {
		if pa.val >= pb.val {
			if pa.val == pb.val {
				pb = pb.next
			}
			cur.next = pa
			pa = pa.next
		} else {
			cur.next = pb
			pb = pb.next
		}
		cur = cur.next
	}

	if pa != nil {
		cur.next = pa
	} else {
		cur.next = pb
	}
	return reverseList(ept.next)
}

// 2.3 遍历确定最大值
func MaxVal(p *ListNode) int {
	var max int = p.val
	for p != nil {
		if p.val > max {
			max = p.val
		}
		p = p.next
	}
	return max
}

// 2.4 翻转链表 空间 O(1) 双指针法
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.next
		head.next = prev
		prev = head
		head = tmp
	}
	return prev
}
