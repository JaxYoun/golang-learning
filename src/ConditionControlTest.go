package main

import "fmt"

func main() {
	num := 120 / 3

	//1.简单if-else条件判断
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	//2.局部变量初始化赋值的if表达式
	if i := 12; i < 100 {
		fmt.Println(i, "is less than one hundred！")
	}
}
