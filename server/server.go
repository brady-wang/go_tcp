package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server starting ...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")

	if err != nil {
		fmt.Println("listen error",err)
		return
	}

	// 里面有个acctpt等待连接
	defer listen.Close()

	// 循环等待客户端连接
	for{
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error",err)
			return
		} else {
			fmt.Printf("accept success client ip = %v\n",conn.RemoteAddr().String())
		}

	}
}