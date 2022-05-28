package main


// 定义一个 struct , 首字母小写说明只能在本包内创建
type person struct {
	Name 	string	// field 首字母大写说明 可以通过 p.Name来直接访问
	Age  	int
	phone   string	// field 首字母小写则只能在本包内 p.phone来访问，其他包不能访问
}

// 首字母大小表面可以在其他包中使用 包名.structName来创建
type Student struct {
	Major 	string
	Class  	int
	email   string
}


// 可以通过 struct 绑定的方法来对私有的field进行封装
func (p person) GetPhone() string {
	return p.phone
}
