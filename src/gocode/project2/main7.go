package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端连接失败，", err)
		return
	}
	fmt.Println("连接成功，", conn)

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	n, err := conn.Write([]byte(str))
	if err != nil {
		return
	}
	fmt.Printf("客户端发送了%c字节数据并退出了\n", n)

}
