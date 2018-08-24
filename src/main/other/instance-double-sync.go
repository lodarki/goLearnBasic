package other

import (
	"fmt"
	"sync"
)

var m *Manager
var lock *sync.Mutex = &sync.Mutex{}

//引入了锁的机制，在GetInstance函数中，每次调用我们都会上一把锁，保证只有一个goroutine执行它，这个时候并发的问题就解决了。
// 不过现在不管什么情况下都会上一把锁，而且加锁的代价是很大的，有没有办法继续对我们的代码进行进一步的优化呢？ 熟悉java的同学可能早就想到了双重的概念，没错，在go中我们也可以使用双重锁机制来提高效率。

func GetInstance() *Manager {
	if m == nil {
		lock.Lock()
		defer lock.Unlock()
		if m == nil {
			m = &Manager{}
		}
	}
	return m
}

type Manager struct{}

func (p Manager) Manage() {
	fmt.Println("manage...")
}
