package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
	for{
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read content error",err)
			return
		}
		content = strings.Trim(content,"\r\n")
		if content == "exit" {
			fmt.Println("客户端退出")
			break
		}
		// 内容发送给服务器
		_, err = conn.Write([]byte(content+"\n"))
		if err != nil {
			fmt.Println("客户端发送消息失败",err)
		}
	}


}