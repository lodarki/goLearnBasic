package socket

import (
	"github.com/astaxie/beego"
	"time"
)

func ReadFromSocketConn() {
	maps := GetSocketMaps()
	for key, value := range maps {
		conn := *value
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			continue
		}
		beego.Info("read from socket Conn : ", key, " content: ", string(b))
	}
}

func WriteToSocketConn(b []byte) {
	maps := GetSocketMaps()
	for key, value := range maps {
		conn := *value
		_, err := conn.Write(b)
		if err != nil {
			continue
		}
		beego.Info("write to socket Conn: ", key, " content: ", string(b))
	}
}

func LoopReadFromSocketConn() {
	for {
		ReadFromSocketConn()
		Sleep5S()
	}
}

func LoopWriteToSocketConn() {
	var sendStr = "this is the message send to socket client hei bro"
	for {
		WriteToSocketConn([]byte(sendStr))
		Sleep5S()
	}
}

func Sleep5S() {
	time.Sleep(time.Duration(5) * time.Second)
}
