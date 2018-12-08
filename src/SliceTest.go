//切片声明、访问、遍历
package main

import "fmt"

//1.由于切片是引用类型（数组是值类型），所以声明而不初始化的切片变量为nil，但是可以求取其长度，且其长度为0
var sli []int

//2.显式声明并赋值
var sli0 = []int{1, 2, 3, 4, 5, 6, 7}

var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	fmt.Println(sli == nil) //true
	fmt.Println(len(sli))   //0

	//3.短变量显示声明
	sli1 := []int{1, 2, 3} //此时切片的长度为5，容量也为5
	fmt.Println(len(sli1), cap(sli1))

	//4.使用make()函数
	sli2 := make([]int, 5) //此时切片的长度为5，容量也为5
	fmt.Println(len(sli2), cap(sli2))

	sli3 := make([]int, 5, 8) //此时切片的长度为5，容量也为8
	fmt.Println("leng=", len(sli3), "cap=", cap(sli3))

	//5.数组截取
	sli4 := arr[:]   //全量截取
	sli5 := arr[3:5] //中部截取，左闭右开区间
	sli6 := arr[:5]  //从0到(5 - 1)
	sli7 := arr[4:]  //从4开始到其后全部
	fmt.Println(len(sli4), cap(sli4))
	fmt.Println(len(sli5), cap(sli5))
	fmt.Println(len(sli6), cap(sli6))
	fmt.Println(len(sli7), cap(sli7))

	//6.切片截取
	sli8 := sli0[:]   //全量截取
	sli9 := sli0[3:5] //中部截取，左闭右开区间
	sli10 := sli0[:5] //从0到(5 - 1)
	sli11 := sli0[4:] //从4开始到其后全部
	fmt.Println(len(sli8), cap(sli8))
	fmt.Println(len(sli9), cap(sli9))
	fmt.Println(len(sli10), cap(sli10))
	fmt.Println(len(sli11), cap(sli11))
}
