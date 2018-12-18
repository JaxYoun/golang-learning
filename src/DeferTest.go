//1.defer提供了延时执行机制，会在一个函数的其他部分执行完毕后才执行defer栈中的多个defer语句，用处类似java中的finally语句
//2.值得注意的是存在多个defer语句时，会根据正常的执行顺序依次入栈（但暂不执行），每个defer语句就是一个栈帧，每个帧都有自己独立的【正常】逻辑数据，函数其他部分执行完毕后再逆序出栈执行
//3.defer语句可以读取有名返回值
//4.【return最先执行-->return负责将结果写入返回值的变量中-->接着defer开始执行一些收尾工作-->最后函数携带当前返回值退出】
//这个执行顺序很重要：
//  a.在显式声明了返回变量的函数中，如果defer语句引用了返回变量则defer语读写到的变量值与最终return的返回值相同
//  b.在未声明返回变量的函数中，defer语句中引用的数据是其在入栈时的局部变量
//以上两种现象可以通过变量地址的作用域来理解：显式声明函数返回变量时，调用方引用了此返回变量，内部return写了此变量，defer语句写了此变量，整个过程都共享了返回值变量
//隐式声明时调用方与return共同读写一块地址，但是被调用函数内部的defer语句使用的仍然是内部函数的局部变量，由于地址的分离所以一旦return后，
//返回数据不会与内部defer语句产生联系
//设计defer关键字的目的是资源管理与异常处理
package main


//1.对于非计算类defer语句，会在声明时确定变量的值
func main() {
  
  test()
  fmt.Println(test0())  //1
  fmt.Println(test1())  //0
}

//1.defer序列遵循后进先出原则
func test() {
  i := 0
  //后执行
  defer fmt.Printf("kkkk", i)  //0
  i++
  
  //先执行
  defer fmt.Printf("kkkk", i)  //1
  i++
  
  return
}

//2.显示声明返回变量，最终返回值与defer是同步的
func test0() (i int) {
  defer func() {
    defer fmt.Println(i)  //2
    i++
  }
  defer func() {
    defer fmt.Println(i)  //1
    i++
  }
  return i
}

//3.未显示声明返回变量，最终返回值与defer是非同步的
func test0() int {
  i := 0
  defer func() {
    defer fmt.Println(i)  //2
    i++
  }
  defer func() {
    defer fmt.Println(i)  //1
    i++
  }
  return i
}
