package main

import (
	"fmt"
	"runtime"
)

/**
1.并行计算是指同一时刻在围观层面上真正的有多个任务同时在执行，需要多核cpu，单核不能实现并行。
2.并发是指从宏观角度看有多个任务在同时执行，单核也可以通过时分多路复用的方式来实现，是“看起来”同时执行多个任务的假象。
*/

/**
1.首先协程时编程语言层面的概念，其调度机制完全由golang自启专有调度器线程来完成。
2.协程最终还是被分配到物理线程上，每个协程会被整体分配到一个线程上去，一个协程内部时顺序执行的。
3.当一个物理线程上被分配多个协程时，协程间也是顺序执行的。
4.当一个物理线程上有多个协程时，有协程发生阻塞，调度器会试图将后续协程迁移到其他线程上去（类似Java中的ForkJoin机制）。
5.在单线程环境下，其上的一个协程得到资源后，如果不在协程里做主动交出操作，就会一致占用，直到此协程执行完毕才会交出来。
5.在协程里调用runtime.Gosched()函数会使当前协程让出一次cpu执行权（类似于yield关键字），后排协程意外获得执行机会。
*/
var (
	fun1 = func() {
		for i := 0; i < 30; i++ {
			if i == 10 {
				runtime.Gosched()
			}
			fmt.Printf("---------%d \n", i)
		}
	}

	fun2 = func() {
		for i := 0; i < 30; i++ {
			if i == 20 {
				runtime.Gosched()
			}
			fmt.Printf("++++++++++%d \n", i)
		}
	}
)

func main() {
	/*go fun1()
	go fun2()
	time.Sleep(time.Minute * 10)*/

	//1.获取当前cpu核心数
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)

	//2.将本次执行分配到几个线程上
	runtime.GOMAXPROCS(1)
	go say("hello")
	speak("world") //换为调用speak函数就能交替执行
}

//在【单线程】前提下
//1.如果去掉此函数中的gosched函数，由于主协程是入口，它排在前面。
//2.所以在主协程中的speak()函数调用将极有【可能】先执行，打印1000个world，子协程可能得不到执行机会，主协程退出了，也可能都打印出来。
//3.在speak寒素中调用gosched函数后就能让协程交替执行。
func speak(input string) {
	for i := 0; i < 1000; i++ {
		runtime.Gosched()
		fmt.Printf("===== %s \n", input)
	}
}

func say(input string) {
	for i := 0; i < 1000; i++ {
		fmt.Printf("===== %s \n", input)
	}
}
