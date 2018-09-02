package main

import (
	"github.com/astaxie/beego"
	"net"
	_ "socket-client-demo01/routers"
	"socket-client-demo01/socket"
	"time"
)

func main() {
	beego.SetLogger("file", `{"filename":"D:/gologs/socketclient.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	beego.Run()
	socketDia()
}

func socketDia() {
	socketServerHost := beego.AppConfig.String("sockettargethost")
	if socketServerHost == "" {
		beego.Error("socket server host is not exist!")
		return
	}

	for {
		conn, dialErr := net.Dial("tcp", socketServerHost)
		if dialErr != nil {
			beego.Error("socket dia error :", dialErr.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		socket.SocketConn = &conn
		beego.Debug(*socket.SocketConn)
		break
	}
}
