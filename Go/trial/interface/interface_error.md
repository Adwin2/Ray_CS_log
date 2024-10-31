error 类型
```go
type error interface{
	Error() string
}
```
errors.New函数  根据传入的错误信息返回一个新error（独特的指针类型
常用其封装函数 fmt.Errorf还可以处理字符串格式化
扩展：syscall包提供了go语言底层api 实现error接口的数字类型Errno
Errno的error方法会从一个字符串表中查找错误信息
