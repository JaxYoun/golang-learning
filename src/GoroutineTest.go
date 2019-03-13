// 1.进程和线程不仅仅是编程语言中的概念，同时在操作系统层面是实实在在存在与之对应的执行单位的
// 而协程在操作系统层面就没有与之对应的执行实体，它纯粹是编程语言界提出的实现机制，可以将一个协程视为一个函数调用，
//协程每执行一次只在线程栈上增加一个对应的栈帧；支持协程的语言
// 会在语言编译器根据语法规则将一个个协程语法单元编译为执行单元，运行时由特殊的调度线程将若干协程
// 分发到多个线程上执行，操作系统自始至终都对协程机制无感知。
// 2.在同一个线程上执行的若干协程是串行阻塞执行的，可以理解为同一个线程上的多个协程是排队顺序执行的，在没有发生阻塞或中断的情况下，
//只有上一个携程完整的执行结束才能执行下一个，即同一个线程上的协程间不会交替穿插执行。
// 3.由于协程被分配到多个线程上执行，而线程的执行是由系统调度的，是无序的，这就导致了协程执行的无序性。
package main

import (
	"fmt"
	"time"
)

//1.显式声明一个普通函数
func myFun(arg int) {
	fmt.Println("----->", arg)
}

func main() {

	//2.使用go关键字调用函数，这个函数就会以协程的方式执行
	go myFun(100)

	//3.以go关键字和匿名函数的方式启动一个协程
	go func(arg float64) {
		fmt.Println("----->", arg)
	}(55.89)

	//4.使用非缓存channel来实现协程间的通信
	ch := make(chan int)

	//5.此协程向管道中写入值
	go func(cha chan int) {
		cha <- 12 //向channel写数据
	}(ch)

	//6.此协程从管道读取值
	go func(cha chan int) {
		kk := <-cha //从channel读数据
		fmt.Println("从channel中读值-->", kk)
	}(ch)

	time.Sleep(time.Second)
	fmt.Println("main end----")

	blockedChannel()
	time.Sleep(3 * time.Second)

}

//7.阻塞式管道式非缓冲的，遵从生产者-消费者模式，一个管道只有生产者写入了数据才能消费，同时只有消费者消费了数据后生产者才能往里写
func blockedChannel() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(cha chan int) {
			cha <- i
		}(ch)
	}

	for i := 0; i < 10; i++ {
		go func(cha chan int) {
			val := <-cha
			fmt.Println("从channel中读值==>", val)
		}(ch)
	}
}
