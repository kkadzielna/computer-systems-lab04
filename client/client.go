package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func read(conn *net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.

	//might not work
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Print(msg)
}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.

}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, _ := net.Dial("tcp", "127.0.0.1:8070")
	for {
		fmt.Println("Enter text: ")
		fmt.Fprintf(conn, *addrPtr)
		read(&conn)
	}
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
