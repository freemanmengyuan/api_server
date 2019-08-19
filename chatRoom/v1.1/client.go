package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("socket err %s", err.Error())
		os.Exit(1)
	}
}

func sendMessage(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)
		if strings.ToUpper(input) == "EXIST" {
			conn.Close()
			break
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			conn.Close()
			fmt.Println("client send message fail")
			break
		}
	}
}

func main() {

	//建立tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkErr(err)
	defer conn.Close()

	fmt.Println("client init.....")
	//协程监听键盘 发送消息
	go sendMessage(conn)

	fmt.Println("client wait.....")
	//阻塞等待服务端信息
	buffer := make([]byte, 1024)
	for {
		numOfBytes, err := conn.Read(buffer)
		checkErr(err)
		if numOfBytes != 0 {
			fmt.Println("received server message", string(buffer[0:numOfBytes]))
		}
	}

}
