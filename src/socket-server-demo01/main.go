package main

import (
	"github.com/astaxie/beego"
	"net"
	_ "socket-server-demo01/routers"
	"socket-server-demo01/socket"
	"strconv"
)

func main() {
	beego.SetLogger("file", `{"filename":"D:/gologs/socketserver.log"}`)
	beego.SetLevel(beego.LevelDebug)
	beego.SetLogFuncCall(true)
	beego.Debug("beego tcpSocketListen")
	go tcpSocketListen()
	go socket.LoopReadFromSocketConn()
	go socket.LoopWriteToSocketConn()
	beego.Run()
}

/**
开始socket监听
*/
func tcpSocketListen() {
	listenPort, parseErr := beego.AppConfig.Int("socketlistenport")
	if parseErr != nil {
		beego.Error("socket listen port is not int value is :", beego.AppConfig.String("socketlistenport"))
		return
	}

	listener, listenErr := net.Listen("tcp", ":"+strconv.Itoa(listenPort))
	if listenErr != nil {
		beego.Error("socket listen failed !")
		return
	}

	for {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			continue
		}

		maps := socket.GetSocketMaps()
		maps[conn.RemoteAddr().String()] = &conn
		beego.Debug(conn)
	}
}
