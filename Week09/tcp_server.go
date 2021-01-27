package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go process(conn)
	}
}

var (
	messageCh = make(chan string)
)

func process(conn net.Conn) {
	defer conn.Close()
	log.Println(fmt.Sprintf("来自%v的链接", conn.RemoteAddr().String()))
	rd := bufio.NewReader(conn)
	go sendMsg(conn)
	for {
		buf, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error:%v\n", err)
			return
		}
		str := string(buf)

		if str == "exit\n" {
			break
		}
		messageCh <- "hello " + str
	}
	close(messageCh)
}

func sendMsg(conn net.Conn) {
	for msg := range messageCh {
		log.Println(msg)
		conn.Write([]byte(msg + "\n"))
	}
}
