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
