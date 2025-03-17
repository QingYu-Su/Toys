# 项目说明

- [gee-web](/gee-web/)：一个简易的Web框架，取名为gee，具有动态路由，分组控制，中间件，HTML模板和错误处理功能，重点参考了[Gin框架](https://github.com/gin-gonic/gin)。教学地址：https://geektutu.com/post/gee.html
- [gee-cache](/gee-cache/)：一个简易的分布式缓存框架，模仿了[groupcache](https://github.com/golang/groupcache)的实现，具有LRU淘汰策略，单机并发缓存，基于http的分布式缓存和一致性哈希选择节点实现负载均衡，使用锁机制防止缓存击穿，使用protobuf优化节点间通信。教学地址：https://geektutu.com/post/geecache.html
- [gee-orm](/gee-orm/)：一个简易的orm框架，orm即对象关系映射（Object Relational Mapping，简称ORM），是通过类名表示数据库表名，对象表示数据库记录，对象属性表示数据库字段的一种映射。本框架模仿了[xorm](https://github.com/go-xorm/xorm)的实现，具有基本的SQL操作，包括表的创建删除和迁移，记录的增删改查，查询条件的链式操作，自定义字段约束，利用钩子函数使得用户可以在执行操作前后自定义行为，支持事务。教学地址：https://geektutu.com/post/geeorm.html
- [gee-rpc](/gee-rpc/)：一个简易的rpc框架，即Remote Procedure Call（远程过程调用），用户可以像调用本地函数一样调用服务器上的服务。本框架选择从零实现 Go 语言官方的标准库 net/rpc，具有协议交换，消息编码，支持异步和同步调用的客户端，可以注册任意结构体的任意方法的服务注册功能，超时处理、支持将HTTP协议转为自定义的rpc协议，具有负载均衡能力的客户端，以及一个独立于服务端和客户端的服务发现和注册中心。教学地址：https://geektutu.com/post/geerpc.html