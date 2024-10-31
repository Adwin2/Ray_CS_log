package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建
	//ages0 := make(map[string]int) //mapping from strings to ints

	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	/*
		/same as↓
		ages1_1 := map[string]int
		ages1_1["alice"] = 31
		ages1_1["charlie"] = 34
		ages2 := map[string]int{}
	*/

	//访问
	//通过key对应的下标语法访问
	ages1["alice"] = 32
	fmt.Println(ages1["alice"]) //"32"

	delete(ages1, "alice") //remove element ages["alice"]

	//map元素不支持取址 容量变化时会发生变化   、可以对单个元素操作

	//key索引相关
	age, ok := ages1["bob"]
	if !ok {
		/*
			bob is not a key in this map;
			age ==0
		*/
		fmt.Println(age, "not a key")
	}
	/*
		or integrate 'em like
		if age, ok := ages1["bob"]; !ok {
			/...
		}
	*/

	//只能和nil进行相等比较
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))

	//value可以是map或slice类的聚合类型
	var graph = make(map[string]map[string]bool)
	//addEdge 惰性初始化map：在每个值首次作为key时才初始化。hasEdge使map的零值也能正常工作：from，to不存在时也可以返回正常结果

	/*
		忽略value(bool类型)的map来当做一个字符串集合
		以下通过map来表示所有的输入行所对应的set集合，以确保已经在集合存在的行不会被重复打印
	*/
	seen := make(map[string]bool) //a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
		if line == "stop" {
			os.Exit(1)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	//false : len不同、对应key不存在、key存在value不同
	return true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
