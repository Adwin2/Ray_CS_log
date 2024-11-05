##### Chapter9 基于共享变量的并发

### 9.1 竞争条件
> *对于两个goroutine中的事件x和y 无法确认先后发生顺序 那么说事件x和y是并发的*
> -  一个函数在线性程序中可以正确工作 在并发情况下依旧可以正确工作的情况下 称该函数是并发安全的
> 并发安全的函数 不需要额外的同步工作 
> - 一个类型可访问的所有方法和操作函数都是并发安全的话 称该类型是并发安全的
> 并发访问类型的方法是危险的行为
> 但包导出函数一般都是并发安全的 （package级别的变量无法被限制在单一的goroutine 修改必须使用互斥条件
**竞争条件指的是程序在多个goroutine交叉执行操作时出现错误。**
特殊情况才会触发 难以复现和分析诊断

两个goroutine并发访问同一个变量，且其中至少一个写操作时候会发生数据竞争 竞争对象是比一个机器字更大的类型时更麻烦
###有三种方式可以避免数据竞争：

## 第一种方法是不要去写变量。
- 考虑一下下面的map，会被“懒”填充，也就是说在每个key被第一次请求到的时候才会去填值。如果Icon是被顺序调用的话，这个程序会工作很正常，但如果Icon被并发调用，那么对于这个map来说就会存在数据竞争。
```go
var icons = make(map[string]image.Image)
	func loadIcon(name string) image.Image

	// NOTE: not concurrency-safe!
	func Icon(name string) image.Image {
		    icon, ok := icons[name]
				    if !ok {
						icon = loadIcon(name)
						icons[name] = icon
					} //顺序调用时填值可以正确运行  并发则未知
			return icon
	}
```
在goroutine之前初始化完毕 且不会修改 那么任意数量访问都是安全的（goroutine只是读取操作）-不需要update的时候使用

## 第二种 避免多个goroutine访问变量  
- （把特定变量限制在特定goroutine中）
其它goroutine只能使用一个channel来发送请求给指定的goroutine查询更新变量。  即所谓的“通过channel通信来共享数据”。 该channel所在goroutine称为monitor goroutine
多人存钱问题 （原 俩goroutine deposit balance返回未知情况
通过第二种方法 把 balance放在monitor goroutine中
```go
package bank

var deposits = make(chan int) //send amount to balances
var balances = make(chan int) // receive balances

func Deposit(amount int) {
	deposits <-amount
}
func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select{
			case amount := <-deposit:
				balance += amount
			case balances <- balance:
		}
	}
}

func init() {
	go teller() //start monitor goroutine
}
```
对于变量会被goroutine严格的顺序访问现象称作 串行绑定
```go
type Cake stuct {
	state string
}

func baker (cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}
```

## 第三种 是允许多个goroutine访问变量 但同时最多只有一个访问  --- 
**互斥**

- **sync.Mutex互斥锁**

将需要锁的函数分成两部分：导出（首字母大写）和不导出（小写）并加注释 且二者不可相互访问
避免重入锁现象（Go不允许）
示例 银行Withdraw
```go
import "sync"
var(
		mu sync.Mutex //guards balance
		balance int
	)

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func Deposit (amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock() //确保访问变量不变
	defer mu.Unlock()
	return balance
}

//This function requires lock
func deposit(amount int) {
	balance += amount
}

```

- **sync.RWMutex读写锁**
“多读单写锁” 允许多个只读操作并行执行 但写操作会完全互斥
如前例的 Banlance函数 只读操作申请互斥锁 多次操作影响效率
```go
var mu sync.RWMutex
var balance int
func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
```
RLock只能在临界区无任何写操作时可用 且慎用 要比无竞争锁的mutex慢

### 内存同步

多核cpu中，并发运行时，在编译器编译后x,y变量有可能是在两个独立的CPU上都有副本（local cache）的，并且此时是被初始化为0的，再之后，由于编译器认为A1,A2这两条语句和B1,B2这两条语句的顺序不影响结果，就有可能调换两者的次序。从而打印出x:0 y:0，y:0 x:0
- 缺少显式的内存同步，编译器和CPU可以随意更改访问内存的指令顺序

### sync.Once惰性初始化

初始化成本高时 延迟到需要的时候再做

对于提过的 顺序调用时填值的 Icon函数  并发执行初始化会产生未知行为
互斥访问icons又没办法并发访问 再引入允许多读的锁。。。
⬇️
解决这种一次性初始化问题 -- sync.once的Do方法
概念上来讲，一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了；互斥量用来保护boolean变量和客户端数据结构。Do这个唯一的方法需要接收初始化函数作为其参数。
```go

var loadIconsOnce sync.Once
var icons map[string]image.Image
// Concurrency-safe.
func Icon(name string) image.Image {

	    loadIconsOnce.Do(loadIcons)
		return icons[name]
}
```
调用Do时锁定mutex, 并检查boolean值 false则初始化

### 竞争条件检测
build test run 等命令后加上-race的flag
竞争检查器会报告所有的已经发生的数据竞争。然而，它只能检测到运行时的竞争条件；并不能证明之后不会发生数据竞争。所以为了使结果尽量正确，请保证你的测试并发地覆盖到了你的包。

由于需要额外的记录，因此构建时加了竞争检测的程序跑起来会慢一些，且需要更大的内存，即使是这样，这些代价对于很多生产环境的程序（工作）来说还是可以接受的。对于一些偶发的竞争条件来说，让竞争检查器来干活可以节省无数日夜的debugging。

