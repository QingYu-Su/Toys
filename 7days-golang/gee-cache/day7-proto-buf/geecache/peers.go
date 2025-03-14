package geecache

import pb "geecache/geecachepb"

// PeerPicker 是一个接口，用于选择一个合适的节点（Peer）来获取缓存数据。
// 它定义了一个方法 PickPeer，该方法根据给定的键选择一个节点。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 是一个接口，用于从远程节点获取缓存数据。
// 它定义了一个 Get 方法，该方法接受一个请求消息和一个响应消息，
// 并通过网络从远程节点获取数据。
type PeerGetter interface {
	// Get 方法从远程节点获取缓存数据。
	// 输入参数：
	//   - in: 指向 pb.Request 的指针，表示请求消息，通常包含缓存组名和键。
	//   - out: 指向 pb.Response 的指针，表示响应消息，用于存储获取到的缓存数据。
	// 返回值：
	//   - error: 表示操作是否成功。如果返回 nil，表示成功；否则返回错误。
	Get(in *pb.Request, out *pb.Response) error
}
