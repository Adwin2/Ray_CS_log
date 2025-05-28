<h1>构造字符串的几种方法性能比较</h1>

<h2>1) fmt.Sprintf()</h2>

>  使用`strconv`系列函数优化类型转换

涉及反射和动态类型解析，频繁调用会产生额外开销

<h2>2) “+”拼接</h2>

类型限制：配合`strconv.Itoa`等方法使用

**Note: 1,2都会生成新的字符串，特定场景内存多次分配**

<h2>3)strings.Builder</h2>

> 高频拼接时推荐

- 通过与分配内存和复用缓冲区，避免重复内存复制
- Write--方法拼接任意类型
- String方法返回底层字节切片的字符串形式

<h2>4)[ ]byte预分配切片</h2>

- 手动控制内存分配，适合已知字符串最终长度的场景（极致性能）
- 一次分配足够容量，避免扩容开销

<h2>Final : 通用方案</h2>

```go
func buildString(name string, age int) string {
	var builder strings.Builder
    builder.Grow(len(name) + 16) // 预分配估算长度
    builder.WriteString("name:")
    builder.WriteString(name)
    builder.WriteString(", age:")
    builder.WriteString(strconv.Itoa(age))
    return builder.String()
}
```



