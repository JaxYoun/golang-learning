//defer提供了延时执行机制，会在一个函数其他部分执行完毕后才执行defer栈中的多个defer语句，用处类似java中的finally语句，值得注意的是
//多个defer会根据逻辑先后顺序依次入栈，每个栈帧都有自己独立的逻辑数据，函数其他部分执行完毕后再逆序出栈执行
//设计defer关键字的目的是资源管理与异常处理
package main


//1.对于非计算类defer语句，会在声明时确定变量的值
func main() {
  
  
}

//1.defer序列会被放到
func test() {
  i := 0
  
  //后执行
  defer fmt.Printf("kkkk", i)  //0
  i++
  
  //先制性
  defer fmt.Printf("kkkk", i)  //1
  i++
  
  return
}
