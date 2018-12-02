//（四）、函数变量（函数作为值）
//*在Golang中，函数也是一种类型，可以像其他类型一样被保存到变量中。
//*可以通过type关键字自定义一个函数相关的类型，但是必须保证该类型与函数实现的参数及返回值的签名一致。
package main

import (
	"fmt"
	"strings"
)

//1.函数可以被当作变量值，赋给一个变量
//2.函数可以被当作参数传递给另一个函数

func main() {
	str := "abCdAEIjkLw"

	//单纯调用工具函数
	fmt.Println(processFun(str))

	//显式的将工具函数作为参数传递给包装函数
	fmt.Println(stringToCase(str, processFun))

	//使用函数类型声明式调用
	fmt.Println(stringToCase0(str, processFun))

}

//工具函数：将字符串奇偶位置上的字母大小写交替出现
func processFun(str string) (result string) {
	for i, v := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(v))
		} else {
			result += strings.ToLower(string(v))
		}
	}
	return
}

//原始传递，函数类型参数的书写过于啰嗦
func stringToCase(str string, fun func(string) string) string {
	return fun(str)
}

//函数类型的使用步骤
//1.定义一个函数类型
//2.实现定义的函数类型
//3.将函数作为参数调用
//类似接口的用法
//设计通用接口时非常灵活、有用

//类型声明式传递
type caseFunc func(string) string

//原始传递，函数类型参数的书写更简洁
func stringToCase0(str string, fun caseFunc) string {
	return fun(str)
}
