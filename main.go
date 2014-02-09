package main

import (
	"bufio"
	"emrysai"
	"fmt"
	"hospitalsql"
	"log"
	"os"
)

func startServer(value chan bool) {
	fmt.Print("start server\n")
	hospitalsql.GetInstance() //初始化数据库
	fmt.Print("start server init ok\n")
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}
		log.Println("command", command)
	}
	value <- true
	fmt.Print("exit server")

}

func showMd5Data() {
	fmt.Println(string(emrysai.GetMD5Data("goweb")))
}

func main() {
	// done := make(chan bool)
	// startServer(done)
	// <-done
	showMd5Data()

}
