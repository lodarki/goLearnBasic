package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"
)

func main() {
	//它内部做了很多事情：
	//
	//解析配置文件
	//
	//beego 会自动解析在 conf 目录下面的配置文件 app.conf，通过修改配置文件相关的属性，我们可以定义：开启的端口，是否开启 session，应用名称等信息。
	//
	//执行用户的 hookfunc
	//
	//beego 会执行用户注册的 hookfunc，默认的已经存在了注册 mime，用户可以通过函数 AddAPPStartHook 注册自己的启动函数。
	//
	//是否开启 session
	//
	//会根据上面配置文件的分析之后判断是否开启 session，如果开启的话就初始化全局的 session。
	//
	//是否编译模板
	//
	//beego 会在启动的时候根据配置把 views 目录下的所有模板进行预编译，然后存在 map 里面，这样可以有效的提高模板运行的效率，无需进行多次编译。
	//
	//是否开启文档功能
	//
	//根据 EnableDocs 配置判断是否开启内置的文档路由功能
	//
	//是否启动管理模块
	//
	//beego 目前做了一个很酷的模块，应用内监控模块，会在 8088 端口做一个内部监听，我们可以通过这个端口查询到 QPS、CPU、内存、GC、goroutine、thread 等统计信息。
	//
	//监听服务端口
	//
	//这是最后一步也就是我们看到的访问 8080 看到的网页端口，内部其实调用了 ListenAndServe，充分利用了 goroutine 的优势
	//
	//一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。

	//beego 默认注册了 static 目录为静态处理的目录，注册样式：URL 前缀和映射的目录（在/main.go文件中beego.Run()之前加入）：
	//
	//StaticDir["/static"] = "static"
	//用户可以设置多个静态文件处理目录，例如你有多个文件下载目录 download1、download2，你可以这样映射（在 /main.go 文件中 beego.Run() 之前加入）：
	//
	//beego.SetStaticPath("/down1", "download1")
	//beego.SetStaticPath("/down2", "download2")
	//这样用户访问 URL http://localhost:8080/down1/123.txt 则会请求 download1 目录下的 123.txt 文件。

	beego.Run()
}
