package socket

import (
	"github.com/astaxie/beego"
	"net"
	"time"
)

var SocketConn *net.Conn

func ReadFromSocket(b *[]byte) (int, error) {
	return (*SocketConn).Read(*b)
}

func WriteToSocket(b []byte) (int, error) {
	return (*SocketConn).Write(b)
}

func LoopPrintReadSocket() {
	readBuffer := make([]byte, 1024)
	for {
		length, e := ReadFromSocket(&readBuffer)
		if e != nil {
			sleep5s()
			continue
		}
		beego.Info("read from socket: ", string(readBuffer[:length]))
		sleep5s()
	}
}

func sleep5s() {
	time.Sleep(time.Duration(5) * time.Second)
}

func LoopWriteSocket() {
	var sendStr = "message from socket client"
	for {
		_, e := WriteToSocket([]byte(sendStr))
		sleep5s()
		if e != nil {
			continue
		}
	}
}
