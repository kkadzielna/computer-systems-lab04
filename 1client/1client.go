package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(msg)
}

func main() {
	msgPtr := flag.String("msg", "default", "msg to send")
	flag.Parse()
	conn, _ := net.Dial("tcp", "127.0.0.1:8010")
	fmt.Fprint(conn, msgPtr)
	read(&conn)
}
