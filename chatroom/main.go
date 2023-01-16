package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

// 创建全局map,存储在线用户  //写完了以后没有map空间
var onlineMap map[string]Client

// 创建全局channel，用来传递用户消息
var message = make(chan string)

func WriteMsgToClient(clnt Client, conn net.Conn) {
	//监听 用户自带Channel上是否有消息
	for msg := range clnt.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()

	//获取用户  网络地址 ip+port
	netAddr := conn.RemoteAddr().String()
	//创建连接用户的  结构体;  默认用户名是Ip+port
	clnt := Client{
		C:    make(chan string),
		Name: netAddr,
		Addr: netAddr,
	}
	//将新连接用户，添加到在在线用户map中
	onlineMap[netAddr] = clnt

	//创建专门用来给当前 用户发送消息的go程
	go WriteMsgToClient(clnt, conn)

	//发送 用户上线消息到 全局message 中
	message <- "[" + netAddr + "]" + clnt.Name + "login"

}
func Manager() {
	//初始化onlineMap
	onlineMap = make(map[string]Client)

	//监听全局channel中是否有数据,有数据存储至msg，无数据阻塞
	for {

		msg := <-message

		//循环发送消息给在线用户
		for _, clnt := range onlineMap {
			clnt.C <- msg

		}
	}
}

func main() {
	//创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1 8000")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listener.Close()

	//创建管理者go程，管理map和全局channel
	go Manager()

	//循环监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}

		go HandlerConnect(conn)
	}

}
