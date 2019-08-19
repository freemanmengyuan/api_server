package main

import (
	"fmt"
	"net"
	"os"
)

/**
 * 简单的聊天室
 * server 负责接受数据  client 负责发送数据
 */
func checkError(err error) {
	if err != nil {
		fmt.Printf("socket err %s", err.Error())
		os.Exit(1)
	}
}

func processInfo(conn net.Conn) {
	buffer := make([]byte, 1024)
	defer conn.Close()
	for {
		numOfBytes, err := conn.Read(buffer)
		if err != nil {
			continue
		}

		if numOfBytes != 0 {
			fmt.Println("server received message", string(buffer[0:numOfBytes]))
		}
	}

}

func main() {
	socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	defer socket.Close()

	fmt.Println("server is wait.......")
	for {
		conn, err := socket.Accept()
		checkError(err)
		go processInfo(conn)
	}
}
