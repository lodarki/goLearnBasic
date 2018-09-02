package socket

import (
	"net"
	"sync"
)

var socketMaps map[string]*net.Conn

var once sync.Once

func GetSocketMaps() map[string]*net.Conn {
	once.Do(func() {
		socketMaps = make(map[string]*net.Conn)
	})
	return socketMaps
}
