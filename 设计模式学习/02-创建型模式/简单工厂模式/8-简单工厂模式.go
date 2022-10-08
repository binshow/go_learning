package main

import "fmt"

// -------------------------------------------
// @file          : 8-简单工厂模式.go
// @author        : binshow
// @time          : 2022/10/8 9:35 PM
// @description   :
// -------------------------------------------

// ========= 抽象层 ===========

type Fruit interface {
	Show()
}

// ========= 基础类 ===========

type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (b *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct{}

func (p *Pear) Show() {
	fmt.Println("我是梨子")
}

// ========= 工厂类，专门生产水果 ===========

type FruitFactory struct{}

func (factory *FruitFactory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// ========= 业务逻辑层 ===========
// 业务逻辑只和 工厂模块进行依赖， 不再关系 Fruit 类是具体如何创建对象的

func main() {
	factory := FruitFactory{}
	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
