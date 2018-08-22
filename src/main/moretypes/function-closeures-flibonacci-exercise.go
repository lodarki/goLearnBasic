package main

import "fmt"

//现在来通过函数做些有趣的事情。
//实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	var ints []int
	return func() int {
		length := len(ints)
		if length == 0 {
			ints = append(ints, 0)
		}
		if length == 1 {
			ints = append(ints, 1)
		}
		if length >= 2 {
			ints = append(ints, ints[length-1]+ints[length-2])
		}
		return ints[len(ints)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
