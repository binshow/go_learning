package main

import (
	"fmt"
	"sync"
)

// 建造者模式

type Message struct {
	Header *Header
	Body   *Body
}

type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

type Body struct {
	Items []string
}


func main() {

	//1. 传统的创建一个Message对象
	// 多层的嵌套实例化
	message := Message{
		Header: &Header{
			SrcAddr:  "192.168.0.1",
			SrcPort:  1234,
			DestAddr: "192.168.0.2",
			DestPort: 8080,
			Items:    make(map[string]string),
		},
		Body:   &Body{
			Items: make([]string, 0),
		},
	}
	// 需要知道对象的实现细节
	message.Header.Items["contents"] = "application/json"
	message.Body.Items = append(message.Body.Items, "record1")
	message.Body.Items = append(message.Body.Items, "record2")


	//2. 使用建造者模式来创建对象

	msg := Builder().
		WithSrcAddr("192.168.0.1").
		WithSrcPort(1234).
		WithDestAddr("192.168.0.2").
		WithDestPort(8080).
		WithHeaderItem("contents", "application/json").
		WithBodyItem("record1").
		WithBodyItem("record2").
		Build()
	if msg.Header.SrcAddr != "192.168.0.1" {
		fmt.Printf("expect src address 192.168.0.1, but actual %s.", message.Header.SrcAddr)
	}
	if msg.Body.Items[0] != "record1" {
		fmt.Printf("expect body item0 record1, but actual %s.", message.Body.Items[0])
	}

}


// Message对象的Builder对象
type builder struct {
	once *sync.Once
	msg *Message
}
// Builder pattern
func Builder() *builder {
	return &builder{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{} , Body: &Body{}},
	}
}

func (b *builder) WithSrcAddr(scrAdder string) *builder {
	b.msg.Header.SrcAddr = scrAdder
	return b
}

func (b *builder) WithSrcPort(srcPort uint64) *builder {
	b.msg.Header.SrcPort = srcPort
	return b
}
func (b *builder) WithDestAddr(destAddr string) *builder {
	b.msg.Header.DestAddr = destAddr
	return b
}
func (b *builder) WithDestPort(destPort uint64) *builder {
	b.msg.Header.DestPort = destPort
	return b
}
func (b *builder) WithHeaderItem(key, value string) *builder {
	// 保证map只初始化一次
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}
func (b *builder) WithBodyItem(record string) *builder {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}
// Build() 创建Message对象，在最后一步调用
func (b *builder) Build() *Message {
	return b.msg
}





