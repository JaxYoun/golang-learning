//select是Golang中实现异步的重要特性
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func(arg chan int) {
		arg <- 1
	}(ch)

	//添加睡眠的目的是让协程有足够的执行时间，保证后续的主函数是在协程执行完后执行的
	time.Sleep(time.Second)

	select {
	case <-ch:
		fmt.Println("1111")
	default:
		fmt.Println("default")
	}

}
