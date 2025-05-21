package gee

import "strings"

// Trie 路由树节点实现
type trieNode struct {
	pattern  string      // 完整的路由路径，仅叶子节点存储，如：/p/:id
	part     string      // 路由中的一部分，如：:id
	children []*trieNode // 子节点切片，如：[doc, tutorial, intro]
	isWild   bool        // 是否精准匹配
}

// 用于找到插入新路由的位置
func (n *trieNode) matchChild(part string) *trieNode {
	for _, child := range n.children {
		// 优先返回精确匹配的子节点
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 用于找到所有匹配的子节点，用于查找路由
func (n *trieNode) matchChildren(part string) []*trieNode {
	nodes := make([]*trieNode, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入新的路由规则
func (n *trieNode) insert(pattern string, parts []string, height int) {
	// 叶子节点时设置完整路由路径
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 查找匹配到的节点
	part := parts[height]
	child := n.matchChild(part)
	// 没有能够匹配的节点时，创建新节点并插入当前节点的子节点切片中
	if child == nil {
		child = &trieNode{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	// 递归插入下一层的路由节点
	child.insert(pattern, parts, height+1)
}

// 匹配已有的路由路径
func (n *trieNode) search(parts []string, height int) *trieNode {
	// 叶子节点或存在通配符
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.part == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		// 递归匹配下一层
		res := child.search(parts, height+1)
		if res != nil {
			return res
		}
	}

	return nil
}
