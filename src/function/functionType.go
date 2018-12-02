package main

import "fmt"

type predicate func(num int) bool

type numFunc func([]int, predicate) []int

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(getSmall(slice, getEvenSliceFromBigSlice, isEven))
	fmt.Println(getSmall(slice, getEvenSliceFromBigSlice, isOdd))
}

func getEvenSliceFromBigSlice(big []int, pre predicate) (small []int) {
	for _, v := range big {
		if pre(v) {
			small = append(small, v)
		}
	}
	return
}

func isEven(num int) bool {
	return num%2 == 0
}

func isOdd(num int) bool {
	return num%2 != 0
}

func getSmall(big []int, fun numFunc, pre predicate) []int {
	return fun(big, pre)
}
