//数组声明和遍历
package main

import "fmt"

//0.数组的显式声明必须指定长度
//1.声明数组但不初始化元素，由于数组为值类型，此时数组变量不为nil，而是每个元素都是默认值
var arr00 [5]int

//2.显示声明一个数组，并初始化其元素，但是元素个数少于数组长度，此时未赋值的元素值为默认值
var arr0 = [5]int{1, 2, 3, 4}

func main() {
	fmt.Println(arr00)
	fmt.Println(arr0)

	//3.短变量声明法
	arr1 := [5]int{6, 7, 8, 9}
	fmt.Println(arr1)

	//4.长度推断式声明
	arr2 := [...]int{2, 3, 4, 5, 6, 7}
	fmt.Println(arr2)

	//5.使用内置函数len()，求取数组长度
	fmt.Println(len(arr2))

	//6.使用数组index随机访问数组元素
	fmt.Println(arr2[4])

	//7.使用index遍历数组，【值得注意的是，for循环体中只会执行一次len()函数】
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	//8.for-range结构遍历数组
	for i, v := range arr2 {
		fmt.Println("intex", i, "value", v)
	}

	//如果返回值前面为匿名变量，必须使用下划线占位
	for _, v := range arr2 {
		fmt.Println("value", v)
	}

	//如果返回值第一个不为匿名变量，后面的为匿名变量，后面连续为匿名变量就可以省略下划线
	//但是如果后续匿名变量排列不连续则必须使用下划线占位
	for i := range arr2 {
		fmt.Println("intex", i)
	}
}
