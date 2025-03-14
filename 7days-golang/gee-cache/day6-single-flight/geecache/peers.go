package geecache

// PeerPicker 是一个接口，用于选择一个合适的节点（Peer）来获取缓存数据。
// 它定义了一个方法 PickPeer，该方法根据给定的键选择一个节点。
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 是一个接口，用于从其他节点获取缓存数据。
// 它定义了一个方法 Get，该方法根据缓存组名和键从远程节点获取数据。
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
