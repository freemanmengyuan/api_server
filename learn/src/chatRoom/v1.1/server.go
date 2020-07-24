package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 简单的聊天室
 * server 负责接受数据  client 负责发送数据
 * 支持高并发 QPS可以达到10W
 */

//处理错误
func checkError(err error) {
	if err != nil {
		fmt.Printf("socket err %s", err.Error())
		os.Exit(1)
	}
}

//接受客户端请求
func processInfo(conn net.Conn) {
	buffer := make([]byte, 1024)
	defer conn.Close()
	for {
		numOfBytes, err := conn.Read(buffer)
		if err != nil {
			continue
		}
		clientaddress := conn.RemoteAddr()
		if numOfBytes != 0 {
			fmt.Printf("has received message {%s} from client ip :%s \n",
				string(buffer[0:numOfBytes]), clientaddress)
		}
	}

}

func main() {
	//主协程建立tcp连接 监听8080端口
	socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	defer socket.Close()

	fmt.Println("server is wait.......")
	//阻塞 开启多个协程接受客户端消息
	for {
		conn, err := socket.Accept()
		checkError(err)
		go processInfo(conn)
	}
}
