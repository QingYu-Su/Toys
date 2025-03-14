# 项目说明

项目教学地址：[7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee.html)

## 第1天
1. 实现lru缓存淘汰策略，使用字典和双向链表实现，该策略存放在lru包内
2. 链表节点的值包含了字典的key和value，方便在删除时同时删除字典
3. 节点值的value是一个接口，只要实现了Len方法即可
4. lru cache实现了获取、淘汰、增加和修改等功能，在增加和修改时，若缓存大小超过了最大字节，则会自动触发淘汰机制，因此外部可以只关心获取、增加和修改。
5. 可以从外部设置回调函数，使得在节点删除时做后续处理

## 第2天
1. 使用ByteView来封装缓存值，它实现了Len方法，本质上只是一个字节切片，并且具有只读属性，外部无法修改其值
2. 实现了cache，是对lru cache的封装，其封装点在于，增加的缓存值只能是ByteView，并且增加了锁，使其线程安全。因为我们每次获取、增加和修改时，都可能导致lru cache底层的链表和字典发生变化。
3. 实现了geecache，用来封装cache，是与外界交互的核心。其中，第一实现了回调Getter，由用户设置，表示缓存不存在时，应该如何获取缓存数据。第二实现了Group，它代表了一组缓存，其中包含了缓存组名称，缓存组Getter，缓存组的cache。
4. geecache在全局设置了groups和全局读写锁，我们可以通过名称获取一个缓存group，并且有全局读写锁来保证线程安全。
5. geecache中的group实现了几个方法，Get方法从 mainCache 中查找缓存，如果存在则返回缓存值。缓存不存在，则调用 load 方法，load 调用 getLocally（分布式场景下会调用 getFromPeer 从其他节点获取），getLocally 调用用户回调函数 g.getter.Get() 获取源数据，并且将源数据添加到缓存 mainCache 中（通过 populateCache 方法）

## 第3天
1. 实现了HTTPPool作为http服务端，用于给其他节点提供缓存值，它有两个值，一个是self，表示自己的地址，包括主机名/IP 和端口。另一个是 basePath，作为节点间通讯地址的前缀，默认是 /_geecache/。
2. 服务端实现逻辑非常简单，首先判断访问路径的前缀是否为 basePath，然后约定访问路径格式为` /<basepath>/<groupname>/<key>`，通过 groupname 得到 group 实例，再使用 group.Get(key) 获取缓存数据。最终使用 w.Write() 将缓存值作为 httpResponse 的 body 返回。

## 第4天
1. 实现了一致性哈希，对于给定的key，如果本节点没有缓存值，则需要从其他节点获取，并且每一次都选择同一节点。底层使用了哈希环，即将哈希值排序，并且首尾相连。
2. 为了解决节点过少导致的数据倾斜，一致性哈希还实现了虚拟节点，一个真实节点对应多个虚拟节点。
3. 一致性哈希命名为Map，提供了New方法，可由用户自定义一个真实节点对应的虚拟节点个数和哈希函数。然后是Add方法，允许提供多个真实节点的名称（string），然后增加对应的虚拟节点并映射为哈希值，存储到keys中，最后给keys排序。
4. 最后实现get方法，输入一个key值（striing），返回一个真实节点名称（string），底层使用了二分查找方法，所以需要在add方法中进行排序。并且通过给索引值取模，实现了环操作。

## 第5天
1. 实现了两个接口，一个是PeerPicker，必须实现PickPeer方法，通过给定的key值，给出特定的PeerGetter。然后就是PeerGetter，必须实现Get方法，通过给定的group名称和key值，给出缓存值
2. 实现了httpGetter作为客户端，其中baseURL 表示将要访问的远程节点的地址(比如https://ip+/_geecache)，实现了Get方法，因此是一个具体的PeerGetter。
3. 为HTTPPool，即服务端增加节点选择功能。具体为，新增成员变量 peers，类型是一致性哈希算法的 Map，用来根据具体的 key 选择节点。新增成员变量 httpGetters，映射远程节点与对应的 httpGetter。每一个远程节点对应一个 httpGetter，因为 httpGetter 与远程节点的地址 baseURL 有关。
4. HTTPPool实现了Set方法，可以传入任意数量的节点名称，实例化了一致性哈希算法，然后添加了传入的节点，并且针对每个节点，都创建了与之对应的httpGetter，然后存储在httpGetters中。每一个httpGetter中的baseurl为节点名称(通常为远程节点的地址)+服务端的basePath（通信前缀，这里默认为/_geecache）。
5. HTTPPool还实现了PickPeer方法，即通过给定的key值，给出特定的PeerGetter，这里给出的是httpGetter，因此HTTPPool是一个具体的PeerPicker。
6. 对Group增加新成员peers，它是一个PeekPicker。新增RegisterPeers方法，将 实现了PeerPicker接口的HTTPPool注入到 Group中。
7. Group新增 getFromPeer() 方法，它接受一个PeerGetter和key值，使用实现了 PeerGetter 接口的 httpGetter 从访问远程节点，获取缓存值。
8. Group修改 load 方法，使用 PickPeer() 方法选择节点，若非本机节点，则调用 getFromPeer() 从远程获取。若是本机节点或失败，则回退到 getLocally()。
9. 总之geecache实现的是面向用户的部分，即用户通过geecache查询key-value。http实现的面向服务器的部分，一个服务器有一个服务端和多个客户端，服务端负责接收查询，客户端负责发送查询。每次用户通过geecache查询时，默认从本地缓存查询，如果没有，则通过其他节点获取，如果没有，则通过用户自定义的数据获取器获取。

## 第6天
1. 为了避免短时间大量的并发请求，导致我们需要发大量的HTTP请求到其他节点，这里设计了一种防止缓存击穿的设计
2. 设计名为signleflight，即单次飞行。如果在服务端同时收到对同一时刻收到对同一个key的请求，则只会发起一次请求，然后阻塞其他请求，获取到的值由所有请求共享。
3. 设计了call，call代表一个正在执行或者是执行完成的请求。其中包括返回值、错误和waitgroup。waitgroup本质是一个计数器，允许多个协程等待。
4. 设计Group，表示一个缓存组中的singleflight，管理不同key的请求call，其中包含锁和map，map中映射了key和call，锁则是用来使map线程安全的。
5. Group实现了Do方法，接收 2 个参数，第一个参数是 key，第二个参数是一个函数 fn。Do 的作用就是，针对相同的 key，无论 Do 被调用多少次，函数 fn 都只会被调用一次，等待 fn 调用结束了，返回返回值或错误。
6. 然后将singleflight组合到geecache的Group中，在load方法中使用Do去包装到远程节点的请求，避免大量并发操作导致的过载。

## 第7天
1. 引入了protobuf来封装节点之间的通信。protobuf 即 Protocol Buffers，Google 开发的一种数据描述语言，是一种轻便高效的结构化数据存储格式，与语言、平台无关，可扩展可序列化。protobuf 以二进制方式存储，占用空间小。
2. 按照 protobuf 的语法，在 .proto 文件中定义数据结构，并使用 protoc 生成 Go 代码（.proto 文件是跨平台的，还可以生成 C、Java 等其他源码文件）。在项目代码中引用生成的 Go 代码。