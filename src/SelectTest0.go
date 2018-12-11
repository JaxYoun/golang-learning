//利用channel实现超时控制
package main

import (
	"fmt"
	"time"
)

var timeOut = make(chan bool)

//1.利用协程和channel并搭配select实现超时控制
func chanelTimeout(chan bool) {
	time.Sleep(time.Second)
	timeOut <- true
}

func main() {
	go chanelTimeout(timeOut)
	select {
	case val := <-timeOut:
		fmt.Println("时间到了===：", val)
	}

	//2.利用time.After函数搭配select实现超时机制
	select {
	case val := <-time.After(3 * time.Second):
		fmt.Println(val.Date())
	}
}
