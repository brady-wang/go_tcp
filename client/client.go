package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	port int
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("connect server 失败",err)
		return
	}

	fmt.Println("conn success")
	// 客户端准备发送单行数据 然后退出
	reader := bufio.NewReader(os.Stdin) // 标准输入
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read content error",err)
		return
	}
	// 内容发送给服务器
	n, err := conn.Write([]byte(content))
	if err != nil {
		fmt.Println("send to server",err)
		return
	}
	fmt.Println(n)

}