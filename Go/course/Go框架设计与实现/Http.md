请求解析  后端路由  业务逻辑
超文本传输协议 Hypertext Transfer Protocol
      考虑：明确的边界；携带什么信息
POST请求组成  ：
请求行/状态行（方法名 协议版本 URL）/协议版本 状态码 状态码描述
请求头/响应头
请求体/响应体
请求流程
HTTP1 HTTP2 QUIC
高内聚、低耦合、易复用、易扩展

应用层设计 
- 提供合理API 名称要易于理解  ； 简单 可见 可测 兼容
中间件层设计
- 配合Handler实现一个完整的请求处理生命周期（调用链）； 拥有预处理逻辑与后处理逻辑； 可注册多中间件； 对上层模块用户逻辑模块易用
路由设计
- 框架路由 为URL匹配对应handler  前缀匹配树
协议层
- 抽象出合适的接口  在连接上读写数据
传输层
BIO block io ； NIO先判断readable再唤醒goroutine （netpoll）


## 性能优化
**针对网络库**
*Buffer设计*
存下全部Header  减少系统调用次数  能够复用内存  能够多次读
Netpoll：存所有Header 拷贝完整Body
 
- go net 适用于流式 小包
- netpoll 中大包友好 时延低

**针对协议**
*Header解析*
找到Header Line 边界 ` \r\n `出现两次 即结束 
> 先找到\n 在看前一个是不是\r 
SIMD 单指令集多数据 sonic加速编解码
核心字段快速解析 使用byte slice存储 高频使用存储为额外的成员变量

*热点资源池化*
RequestContext池 
syncpool

## 企业实践
追求性能
追求易用(API快速上手)，减少误用
打通内部生态
文档建设、用户群建设

内部http框架：Hertz
1万+服务 3千万+QPS
