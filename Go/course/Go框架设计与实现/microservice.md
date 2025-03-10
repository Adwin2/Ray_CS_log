微服务架构是当前大多数互联网公司的标准架构

## 微服务架构中的...
**核心要素**
服务治理：服务注册 发现 负载均衡 扩缩容 流量治理 稳定性治理……
可观测性：日志采集 分析 监控打点 大盘 异常报警 链路追踪
安全：身份验证 认证授权 访问令牌 审计 传输加密 ……
## 基本概念
服务service  具有相同逻辑的运行实体
实例instance 每个运行实体
一个实例对应一或多个进程 
集群cluster 指服务内部的逻辑划分，包含多个实例
常见的实例承载形式：进程、VM、k8s pod
有\无状态服务 －－服务的实例是都存储乐可持久化的数据

- 服务间通信  ）单体服务 模块通信 函数调用 ）微服务 网络传输(HTTP gRPC Thrift)
- 服务注册及发现 ）ip:port❌目标服务的地址可能是动态的 ）DNS❌本地dns缓存导致延时 负载均衡问题 不支持服务实例的探活检查 无法配置端口 
解决思路：增加统一的服务注册中心 存储服务名到服务实例的映射
- 服务实例的上下线 上线-健康检查（health check is running all the time)

**流量特征**
- 统一网关入口
- 内网(intranet)通信采用RPC
- 网状调用链路

## 核心服务治理功能
**服务发布(deployment)**--升级运行新的代码 ）在线服务
难点：服务不可用(halt) 抖动(errors) 回滚
- 蓝绿部署 分为蓝绿集群 依次切换流量并升级  ）简单稳定 但需要两倍资源 宜流量低的时候
- 灰度发布（金丝雀发布）单个依次发布 ❌回滚复杂（精细化回滚 k8s..平台能力 ） 多次改变服务注册中心
# 流量治理
基于地区 集群 实例 请求.. 对端到端流量路由路径精确治理
# 负载均衡(Load Balance)
常见lb策略 ) Round Robin) Random) Ring hash) Least Request
# 稳定性治理
限流 rate limit （QPS限制）
熔断 
过载保护
降级 important -common

## 重试策略实践
Retry 避免偶发错误 提高sla (Service-level Agreement)
降低错误率 长尾延时（，重试有机会提前返回） 容忍暂时性错误 避开下游故障实例
- 难点：幂等性 重试风暴）设定重试比例；防止链路重试“请求失败，但别重试”status
理想情况下，只有最下一层重试

Hedged requests 先向下游发送request等待先到达的响应 （重试组件）

） 超时设置
