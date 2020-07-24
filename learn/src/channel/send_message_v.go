package main

import (
	"os"
	"os/exec"
	"time"
)

func execShell(file string) {
	cmd := exec.Command("php", file)
	cmd.Start()
}

func mainTemp() {
	file := os.Args[1]
	for i := 1; i <= 60; i++ {
		println(file)
		go execShell(file)
		time.Sleep(time.Second)
	}
}
