package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Augu", 9: "Sept", 10: "Oct", 11: "Nov", 12: "Dec"}

	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Println(Q2)
	fmt.Println(summer)

	//	fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend the Slice
	fmt.Println(endlessSummer)  // 底层共享内存 O(1)

	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} //rotate s by two positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)

	//slice之间不能使用==符号比较（函数如下
	//(唯一：[]byte类型时的bytes.Equal函数
	if s == nil { /* */
	}
	/*
		var s0 []int    / len(s) == 0 s == nil
		s0 = nil        /len(s) == 0 s == nil
		s0 = []int(nil) /len(s) == 0 s == nil
		s0 = []int{}    /len(s) == 0 s != nil
		:len(s) == 0测试是否为空 补:reverse(nil) 也是安全的
	*/

	/*	内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分省略时，容量等于长度
		len := 5
		cap := 6
		make([]T, len)        ←该slice为整个数组的view
		make([]T, len, cap)   /same as make([]T, cap)[:len]
		↑该slice是底层数组的前len长度
	*/

	/* append函数 ：用于向slice追加元素*/
	fmt.Println("\n4.2.1 append")

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	//append 函数容量够的时候返回地址不变 不够时扩展并重新分配地址 : runes = append(runes, r)

	var stack []int
	v := 0
	//fmt.Println("Slice模拟Stack")
	stack = append(stack, v) //push v
	//top := stack[len(stack)-1]   // top of stack
	stack = stack[:len(stack)-1] // pop

	//删除slice中间元素而保持顺序不变
	s2 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s2, 2))
	//不保持顺序就直接slice[i] = slice[len(slice)-1] pop即可

}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y ...int) []int {
	var z []int
	//zlen := len(x) + len(y)
	copy(z[len(x):], y) //copy(dst, src) 默认返回成功复制的元素个数 等于两个slice中较小的长度
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
