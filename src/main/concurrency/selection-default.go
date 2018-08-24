package main

import (
	"fmt"
	"time"
)

//当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
//
//为了非阻塞的发送或者接收，可使用 default 分支：
//
//  select {
//  case i := <-c:
//  // 使用 i
//  default:
//  // 从 c 读取会阻塞
//  }
func main() {
	tick := time.Tick(100 * time.Millisecond)  // 返回的是一个channel
	boom := time.After(500 * time.Millisecond) // 返回的是一个channel
	for {
		select {
		case i := <-tick:
			fmt.Println("tick.", i)
		case j := <-boom:
			fmt.Println("BOOM!", j)
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
