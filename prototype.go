//
// IRC Server Prototype
//
// Ultra-minimal prototype for proof-of-concept.
//

package main

import (
	"log"
	"net"
	"strings"
)

const (
	MESSAGE_TERMINATOR = "\r\n"
	SERVER_IDENT = ":localhost "
	SERVER_HOST = "localhost"
)


type IRCUser struct {
	Name string
	Ident string
	Nick string
	Remote net.Addr
}

// Main start-up
//
// Creates the listening server and starts the app
func main() {

	log.Println("I  Starting")

	// Listen on *:6667 (standard for IRC)
	server, err := net.Listen("tcp",":6667")

	// If we can't listen, bail out.
	if err != nil {
		log.Fatalln("F  Couldn't listen on *:6667 ", err)
	}

	log.Println("I   Accepting Client Connections")

	for {
		client, err := server.Accept()

		// If the client has an error
		// We don't care, just move on
		if err != nil {
			continue
		}

		// handle the client connection in a goroutine
		go handleClient(client)

	}
	
}


// Handles a client connection
func handleClient(client net.Conn) {

	user := IRCUser{
		Remote: client.RemoteAddr(),
	}
	
	client_recv     := make([]byte, 8192)   /* Received client data buffer (8kiB) */
	var client_recv_len int                 /* Length of received client data frame */
	var err error                           /* Error */

	for {
		client_recv_len, err = client.Read(client_recv)

		// Check for, and disconnect broken/disconnected clients
		if err != nil && client_recv_len == 0{
			break
		}

		log.Println("D Client Message...", string(client_recv))

		// Parse what the user sent (in its simplist form
		fields := strings.Fields(string(client_recv))
		command := fields[0]
		params := fields[1:]

		// If the client command was QUIT
		// Then break and disconnect them
		if command == "QUIT" {
			break
		}

		// Handle each command in a goroutine so that they are non-blocking.
		// We can still send the client data from within the goroutine
		// so they can still get the response to their command.
		go user.handleCommand(client, command, params)

	}

	client.Close()
}

// Handles a single command issued by a client
func (user *IRCUser) handleCommand(client net.Conn, command string, params []string) {

	var reply string

	if command == "USER" {
		user.Ident = params[len(params)-1]
		reply = "001 leo\r\n002 localhost goircd-dev\r\n003 created today\r\n\004 new"
	}

	if command == "NICK" {
		user.Nick = params[len(params)-1]
	}

	if len(reply) > 0 {
		log.Println("D Sending reply... ", SERVER_IDENT, reply)
		client.Write([]byte(SERVER_IDENT + reply + MESSAGE_TERMINATOR))
	}

}
