package gee

import (
	"fmt"
	"strings"
)

// node 表示前缀树中的一个节点，用于存储路径模式的片段
type node struct {
	pattern  string  // 完整的路径模式，例如 "/user/:id"
	part     string  // 当前节点对应的路径片段，例如 "user" 或 ":id"
	children []*node // 子节点列表，表示路径的下一级片段
	iswild   bool    // 是否为通配符节点，通配符以 ':' 或 '*' 开头
}

// String 方法用于格式化输出节点的信息，方便调试
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.iswild)
}

// insert 方法将一个新的路径模式插入到前缀树中
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { // 如果已经到达路径模式的最后一个片段
		n.pattern = pattern // 将完整的路径模式存储在当前节点
		return
	}

	part := parts[height]       // 获取当前层级的路径片段
	child := n.matchChild(part) // 查找当前节点的子节点中是否有匹配的节点
	if child == nil {           // 如果没有匹配的子节点
		child = &node{part: part, iswild: part[0] == ':' || part[0] == '*'} // 创建一个新的子节点
		n.children = append(n.children, child)                              // 将新节点添加到子节点列表中
	}
	child.insert(pattern, parts, height+1) // 递归插入到子节点中
}

// search 方法在前缀树中搜索一个路径模式
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") { // 如果已经到达路径的最后一个片段，或者当前节点是通配符节点
		if n.pattern == "" { // 如果当前节点没有存储完整的路径模式
			return nil // 返回 nil，表示没有找到匹配的路径
		}
		return n // 返回当前节点，表示找到匹配的路径
	}

	part := parts[height]             // 获取当前层级的路径片段
	children := n.matchChildren(part) // 查找当前节点的所有匹配子节点
	for _, child := range children {  // 遍历匹配的子节点
		result := child.search(parts, height+1) // 递归搜索子节点
		if result != nil {                      // 如果找到匹配的路径
			return result // 返回匹配的节点
		}
	}
	return nil // 如果没有找到匹配的路径，返回 nil
}

// travel 方法用于遍历前缀树，并将所有存储了完整路径模式的节点添加到列表中
func (n *node) travel(list *([]*node)) {
	if n.pattern != "" { // 如果当前节点存储了完整的路径模式
		*list = append(*list, n) // 将当前节点添加到列表中
	}
	for _, child := range n.children { // 遍历子节点
		child.travel(list) // 递归遍历子节点
	}
}

// matchChild 方法查找当前节点的子节点中是否有与给定路径片段匹配的节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children { // 遍历子节点
		if child.part == part || child.iswild { // 如果子节点的路径片段与给定片段匹配，或者子节点是通配符节点
			return child // 返回匹配的子节点
		}
	}
	return nil // 如果没有找到匹配的子节点，返回 nil
}

// matchChildren 方法查找当前节点的所有匹配子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)          // 创建一个空的节点列表
	for _, child := range n.children { // 遍历子节点
		if child.part == part || child.iswild { // 如果子节点的路径片段与给定片段匹配，或者子节点是通配符节点
			nodes = append(nodes, child) // 将匹配的子节点添加到列表中
		}
	}
	return nodes // 返回匹配的子节点列表
}
