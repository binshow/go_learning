package main

// 责任链模式
/**

 职责链模式(Chain of Responsibility Pattern)：
避免请求发送者与接收者耦合在一起，让多个对象都有可能接收请求，
将这些对象连接成一条链，并且沿着这条链传递请求，直到有对象处理它为止

职责链模式包含如下角色：
• Handler: 抽象处理者：定义出一个处理请求的接口。如果需要，接口可以定义出一个方法，以设定和返回对下家的引用。这个角色通常由一个抽象类或接口实现。
• ConcreteHandler: 具体处理者：具体处理者接到请求后，可以选择将请求处理掉，或者将请求传给下家。由于具体处理者持有对下家的引用，因此，如果需要，具体处理者可以访问下家。
• Client: 客户端

 */


// 以医院为例子


func main() {
	//1. 初始化责任链：病人来了 -->  挂号处挂号 --> 找医生看病 --> 收银处缴费 --> 药房拿药
	medical := &medical{}
	//Set next for cashier department
	cashier := &cashier{}
	cashier.setNext(medical)
	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(cashier)
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	//2. 病人来看病了
	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}
