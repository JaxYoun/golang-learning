//闭包概念与测试
//内部匿名函数 + 被匿名函数引用的内部局部变量构成的应用上下文，外层函数内层匿名函数
//对闭包外层函数的调用返回的是内层的匿名函数的引用，这样的程序体具有记忆状态数据功能
//其实Lambda表达式就是闭包的一种实现

package main

import "fmt"

func main() {
  slice := []int{1, 2, 3, 4}
  fmt.Println(OuterFunc()(slice...))  //此处的切片作为可变参数集合，需要使用省略号
  
  arr := [...]int{1, 23, 3, 4}
  fmt.Println(OuterFunc()(arr[:]...))  //如果参数集合是数组，需要将数组转为切片，然后将切片作为参数传递，但仍然需要使用省略号
  
  fmt.Println(OuterFunc()(1, 2, 3, 4, 5))  //传入零散值
  
  //对同一个闭包函数的多次调用，数据结果具有连续性，闭包执行体对状态数据具有记忆功能
  fun1 := MyClosure()
  fmt.Pringln(fun1(1, 2, 3))  //6
  fmt.Pringln(fun1(1, 2, 3))  //12
  
  //不同函数的调用数据不具有连续性
  //【值得注意的是，每次返回的内层匿名函数的引用是相同的，这可以理解，其实返回的匿名函数本质是一个真实的函数，所以其引用也只有一份】
  fun2 := MyClosure()
  fmt.Pringln(fun2(1, 2, 3))  //6
  fmt.Pringln(fun2(1, 2, 3))  //12
  
}

//1.常规写法，稍显嗦，但可读性强
func MyClosure() func(...int) int {
  sum := 0
  fun := func(slice ...int) int {  //匿名内部函数接收可变参数
    for _, v := range slice {
      sum += v
    }
    return sum
  }
  return fun
}

//2.简练写法
func main() {
  slice := []int{2, 3, 4}
  
  //a.返回内层函数，最后独立调用
  reFun := func() func(...int) int {
    sum := 0
    return func(arr ...int) int {
      for _, v := range arr{
        sum += v
      }
      retun sum
    }
  }()  //即时执行外层匿名函数，返回内层函数待后续调用
  fmt.Println(reFun(1, 2, 3, 4))  //可变参数
  fmt.Println(reFun(slice...))  //切片作为可变参数
  
  //b.返回内层函数调用执行的最终的结果
  result := func() func(...int) int {
    sum := 0
    return func(arr ...int) int {
      for _, v := range arr{
        sum += v
      }
      retun sum
    }
  }()(slice...)  //两次即时执行
  fmt.Println("%T---%v", result, result)  //in---10
}








