package basic

var i, j int = 1, 2

//变量定义可以包含初始值，每个变量对应一个。
//如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
func main() {
	var c, python, java = true, false, "no!"
	println(i, j, c, python, java)
}
