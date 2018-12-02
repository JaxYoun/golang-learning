//使用匿名函数实现函数回调
package main

import "fmt"

func main() {
	//定义一个切片，对其中的元素遍历进行开方和平方运算
	slice := []int{1, 2, 3, 4, 5, 9, 6, 7, 8}
	sl := filter(slice, func(i int) int {
		return i * i
	})
	fmt.Println(sl)

}

type myFunc func(int) int

func filter(slice []int, f myFunc) (re []int) {
	for _, v := range slice {
		re = append(re, f(v))
	}
	return
}
