//变量的声明和初始化
package main

import (
	"fmt"
	"strconv"
)

func main() {
	juHe()
	fmt.Println("===============")
	fenSan()
	fmt.Println("===============")
	shortVar()
	fmt.Println("===============")
	convert()
	fmt.Println("===============")
	cross()
	fmt.Println("===============")
}

//只使用一个var关键字，批量声明多个变量
func juHe() {

	var (
		a    int
		b    string
		c    float32
		d    int8 = 100
		e    bool
		f    func() bool
		g    = "yang"
		h    = 400.2
		slic []int8
		arr  [3]int8
	)

	fmt.Printf("%T ==> %v	\n", a, a)
	fmt.Printf("%T ==> %v ==> %q \n", b, b, b)
	fmt.Printf("%T ==> %v	\n", c, c)
	fmt.Printf("%T ==> %v	\n", d, d)
	fmt.Printf("%T ==> %v	\n", e, e)
	fmt.Printf("%T ==> %v	\n", f, f)
	fmt.Printf("%T ==> %v	\n", g, g)
	fmt.Printf("%T ==> %v	\n", h, h)
	fmt.Printf("%T ==> %v	\n", slic, slic)
	fmt.Printf("%T ==> %v	\n", arr, arr)
}

//每个变量都使用一个var关键字进行声明
func fenSan() {
	//默认值为0
	var a int

	//默认值为空字符串【""】
	var b string

	//默认值为0
	var c float32

	//【%v】代表变量的值【仅当不确定变量的类型时使用进行占位】
	//【%d】代表数字类型的值
	//【%s】代表字符串类型的值
	//【%q】代表变量的引用
	//【%T】代表变量的类型
	fmt.Printf("%T ==> %v	\n", a, a)
	fmt.Printf("%T ==> %v ==> %q \n", b, b, b)
	fmt.Printf("%T ==> %v	\n", c, c)
}

//短变量初始化声明，此方式仅可用于局部变量的声明
func shortVar() {
	a := 90.3
	//a := 90.4  不允许单独的重复变量声明
	a, b := 9.9, "jiang" //批量的变量声明中如果其中包含之前已声明变量，那么会覆盖重复变量的旧值【前提时数据类型不变】
	fmt.Printf("%T ==> %v	\n", a, a)
	fmt.Printf("%T ==> %v	\n", b, b)
}

func convert() {
	a := 89
	b := strconv.Itoa(a)
	fmt.Printf("%d\t%s\n", a, b)

	c, err := strconv.Atoi(b)
	fmt.Println(err)
	fmt.Printf("%d\t\n", c)

	//匿名变量：【_】,只起到占位的作用，不会分配内存空间,这个位置上的变量也不能被使用，主要用于函数中
	d, _ := strconv.Atoi(b)
	fmt.Printf("%d\t\n", d)
}

//支持若干变量交换值，比其他语言更简洁
func cross() {
	x := 10
	y := 20
	z := 30
	fmt.Println(x, "===>", y, "===>", z)

	x, y, z = z, x, y
	fmt.Println(x, "===>", y, "===>", z)

}
