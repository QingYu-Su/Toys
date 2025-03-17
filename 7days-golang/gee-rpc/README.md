# 项目说明

项目教学地址：[7天用Go从零实现RPC框架GeeRPC](https://geektutu.com/post/geerpc.html)

## 第1天
1. 在codec包中抽象出请求和响应共有的部分，即header和body，header包括服务方法，请求序号和错误（客户端为空），body在客户端时为参数，在服务端时为响应值。
2. 在codec包中抽象出对消息体（即请求和响应的组合）的编解码接口Codec，其中定义了从流中读取header、读取body和写入消息（同时写入header和body）的方法。方便不同的协议实现不同的编解码。
3. 在codec包中抽象出Codec的构造函数，可以通过Codec的Type来得到构造函数。目前实现了Gob一种编解码体，即GobCodec。
4. 客户端与服务端的通信需要协商一些内容，对于 GeeRPC 来说，目前需要协商的唯一一项内容是消息的编解码方式。我们将这部分信息，放到sever.go中的结构体 Option 中承载。Option中包含MagicNumber和CodecType，这里固定使用json进行编解码，魔数表示协议同步，CodecType表示后续消息体的编解码方式。在单次连接中，Option固定在报文最开始，后续的消息可以有多个，但都是用固定的CodecType进行编解码了，即Option-Header-Body-Header-Body。。。
5. 在server包中定义了server结构体，实现了accept方法，只需要传入net.listener，即告诉是在哪个端口进行监听即可。每次监听到一个连接，则开启一个go程去处理，使得服务端可以在一个端口监听多个连接。
6. 在server结构体中还实现了ServeConn和serveCodec，ServeConn解析Option然后交由特定的编解码接口，serveCodec解析消息，每次成功解析到一个消息，则开启一个go程处理。这里需要注意的是，虽然一次连接可以接收多个请求，但回复请求的报文必须是逐个发送的，并发容易导致多个回复报文交织在一起，客户端无法解析。在这里使用锁(sending)保证。并且，所有请求都处理完成后，才能关闭连接，这里使用WaitGroup保证。
7. 在server包中定义了request结构体，包括一个header、一个请求参数和一个响应值。在读取消息时，会读取header和请求参数，在响应时，只修改响应值。这些过程交由server结构体的readRequest、sendResponse和handleRequest方法处理。

## 第2天
1. 在client.go中实现了结构体Call，表示一次RPC请求所需的信息，包括请求参数和响应，以及一个*Call类型的无缓冲通道Done，只要Call的任务完成，就会往通道写入已经完成了的Call，使得支持异步调用Call。
2. 在client.go中实现了Client结构体和方法，支持同步（Call）和异步方式（Go）使用Call，并且保证了请求方式是串行执行的。
3. 创建Clinet后，会自动开启一个协程不断去接收响应，并将对应响应写入Call然后调用Done。

## 第3天
1. 构建service.go，将结构体的方法映射为服务，可以使得服务器注册任意结构体的任意方法，因此需要使用反射去实现。
2. 在service.go定义结构体methodType，它包含一个具体方法的完整信息，包括方法本身，第一个参数的类型（请求参数），第二个参数的类型（响应参数），以及统计多少次调用的值。实现了 2 个方法 newArgv 和 newReplyv，用于创建对应类型的实例。
3. 在service.go定义结构体 service，它包含一个具体类的完整信息，包括名称，结构体类型，结构体本身，以及一个map存储结构体的所有符合条件的方法。它实现了在创建时根据结构体初始化，然后还有一个重要的方法Call用来调用特定方法。注意，在使用反射获取一个结构体函数时，如果要调用，第一个参数必定是结构体本身。
4. 将service集成到服务器，具有注册和发现服务的方法，并且在获取请求和处理请求时，也要使用到服务。

## 第4天
1. 给服务端和客户端添加超时处理机制，总的来说，可以通过给某段代码使用go程处理，然后完成之后往通道写入结果，然后靠select一个time.After(timeout)和结果通道去看是否超时。

## 第5天
1. 使得gee-rpc支持http协议，重点在于协议同步，首先客户端发起HTTP CONNECT请求到服务端，服务端接收后返回一个ok，然后将底层的TCP连接转化为rpc连接，客户端接收ok后，也转换为rpc连接
2. 支持 HTTP 协议的好处在于，RPC 服务仅仅使用了监听端口的 /_geerpc 路径，在其他路径上我们可以提供诸如日志、统计等更为丰富的功能。

## 第6天
1. 创建xclient包，构建了服务发现dicovery和支持负载均衡的客户端xclient
2. dicovery是一个接口，包括了服务发现最基本的几个方法，刷新，更新，获取和获取全部服务地址。并且实现了一个不需要注册中心，服务列表由手工维护的结构体MultiServersDiscovery。
3. XClient是一个支持负载均衡的客户端，其中包含多个与服务端通信的客户端，在执行call命令时，会选择一个客户端进行通信。

## 第7天
1. 添加具有心跳机制的注册中心registry，使用HTTP服务端为承载，是rpc客户端和服务端的中间件
2. 在xclient中添加GeeRegistryDiscovery，通过注册中心获取服务列表