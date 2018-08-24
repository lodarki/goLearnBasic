package main

import "fmt"

//一个结构体（`struct`）就是一个字段的集合。
//（而 type 的含义跟其字面意思相符。）
//type Vertex struct {
//	X int
//	Y int
//}

func main() {
	v := Vertex{X: 1, Y: 2}
	v.X = 4
	fmt.Println(v.X)
}
