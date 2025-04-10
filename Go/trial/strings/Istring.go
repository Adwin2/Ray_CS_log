package main

//map[string]func(val)val

type Istring interface {
	// 连接字符串切片
	Join(elems []string, sep string) string

	// 重复一个字符串多次
	Repeat(s string, cnt int) string

	// 替换旧字符为新字符
	Replace()

	// 判断字符串是否包含子字符串
	Contains()

	// 获取子字符串在字符串中的索引位置
	Index()

	// 判断字符串是否以指定前缀开头
	HasPrefix()

	// ...指定后缀结尾
	HasSuffix()

	// 分割字符串为切片
	Split()

	// 字符串转为小写
	ToLower()

	// 字符串转为大写
	ToUpper()

	// 修剪字符串两端的指定字符
	Trim()

	// 比较字符串大小
	Compare()
}

// func main() {
// 	fmt.Println("vim-go")
// }
