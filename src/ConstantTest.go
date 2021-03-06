//常量、常量组、iota关键字
//1.常量使用const 关键字定义，常量的类型必须是基本数据类型如：数字、bool、string等，可以显示声明类型也可以利用自动推断机制
//2.常量组使用const(a = 12; b = 22)定义
//3.常量的访问权限也遵从首字母大小写原则，但是常量被定义后如果无访问，仍然可以通过编译
//4.常量组可以当作枚举量来使用
//5.常量组中可以使用iota关键字或包含iota的表达式为元素赋值，每个iota都代表的是当前元素在常量组中的位置索引，是一个不定的值
//6.除了用在常量组中，iota同样也可用于单个常量的声明表达式中
package main

import "fmt"

const _A int = 12 //显式声明常量的类型，常量未使用编译可以通过
const _B = 13     //推断常量类型

const (
	PI   = 3.14
	Nike = "nk"
	ali         //如果未显式赋值，会将沿用上一条记录的值
	num  = iota //这个num接收的iota值是编译器在编译时根据常量组元素个数推算出来的
)

func main() {
	fmt.Println(_B)
	fmt.Println(num) //3
}
