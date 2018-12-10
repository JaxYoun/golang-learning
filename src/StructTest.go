//可以通过匿名父类变量的形式实现继承，但不具有父子类之间的多态性
package main

import (
	"fmt"
)

//1.声明父类
type Human struct {
	Id      int
	Name    string
	Age     int
	Married bool
}

//2.父类方法声明
func (this Human) copyBindExtend(i int) int {
	this.Id += i
	return this.Id
}

//3.声明子类
type User struct {
	//4.采用匿名变量的形式实现继承
	Human
	Account int
}

//5.拷贝绑定式方法声明，这种方式声明的方法不能将执行过程作用于原对象，只能作用于对象的拷贝，不具备Java中类似于对this的操作
func (this User) copyBind(i int) int {
	this.Age += i
	return this.Age
}

//6.引用绑定式方法声明，这种方法可以操作对象应用指向的原对象，类似Java中的this
func (this *User) referenceBind(i int) int {
	this.Age += i
	return this.Age
}

func main() {
	extendTest()
	methodTest()
}

//继承特性测试
func extendTest() {
	//7.实例化子类对象，继承父类属性
	user := new(User)
	fmt.Println(user.Account, user.Married)

	//8.实例化子类对象，继承父类属性
	user0 := User{}
	fmt.Println(user0.Account, user0.Married)

	//9.实例化子类对象，继承父类方法
	user1 := User{Human{}, 1}
	user1.copyBindExtend(300)
}

//方法特性测试
func methodTest() {
	user := new(User)
	fmt.Println("执行方法前===", user.Age) //0

	fmt.Println(user.copyBind(10))    //10
	fmt.Println("拷贝绑定后===", user.Age) //0

	fmt.Println(user.referenceBind(10)) //10
	fmt.Println("引用绑定后===", user.Age)   //10

	fmt.Println(user.copyBindExtend(12))
	fmt.Println("继承拷贝绑定", user.Id)
}
