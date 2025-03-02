# 50 Shades of Go : Traps, Gotchas, and Cpmmon Mistakes for New Golang Devs

( **Edited!** )

## 初级

1. 不允许`{`单独一行
2. 短变量声明只能在函数内部使用
3. 短变量声明不能用来设置字段值
4. 同名短变量声明在不同作用域出现会导致幽灵变量现象 `go tool vet -shadow your_file.go` 检查幽灵变量
5. nil不能用来初始化未指定类型的变量
6. 字符串不允许使用nil值 。nil只能赋值给**指针，channel，func，interface，map，slice**
7. 不能直接使用nil的Slice和Map
8. map使用make分配内存时可指定capacity，但不可使用cap函数
9. 数组用于函数传参时是值传递，只有**map，slice，channel，指针**是引用传递

    ```go
    x := [3]int{1, 2, 3}

    func(arr *[3]int) {
        (*arr)[0] = 7
        fmt.Println(arr) // &[7, 2, 3]
    }(&x)
    fmt.Println(x) // [7, 2, 3]
    ```

10. `range`返回键值对， 默认 索引 + 值
11. `map[key]`始终有返回值，默认0
12. 字符串不可变

    ```go
    x := "text"

    xbytes := []byte(x)
    xbytes[0] = 'T'

    fmt.Println(string(xbytes))
    ```

13. 字符串与[]byte 之间的转换是复制（内存损耗），可用`map[string][]byte` 建立字符串与[]byte的映射，也可range来**避免内存分配**，提高性能

    ```go
    for i, v := range []byte(str) { ...
    }
    ```

14. string 索引操作返回的是byte(或uint8)，获取字符可用for range，也可使用`unicode/utf8`和`golang.org/x/exp/utf8string`包的`At()`方法
15. `len(str)`返回的是字符串的字节数，获取字符串的rune数通过`unicode/utf8.RuneCountInString()`函数，注意有些字符由多个rune组成(如é是两个rune组成)。
16. slice， array， map多行书写时最后的逗号不可省略
17. 内置数据结构的操作并不同步，但可以配合Go并发特性(goroutine channel)
18. `for .. range ..` 以rune类型遍历string。for range总是尝试将字符串解析成utf8的文本，对于它无法解析的字节，它会返回**oxfffd**的rune字符。因此，任何**包含非utf8**的文本，一定要先将其转换成字符切片([]byte)。
    > 一个字符，也可以有多个rune组成。需要处理字符，尽量使用`golang.org/x/text/unicode/norm`包。

    ```go
    data := "A\xfe\x02\xff\x04"
    for _,v := range data {
        fmt.Printf("%#x ",v)
    }
    //prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

    fmt.Println()
    for _,v := range []byte(data) {
        fmt.Printf("%#x ",v)
    }
    //prints: 0x41 0xfe 0x2 0xff 0x4 (good)
    ```

19. 使用`for .. range ..`遍历map 每次顺序是随机的。
20. switch case匹配规则：匹配条件后默认退出，除非使用`fallthrough`继续匹配；不同于其他语言依赖break退出。
21. Go只存在后置自增自减
22. 位运算的**非**操作是 ^(跟**异或**位运算符号一致)；不同于其他语言的 ~
23. 位运算（与、或、异或、取反）优先级高于四则运算（加、减、乘、除、取余），有别于C 。
24. struct在序列化时以小写字母开头的字段不会encode，decode时显示为0值。
25. 主程序结束即退出。可通过channel实现主协程等待goroutine完成。（或sync.WaitGroup）
26. 无缓存channel的阻塞问题

    ```go
    ch := make(chan int)

    var ch chan int // 此时channel值为 nil 同样会永远阻塞
    ```

27. 从closed的channel读取数据是安全的，可通过返回值的第二个参数判断是否关闭。而向closed写channel会导致panic
28. 方法接收者是指针类型（*T），是对原对象的引用，方法中对其修改就是对原对象修改。 否则只是值复制。
29. log包中的`log.Fatal`和`log.Panic`不仅仅记录日志，还会中止程序。它不同于Logging库。

## 中级

1. 关闭HTTP的Response.Body
   使用defer语句关闭资源时要注意nil值，在defer语句之前要进行nil值处理

    ```go
    package main

    import (
        "fmt"
        "net/http"
        "io/ioutil"
    )

    func main() {
        resp, err := http.Get("https://api.ipify.org?format=json") 

        if resp != nil {
            defer resp.Body.Close()
        }
        //defer位置： 1)nil判断之前： defer执行引发空引用的panic
        if err != nil {
        fmt.Println(err)
        return
        }

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println(string(body))
    }
    ```

    > 在Go 1.5之前resp.Body.Close()会读取并丢失body中的数据，保证在启用keepaliva的http时能够在下一次请求时重用。
    在Go 1.5之后，就需要在关闭前手动处理。
    _, err = io.Copy(ioutil.Discard, resp.Body)  
    如果只是读取Body的部分，就很有必要在关闭Body之前做这种手动处理。例如处理json api响应时json.NewDecoder(resp.Body).Decode(&data)就需要处理掉剩余的数据。

2. 关闭HTTP连接：
3. Json反序列化数字到interface{}类型的值中，默认解析为float64，使用时注意。
4. Struct、Array、Slice、Map **的**比较
   - Struct 和 Array在所有元素都可比较时才可以比较；否则编译错误
   - Go提供了一些用于比较不能直接==比较的函数
     reflect.DeepEqual() 对于**nil值**的Slice与**空元素**的Slice不相等；这点不同于bytes.Equal()函数。
   - 忽略大小写来比较包含文字数据的字节切片
     strings.EqualFold() , bytes.EqualFold()  //ToUpper() ToLower() 只能处理英文文字
   - 如果要比较用于验证用户数据密钥信息的字节切片时，使用reflect.DeepEqual()、bytes.Equal()、bytes.Compare()会使应用程序遭受计时攻击(Timing Attack)，可使用crypto/subtle.ConstantTimeCompare()避免泄漏时间信息。
5. 从panic中恢复
   recover()函数可以捕获、拦截panic，必须在defer函数/语句中直接调用
6. 在Slice、Array、Map的`for .. range ..`子句中修改和引用数据项
   使用range获取的数据是从集合元素中复制过来的，并非原始数据（语法糖），但使用索引可以访问原始数据

    ```go
    data := []int{1,2,3}
    for _, v := range data {
        v *= 10  // 不改变原数据
    }

    data2 := []int{1,2,3}
    for i, v := range data2 {
        data[i] *= 10  // 改变原数据
    }
    //拓：
    data3 := []*struct{num int} {{1},{2},{3}}
    for _, v := range data3 {
        v.num *= 10
    }
    fmt.Println(*data3[0], *data3[1], *data3[2])
    ```

7. Slice中的隐藏数据
   从Slice上生成切片新的Slice，新slice会直接引用原始数组，两个slice对同一数组的操作会相互影响。可通过手动分配空间来避免相互影响。
8. Slice超范围数据覆盖
   新生成切片之间capicity区域是重叠的，因此在添加数据时易造成数据覆盖问题。
    slice使用append添加的内容时超出capicity时，会重新分配空间。
    利用这一点，将要修改的切片指定capicity为切片当前length，可避免切片之间的超范围覆盖影响。

   ```go
    path := []byte("AAAA/BBBBBBBBB")
    sepIndex := bytes.IndexByte(path,'/') //bytes.IndexByte(str, char)
    dir1 := path[:sepIndex]
    // 解决方法
    // dir1 := path[:sepIndex:sepIndex] //full slice expression
    dir2 := path[sepIndex+1:]
    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAA
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => BBBBBBBBB

    dir1 = append(dir1,"suffix"...)
    path = bytes.Join([][]byte{dir1,dir2},[]byte{'/'})

    fmt.Println("dir1 =>",string(dir1)) //prints: dir1 => AAAAsuffix
    fmt.Println("dir2 =>",string(dir2)) //prints: dir2 => uffixBBBB (not ok)

    fmt.Println("new path =>",string(path))   
   ```

9. Slice增加元素重新分配内存导致的怪事
   slice在添加元素前，与其它切片共享同一数据区域，修改会相互影响；但添加元素导致内存重新分配之后，不再指向原来的数据区域，修改元素，不再影响其它切片。
10. 类型重定义与方法继承
    从一个已存在的(non-interface)非接口类型重新定义一个新类型时，不会继承原类型的任何方法。
    可以通过定义一个组合匿名变量的类型，来实现对此匿名变量类型的继承。（类似适配器模式）
    但是从一个已存在接口重新定义一个新接口时，新接口会继承原接口所有方法。
11. 从“for switch/select”代码块中跳出
    无label的break只会跳出最内层的switch/select代码块。
    如需要从switch/select代码块中跳出外层的for循环，可以在for循环外部定义label，供break跳出。

    return当然也是可以的，如果在这里可以用的话。

    ```go
    //Go语言中默认的break语句只能终止当前最内层的switch/select代码块，无法直接跳出外层的for循环。例如：

    for {
        switch val := someFunc(); val {
        case 1:
            break  // 仅跳出switch，循环继续执行
        case 2:
            // do something
        }
    }
    //标签break的用法
    //通过在for循环外定义标签，配合break + 标签名实现跨层跳出：

    OuterLoop:  // 定义标签
    for i := 0; i < 5; i++ {
        switch {
        case i == 2:
            break OuterLoop // 直接跳出整个for循环
        default:
            fmt.Println(i)
        }
    }
    // 输出: 0 1
    //若跳出后无需执行后续逻辑，可以直接用return退出函数：

    func process() {
        for {
            switch {
            case condition:
                return // 直接退出函数
            }
        }
        // 后续代码不会执行
    }
    ```

12. 在for迭代过程中，迭代变量会一直保留，只是每次迭代值不一样。
    因此在for循环中在闭包里直接引用迭代变量，在执行时直接取迭代变量的值，而不是闭包所在迭代的变量值。

    如果闭包要取所在迭代变量的值，就需要for中定义一个变量来保存所在迭代的值，或者通过闭包函数传参。

    ```go
    package main

    import (  
        "fmt"
        "time"
    )

    func forState1(){
        data := []string{"one","two","three"}

        for _,v := range data {
            go func() {
                fmt.Println(v)
            }()
        }
        time.Sleep(3 * time.Second)    //goroutines print: three, three, three

        for _,v := range data {
            vcopy := v // 使用临时变量
            go func() {
                fmt.Println(vcopy)
            }()
        }
        time.Sleep(3 * time.Second)    //goroutines print: one, two, three

        for _,v := range data {
            go func(in string) {
                fmt.Println(in)
            }(v)
        }
        time.Sleep(3 * time.Second)    //goroutines print: one, two, three
    }

    func main() {  
        forState1()
    }
    ```

再看一个例子

    ```go
    package main

    import (  
        "fmt"
        "time"
    )

    type field struct {  
        name string
    }

    func (p *field) print() {  
        fmt.Println(p.name)
    }

    func main() {  
        data := []field{{"one"},{"two"},{"three"}}
        for _,v := range data {
            // 解决办法：添加如下语句
            // v := v
            go v.print()
        }
        time.Sleep(3 * time.Second)     //goroutines print: three, three, three

        data2 := []*field{{"one"}, {"two"}, {"three"}}  // 注意data2是指针数组
        for _, v := range data2 {
            go v.print()                // go执行是函数，函数执行之前，函数的接受对象已经传过来
        }
        time.Sleep(3 * time.Second)     //goroutines print: one, two, three
    }
    ```
13. defer函数调用参数
defer后不论函数还是方法，输入参数的值在defer声明时已计算好
要特别注意的是，defer后面是方法调用语句时，方法的接受者是在**defer语句执行时**传递的，而不是defer声明时传入的。

    ```go
    type field struct{
        num int
    }
    func(t *field) print(n int){
        fmt.println(t.num, n)
    }
    func main() {    
        var i int = 1
        defer fmt.Println("result2 =>",func() int { return i * 2 }())
        i++

        v := field{1}
        defer v.print(func() int { return i * 2 }())
        v = field{2}
        i++

        // prints: 
        // 2 4
        // result => 2 (not ok if you expected 4)
    }
    ```

14. defer在当前函数结束后调用，与变量的作用范围无关
15. 类型断言失败时会返回T类型的“0值”，而不是变量原始值。

    ```go
    var data interface{} = "great"

    if data, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",data)
    } else {
        fmt.Println("[not an int] value =>",data)         //prints: [not an int] value => 0 (not "great")
    }

    if res, ok := data.(int); ok {
        fmt.Println("[is an int] value =>",res)
    } else {
        fmt.Println("[not an int] value =>",data)         //prints: [not an int] value => great (as expected)
    }
    ```
16. 阻塞的goroutine与资源泄露
    ```go
    func First(query string, replicas ...Search) Result {  
        c := make(chan Result)
        // 解决1：使用缓冲的channel： c := make(chan Result,len(replicas))
        searchReplica := func(i int) { c <- replicas[i](query) }
        // 解决2：使用select-default，防止阻塞
        // searchReplica := func(i int) {
        //     select {
        //     case c <- replicas[i](query):
        //     default:
        //     }
        // }
        // 解决3：使用特殊的channel来中断原有工作
        // done := make(chan struct{})
        // defer close(done)
        // searchReplica := func(i int) {
        //     select {
        //     case c <- replicas[i](query):
        //     case <- done:
        //     }
        // }
        for i := range replicas {
            go searchReplica(i)
        }
        return <-c
    }
    ```

## 高级

1. 用值实例上调用接收者为指针的方法
对于可寻址(addressable)的值变量(而不是指针)，可以直接调用接受对象为指针类型的方法。
换句话说，就不需要为可寻址值变量定义以接受对象为值类型的方法了。

但是，并不是所有变量都是可寻址的，像Map的元素就是不可寻址的。

```go
package main

import "fmt"

type data struct {  
    name string
}

func (p *data) print() {  
    fmt.Println("name:",p.name)
}

type printer interface {  
    print()
}

func main() {  
    d1 := data{"one"}
    d1.print() //ok

    // var in printer = data{"two"} //error
    var in printer = &data{"two"}
    in.print()

    m := map[string]data {"x":data{"three"}}
    //m["x"].print() //error
    d2 = m["x"]
    d2.print()      // ok
}
```
