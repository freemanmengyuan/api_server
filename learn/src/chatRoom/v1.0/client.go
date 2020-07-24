package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("socket err %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	checkErr(err)
	defer conn.Close()

	conn.Write([]byte("hello world"))

	fmt.Println("client hased send message")

}
