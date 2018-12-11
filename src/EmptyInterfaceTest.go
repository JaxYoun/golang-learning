//Glang中的空接口类似于Java中的Object类，可以作为任何类的抽象接口，带来更大的灵活性
package main

import "fmt"

func main() {

	//空接口变量可以赋任何类型的值
	var val interface{}
	val = 8.1
	fmt.Printf("%T\n", val)

	//可以通过判断变量的类型，然后决定是否执行后续逻辑
	if v, bul := val.(float32); bul { //类型查询语法
		fmt.Println(v, bul)
	} else {
		fmt.Println("type missmatch!")
	}

	switch val.(type) {
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
