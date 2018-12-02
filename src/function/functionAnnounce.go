//The ways to annoucing a function!
package main

import "fmt"

//1.函数支持多个参数和多个返回值
//2.函数签名中类型相同且相邻的参数或返回值可以批量声明
//3.函数的返回值类型不可省略，但是变量名可以省略，如果有多个返回值需要使用小括号
//4.如果函数签名只做了返回值类型的声明，函数体内必须声明对应类型的变量，并显式返回这些变量
//5.如果函数签名中对返回变量名做了声明，当在函数体内对这些返回变量做了赋值操作后，函数体内可以不一一显式返回
//6.如果一个函数有多个返回值，但是在被调用时又不使用其所有返回值，可以使用下划线对不使用返回值进行占位，
// 形成匿名变量不为其分配内存，已达到部分规避的目的

func main() {
	fmt.Println(mySum(10, 20))
	fmt.Println(mySum0(30, 40))
	fmt.Println(mySum1(50, 60))

	//使用匿名变量
	sum, result0, _ := myFun0(1, 2, "yang")
	fmt.Println(sum, result0)
}

//1.多参数单返回值，且省略了返回值变量
func mySum(a int, b int) int {
	he := a + b
	return he
}

//2.多参数、多返回值批量声明，内部只需赋值，不用一一显式返回
func mySum0(a, b int) (he, shang int) {
	he = a + b
	shang = a * b
	return
}

//3.省略返回变量名，必须内部声明并显式一一返回
func mySum1(a, b int) (int, int) {
	he := a + b
	shang := a * b
	return he, shang
}

//4.综合案例
func myFun0(num0, num1 int, name string) (sum int, result0, result1 string) {
	sum = num0 + num1
	result0 = name
	result1 = result0
	return
}
