// 1.带缓冲功能的channel
package main

import (
	"fmt"
	"time"
)

var blockChannel = make(chan int, 0)

var bufferedChannel = make(chan int, 2)

func blockedTest(arg chan int) {
	// 对于阻塞式管道，其后续操作会被阻塞，直到阻塞式管道被消费；
	// 对于缓冲式管道则不会立即阻塞，但是到达缓冲上界后就会阻塞，直到有消费者来消费掉其中的数据
	arg <- 1 //不阻塞
	arg <- 2 //不阻塞
	arg <- 3 //会阻塞
	fmt.Println("goroutine end----")
}

func main() {
	go blockedTest(bufferedChannel)
	//<-blockChannel
	fmt.Println("out end =====")
	<-bufferedChannel
	time.Sleep(time.Second)
}
