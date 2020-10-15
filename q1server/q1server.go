package main

import ( "net"
		 "fmt"
		 "bufio" )

func handleConnection(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(*conn, "kurwaaaa")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8070")
	for {
		conn, _ := ln.Accept()
		go handleConnection(&conn)
	}
}