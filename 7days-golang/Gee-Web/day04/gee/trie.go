package gee

import "strings"

/**
手写前缀树实现：
1. 参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc。
2. 通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，
这种模式常用于静态服务器，能够递归地匹配子路径
 */

type node struct {
	pattern string   // 待匹配路由，例如 /p/:lang
	part    string   // 当前节点保存路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild	bool   // 是否精确匹配，part 含有 : 或 * 时为true
}

// 找到当前节点的子节点中是否有对应part的节点，如果有就返回第一个对应的子节点
func (n *node) matchChild(part string) *node {
	for _,child := range n.children{
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 找到当前节点的子节点中是否有对应part的节点，返回所有对应的子节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children{
		if child.part == part || child.isWild {
			nodes = append(nodes , child)
		}
	}
	return nodes
}


// 前缀树节点的插入，有多个part
func (n *node) insert(pattern string , parts []string , height int)  {

	// 如果树的高度已经到了最后一个part了,就说明到达叶子节点了
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	// 找到当前需要对应part的子节点，如果没有就构造
	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		 child = &node{
			 part:     part,
			 isWild: part[0] == ':' || part[0] == '*',
		 }
		 n.children = append(n.children , child)
	}

	// 如果有对应的子节点，就在子节点下面继续插入
	child.insert(pattern , parts , height+1)
}


// 查询part在树中的路径，返回最后一个子节点
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
