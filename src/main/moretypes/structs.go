package main

import "fmt"

//一个结构体（`struct`）就是一个字段的集合。
//（而 type 的含义跟其字面意思相符。）
type Vertex struct {
	X int
	Y int
}

func (v Vertex) PrintInfo() {
	fmt.Println(v.X, v.Y)
}

func main() {
	vertex := Vertex{1, 2}
	vertex.PrintInfo()
}
