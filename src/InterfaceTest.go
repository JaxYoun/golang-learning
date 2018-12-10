//Golang中的接口实现采用的是非侵入式的方式，即不需要像Java那样显式的表明实现关系，
//只要某个类实现了一个接口的所有方法，这个类就实现了该接口，就具有部分接口的多态性
//值得注意的式实现的方法必须为拷贝绑定型。

//当两个接口中声明的方法呈现全包含关系时，可以将方法多的接口类型的变量赋值给方法少的接口类型变量
package main

import "fmt"

type Runnable interface {
	run(far int)
	fly(hight int)
}

type Callable interface {
	fly(age int)
}

type Duck struct {
	name string
}

func (this Duck) run(far int) {
	fmt.Println(this.name)
}

func (this Duck) fly(far int) {
	fmt.Println(this.name)
}

func (this Duck) call(far int) {
	fmt.Println(this.name)
}

func main() {
	duck := Duck{"ga ga ga ga..."}
	duck.run(12)

	var runner Runnable
	runner = duck
	runner.run(333)

	//满足接口的多态性
	var runne2 Runnable = Duck{"========="}
	runne2.fly(999)

	//满足接口的多态性
	hhh(duck)

	//满足接口的多态性
	hhh(runne2)

	//将接口变量赋值给接口变量
	var callable Callable = runne2
	callable.fly(89)
}

func hhh(runner Runnable) {
	runner.fly(12)
}
