package basic

import (
	"fmt"
	"math"
)

//类型转换
func main() {
	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = int(f)
	fmt.Println(x, y, z)
}
