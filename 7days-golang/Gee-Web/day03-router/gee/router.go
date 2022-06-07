package gee

import (
	"net/http"
	"strings"
)

// 使用 Trie树实现动态路由 Dynamic route

type router struct {
	roots    map[string]*node       //存储每种请求方法的 Trie树根节点 ， 方法包括 get、post、put等
	handlers map[string]HandlerFunc //存储每种请求方式的 HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

// Only one * is allowed
// 解析路由
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// 添加路由 ===>  插入前缀树的节点
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) //

	// 用请求的方法来创建一个根节点
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	// 根节点下面继续插入
	r.roots[method].insert(pattern, parts, 0)

	key := method + "-" + pattern //GET-/
	r.handlers[key] = handler
}

// 查找路由 ===>  查询前缀树的节点
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	//fmt.Println(searchParts) // [hello geektutu]

	params := make(map[string]string)
	// 查询对应的方法有没有根节点
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	// 查询对应的路径数组有没有节点，返回的是 叶子节点
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// 根据方法查询所有路由
func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c) // 执行对应的handler
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
