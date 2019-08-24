package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

/**
 * 简单的聊天室
 * server 负责接受数据  client 负责发送数据
 * 支持高并发 QPS可以达到10W
 * 定义消息体 记录所有的连接 实现转化功能
 */

var onlineConnect = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quitChan = make(chan bool)

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
			break
		}
		if numOfBytes != 0 {
			//将消息写入消息队列
			messageQueue <- string(buffer[0:numOfBytes])
		}
	}

}

//消费队列
func ConsumerMessage() {
	for {
		select {
		case message := <-messageQueue:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

//解析消息
func doProcessMessage(message string) {
	content := strings.Split(message, "#")
	if len(content) > 1 {
		addr := content[0]
		sendMessage := content[1]
		if conn, ok := onlineConnect[addr]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {
				fmt.Print("online conn fail")
			}

		}
	}
}

func main() {
	//主协程建立tcp连接 监听8080端口
	socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkError(err)
	defer socket.Close()

	fmt.Println("server online.......")
	//消费消息
	go ConsumerMessage()

	//阻塞 开启多个协程接受客户端消息
	for {
		conn, err := socket.Accept()
		checkError(err)
		//将所有的connect存储到onlineConnect映射表里
		conaddr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConnect[conaddr] = conn
		fmt.Println("客户端即时统计列表")
		for i := range onlineConnect {
			fmt.Println("客户端：" + i + " 上线")
		}
		go processInfo(conn)
	}
}
