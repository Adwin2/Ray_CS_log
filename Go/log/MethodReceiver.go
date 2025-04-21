package main

// 错误实现
// func (c *ServerConfig) Load() {
//    c = ParseFlags() // 这个赋值只影响方法内的副本
// 方法内的c是外部指针的副本。只会影响副本，不会影响外部指针
// }

// main.go
//
//	func main() {
//		var conf *ServerConfig
//		conf.Load() // 调用后 conf 仍为 nil
//	}

// 如果需要保持 Load 方法
func (c **ServerConfig) Load() { // 使用双指针
	*c = ParseFlags()
}

// 调用方式
func main() {
	var conf *ServerConfig
	(&conf).Load() // 现在 conf 会被正确赋值
}

// 拓展：构造函数模式
func NewServerConfig() (*ServerConfig, error) {
	return ParseFlags(), nil
}

// 使用
// conf, _ := NewServerConfig()
