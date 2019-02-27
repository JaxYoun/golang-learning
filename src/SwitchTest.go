//Golang中的switch与Java中的作用相似，但是存在某些qubie
//1.switch后支持所有的基本数据类型，ing，float，bool，string等
//2.switch后可以不带参数或表达式，此时相当于switch true，只执行为true的case，此时的case后需要写bool表达式来达到分支的目的，类似if-else语句
//3.case子句后可以带多个条件，表示只要满足这多个中的一个条件就可以执行这个分支
//4.case子句中的break不是强制的【但，如果逻辑需要，可以使用break关键字达到跳出本分支，并跳出本switch的目的】，默认只会执行一个分支而不会出现条件击穿问题
//5.跟Java一样，如果没有任何条件满足，就会执行default分支
//6.如果想要达到人为击穿的效果，可以在case间使用fallthrough关键字，让上一个case的隐式阻断失效
package main

import "fmt"

func main() {
	a := 11

	//1.switch后跟变量，此时case后的值类型必须同这个变量类型相同
	switch a {
	case 10:
		fmt.Println("ten") //隐式阻断，会被fallthrough关键字破坏
	case 11:
		fmt.Println("eleven")
		break //显式阻断，不会被fallthrough关键字破坏
	default:
		fmt.Println("none matchs!")
	}

	//2.switch后不跟变量，此时case后只能跟bool表达式
	switch {
	case a > 10:

	case a < 100:

	}

}
