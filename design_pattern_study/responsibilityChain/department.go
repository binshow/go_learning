package main

// 看病的行为
type department interface {
	execute(*patient) //当前部门对这个病人处理
	setNext(department) // 调用下一个
}
