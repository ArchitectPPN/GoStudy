package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close() // 处理完之后，要关闭这个链接
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn fail, err%v\n", err)
			break
		}

		rec := string(buf[:n])
		fmt.Printf("接收到的数据：%v\n", rec)
		conn.Write([]byte("ok")) // 收到的数据写回客户端
	}
}

func main() {
	// 1. 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
	}

	for {
		// 2. 等待客户端建立链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}

		// 3. 启动一个单独的goroutine去处理连接
		go process(conn)
	}
}
