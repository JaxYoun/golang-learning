//map数据类型测试
//1.声明与构造
//2.访问与迭代
package main

import "fmt"

//1.只做类型声明
var map0 map[string]string

//2.声明并初始化，但是使用类型推断
var map1 = make(map[string]int64)

//3.声明并带上类型，同时初始化，但是此处的类型声明太冗余
var map2 map[string]int64 = make(map[string]int64)

func main() {
  //4.map对象初始化
  map0 = make(map[string]string
  //写元素
  map0["name"] = "yang"
  fmt.Println(map0["name"])
  
  //4.遍历map对象
  for k, v := range map0{
    fmt.Println(k, v)
  }
}
