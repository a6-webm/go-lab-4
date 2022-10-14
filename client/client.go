package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err, ", couldn't read from server")
			time.Sleep(2 * time.Second)
		}
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for {
		msg, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println(err, ", couldn't read from stdin")
		}
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err, ", couldn't write to server")
		}
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	fmt.Println("Dialing ", *addrPtr)
	conn, err := net.Dial("tcp", *addrPtr)
	for err != nil {
		fmt.Println(err, ", couldn't reach server, trying again in 2 seconds")
		time.Sleep(2 * time.Second)
		conn, err = net.Dial("tcp", *addrPtr)
	}

	fmt.Println("Connected to ", *addrPtr)
	go write(conn)
	go read(conn)
	for {
	}
}
