package basic

import "fmt"

//类型推导
func main() {
	v := 42
	c := 3.142
	d := 0.876 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("v is of type %T\n", c)
	fmt.Printf("v is of type %T\n", d)
}
