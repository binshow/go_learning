package main


//func main() {
//	var t T
//	t.Get()
//	t.Set(1)
//
//	// 下面这种使用类型调用方法的表达式称为 Method Expressopn
//	// 类型T只能调用 T的方法集合，*T也只能调用 *T的方法集合
//	var t2 T
//	T.Get(t2)
//	(*T).Set(&t2 , 1)
//
//	f1 := (*T).Set
//	f2 := T.Get
//	fmt.Printf("the type of f1 is %T\n" , f1) // the type of f1 is func(*main.T, int) int
//	fmt.Printf("the type of f2 is %T\n" , f2) // the type of f2 is func(main.T) int
//	f1(&t , 3)
//	fmt.Println(f2(t)) // 3
//}
