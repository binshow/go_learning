package gee

import (
	"fmt"
	"strings"
)

// -------------------------------------------
// @file          : trie.go
// @author        : binshow
// @time          : 2022/6/7 10:58 PM
// @description   : 使用前缀树来实现动态路由，比如 /hello/:name 同时支持 /hello/tom 和 /hello/jack
// -------------------------------------------


// HTTP请求的路径恰好是由/分隔的多段构成的，因此，每一段可以作为前缀树的一个节点。
// 我们通过树结构查询，如果中间某一层的节点都不满足条件，那么就说明没有匹配到的路由，查询结束


//前缀树的节点
type node struct {
	pattern  string  // 完整的待匹配路由，例如 /p/:lang
	part     string  // 当前节点保存路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

func (n *node) String() {
	fmt.Printf("n.pattern = %s , n.part = %s , n.children = %v , n.isWild = %v\n" , n.pattern , n.part , n.children , n.isWild )
}


//前缀树节点的插入：pattern 完整的路由地址 ， parts 拆分的每个路由地址 ， height 已经确定的parts中的数量
func (n *node) insert(pattern string, parts []string, height int) {
	//1. 递归截止条件，parts中的内容已经全部赋值到节点上了，同时也说明已经到达叶子节点了
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	//2. 找到当前需要赋值的 parts的内容
	part := parts[height]

	//3. 找到当前节点的 子节点，看是否有子节点已经存在这个part的值了
	child := n.matchChild(part)

	//4. 如果没有就新建一个子节点
	if child == nil {
		child = &node{
			part: part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children , child)
	}

	//5. 现在已经有对应的子节点了，就递归的往这个子节点下面插入后序的值
	child.insert(pattern , parts , height+1)

}


//前缀树节点的查询，在n的子节点下查询 parts 所在的路径
// 有一点需要注意，/p/:lang/doc只有在第三层节点，即doc节点，pattern才会设置为/p/:lang/doc。p和:lang节点的pattern属性皆为空。
// 因此，当匹配结束时，我们可以使用n.pattern == ""来判断路由规则是否匹配成功。
func (n *node) search(parts []string, height int) *node {
	//1. 递归截止条件，已经查询到最后一层了, 如果某一层节点 为 *，那后面的都不用查询了
	if len(parts) == height || strings.HasPrefix(n.part , "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	//2. 找第一个路径所在的子节点集合
	part := parts[height]
	children := n.matchChildren(part)

	//3. 遍历每个节点，递归的往下找，直到找不到为止
	for _, child := range children {
		result := child.search(parts , height+1)
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


// 判断当前节点的 子节点中 是否含有 part 的值
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
