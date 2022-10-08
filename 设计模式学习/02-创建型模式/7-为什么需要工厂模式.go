package main

import "fmt"

// -------------------------------------------
// @file          : 7-为什么需要工厂模式.go
// @author        : binshow
// @time          : 2022/10/8 9:12 PM
// @description   :
// -------------------------------------------

// 下面这样设计的缺点：
//1.大量的if-else代码块
//2. Fruit 类负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类实现，违反了"单一职责原则"
//3. 当需要增加新类型的水果时，需要修改 Fruit 类的构造函数 NewFruit 和其他方法源代码，违反了"开闭原则"

//水果类
type Fruit struct{}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("i am apple")
	} else if name == "banana" {
		fmt.Println("i am banana")
	} else if name == "pear" {
		fmt.Println("i am pear")
	}
}

//创建一个Fruit对象
func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
	} else if name == "banana" {
	} else if name == "pear" {
	}
	return fruit
}

// 在业务逻辑层中，业务层的开发逻辑也需要依赖于 Fruit 模块的更新和改变，也就是说：
//  业务逻辑层 ----> 基础类模块

//  如何解耦呢？这样就引入了 工厂模式
//  业务逻辑层 ----> 工厂模式 ----> 基础类模块

func main() {

	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")

	pear := NewFruit("pear")
	pear.Show("pear")

}
