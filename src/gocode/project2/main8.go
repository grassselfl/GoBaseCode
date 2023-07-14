package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	//循环处理客户端发送的数据
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data) //此处会阻塞
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接")
			}
			return
		}
		fmt.Println(string(data[:n])) //此处一定要截一下艾，否则会有问题
	}
}

func main() {
	fmt.Println("服务端启动")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败", err)
	}

	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("监听客户端失败", err2)
		} else {
			fmt.Printf("conn = %v, msg = %v\n", conn, conn.RemoteAddr().String())
			fmt.Println("#{err2}, ")
			go process(conn)
		}
	}
}
