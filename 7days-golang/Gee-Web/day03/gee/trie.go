package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 当前节点保存路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 找到当前节点的子节点中是否有对应part的节点，如果有就返回第一个对应的子节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 找到当前节点的子节点中是否有对应part的节点，返回所有对应的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 前缀树节点的插入，有多个part. 对应于路由注册
func (n *node) insert(pattern string, parts []string, height int) {

	// 如果树的高度已经到了最后一个part了,就说明到达叶子节点了
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 找到当前需要对应part的子节点，如果没有就构造
	part := parts[height]
	child := n.matchChild(part)

	// 递归查找每一层的节点，如果没有匹配到当前part的节点，则新建一个
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}

	// 如果有对应的子节点，就在子节点下面继续插入
	child.insert(pattern, parts, height+1)
}

// 查询part在树中的路径，返回最后一个子节点
// 有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。
// 因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	// 第一个路径的节点
	part := parts[height]
	children := n.matchChildren(part)

	// 遍历每个节点，继续往下找，直到找不到为止
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
