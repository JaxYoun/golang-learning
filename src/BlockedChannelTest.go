// 1.Golang中协程间的数据同步有两种方式可以实现：锁（*sync.Mutex）和管道（channel），
// 锁的原理和多线程类似，实现比较复杂，所以优先使用管道
package main

import (
	"fmt"
	"time"
)

var cha = make(chan int)

func Write(cha chan int) {
	fmt.Println("Write 前")
	cha <- 1
	fmt.Println("Write 后")
}

func Read(cha chan int) {
	fmt.Println("Read 前")
	val := <-cha
	fmt.Println("Read 后")
	fmt.Println(val)
}

func kk() {
	chaArr := make([]chan int, 10)
	for i := range chaArr {
		fmt.Println("index---", i)
		chaArr[i] = make(chan int)
		go func(arg chan int) {
			arg <- i
		}(chaArr[i])
	}

	time.Sleep(3 * time.Second)

	for _, v := range chaArr {
		go func(arg chan int) {
			val := <-arg
			fmt.Println("++++++++", val)
		}(v)
	}

}

func main() {

	//非缓冲channel会引发消费者和生产者的阻塞，两种协程直到满足各自的条件再继续执行，利用此机制能很好的实现协程间的数据同步
	go Read(cha)
	go Write(cha)
	time.Sleep(time.Second)

	//kk()
	//time.Sleep(5 * time.Second)
}
