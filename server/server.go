package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, _ := ln.Accept()
		//if err != nil {
		//	handleError(err)
		//}
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	fmt.Println("SERVER: connected User: ", clientid)
	for {
		reader := bufio.NewReader(client)
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err, ", couldn't read string from client")
		}
		msgs <- Message{clientid, fmt.Sprint("User ", clientid, ": ", msg)}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)
	fmt.Println("SERVER: listening on ", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			clientid := len(clients)
			clients[clientid] = conn
			go handleClient(conn, clientid, msgs)
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i, conn := range clients {
				if i != msg.sender {
					fmt.Fprint(conn, msg)
				} else {
					fmt.Fprintln(conn, "SERVER: message was received by server")
				}
			}
		}
	}
}
