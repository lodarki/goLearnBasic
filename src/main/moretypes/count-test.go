package main

import (
	"fmt"
)

func main() {
	//本金
	var origin = 34000.0000
	//日利率
	var rate = 0.0005
	//期限（年）
	var expired = 1
	var result = origin
	for i := 1; i <= expired*365; i++ {
		result += result * rate
	}
	fmt.Printf("原始金额为 %v \n", origin)
	fmt.Printf("日利率为 %v \n", rate)
	fmt.Printf("期限为 %v 年\n", expired)
	fmt.Printf("最终应还金额为 %v \n", result)
	hundredRate := ((result / origin) - 1) * 100
	fmt.Printf("年利率为百分之 %v", hundredRate/float64(expired))
}
