package main

import (
	"fmt"
	"math"
)

//匿名函数：是指在使用时才定义的，没有函数名，只有函数体函数，匿名函数也可以赋值给一个变量，也可以当作变量被传递
//匿名函数经常被用在实现函数回调、和闭包等场景

func main() {

	//1.在定义时调用匿名函数,形成即时执行体
	//无参，无返回值
	func() {
		fmt.Println("=======")
	}()

	//有参无返回值
	func(num int) {
		fmt.Println("Your input is: ", num)
	}(1000)

	//有参，有返回值
	sqrt := func(num float64) float64 {
		return math.Sqrt(num)
	}(25)
	fmt.Println(sqrt)

	//2.将匿名函数赋值给变量
	fun := func(a int) {
		fmt.Println(a)
	}
	fun(8)
}
