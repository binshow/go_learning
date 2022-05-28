package main

import "fmt"

type mapKey struct {
	key  int
}

// 如果要使用struct来作为key，必须要保证 struct 对象在逻辑上是一个不可变的
func main() {
	var m = make(map[mapKey]string)
	key := mapKey{key: 10}

	m[key] = "hello"
	fmt.Printf("m[key] = %s\n" , m[key])

	// 修改key的字段的值后再次查询map，无法获取刚才add进去的值
	key.key = 100
	fmt.Printf("再次查询m[key]=%s\n", m[key])
}
