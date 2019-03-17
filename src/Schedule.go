package main

//timer和ticker测试
import (
	"fmt"
	"time"
)

func main() {
	//timerTest()
	//loopTimerTest()
	//tickerTest()
	//selectTickerTest()
	producerAndConsummerTest()

	time.Sleep(time.Minute * 5)
}

/*timer只会在预设时间段到达后执行一次，到达之前可以重置，可以关闭
其机制是当到达预定时间后，timer实例会向其内置channel C写入一个当前时间，来解除阻塞，触发执行
由于只执行一次，所以要达到ticker那样的周期性执行就需要引入循环体，在循环体中反复的调用timer的reset方法
*/
func simplelTimerTest() {
	//1.实例化一个timer
	timer := time.NewTimer(5 * time.Second)

	//2.重置到期时间
	timer.Reset(5 * time.Second)

	//4.停止
	//timer.Stop()

	//3.利用timer内置channel实现阻塞
	<-timer.C
	go func() {
		fmt.Println("timer now")
	}()
}

func loopTimerTest() {
	timer := time.NewTimer(time.Second * 3)
	for {
		select {
		case <-timer.C:
			{
				go func() {
					fmt.Println("loopTimerTest")
				}()
				timer.Reset(time.Second * 3)
			}

		//不确定此处的default是否必要
		default:
			{
				continue
			}
		}
	}
}

/*ticker会在预设的周期到达后反复执行（会顺序推迟），只支持停止，不可重置
由于推迟的问题，所以尽量每一次执行都在单独的协程中执行
*/
func simpleTickerTest() {
	//1.实例化一个定时调度其
	ticker := time.NewTicker(3 * time.Second)

	//2.停止调度器
	ticker.Stop()

	//3.周期性阻塞-触发执行
	for range ticker.C {
		fmt.Println("ticker")
	}
}

func selectTickerTest() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			{
				go func() {
					fmt.Println("----------")
				}()
			}
		}
	}
}

func producerAndConsummerTest() {
	//1.声明一个容量为100的通道
	input := make(chan int, 10)

	//2.启动周期性生产
	go tickerProducer(&input)

	//3.启动周期性消费
	go tickerConsumer(&input)

	// 4.简单消费
	//simpleConsumer(&input)
}

func tickerProducer(input *chan int) {
	index := 0
	ticker := time.NewTicker(time.Second * 3)
	for range ticker.C {
		fmt.Println(index)
		*input <- index
		index++
	}
}

func tickerConsumer(input *chan int) {
	ticker := time.NewTicker(time.Second * 3)
	for range ticker.C {
		fmt.Printf("========== %d \n", <-*input)
	}
}

func simpleConsumer(input *chan int) {
	/*for i := range *input {
		fmt.Printf("+++++++ %d \n", i)
	}*/

	for {
		select {
		case i := <-*input:
			{
				fmt.Printf("+++++++ %d \n", i)
			}
		}
	}
}
