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
	User string
	Nick string
	Host string
	connection net.Conn
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
		user := IRCUser{
			connection: client,
		}

		go user.Handle()
		
	}
	
}


// Handles a client connection
func (user *IRCUser) Handle() {
	
	client_recv     := make([]byte, 8192)   /* Received client data buffer (8kiB) */
	var client_recv_len int                 /* Length of received client data frame */
	var err error                           /* Error */

	for {
		// Clear the client_recv slice:
		client_recv = make([]byte, 8192)
		
		client_recv_len, err = user.connection.Read(client_recv)

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
		go user.handleCommand(command, params)

	}

	user.connection.Close()
}

// Handles a single command issued by a client
func (user *IRCUser) handleCommand(command string, params []string) {

	var reply string

	if command == "USER" {
		user.User = lastParam(params)
	}

	if command == "NICK" {
		user.Nick = lastParam(params)

		log.Println("D  User nick:", user.Nick)
		
		reply = "001 " + user.Nick + " :Welcome to the IRC chat."
	}

	if command == "PING" {
		reply = "PONG"
	}

	if command == "JOIN" {
		reply = ":" + user.Ident() + " " + "JOIN" + lastParam(params)

		log.Println("D Reply..." , reply)
		user.connection.Write([]byte(reply + MESSAGE_TERMINATOR))
		return
	}

	if len(reply) > 0 {
		log.Println("D Sending reply... ", SERVER_IDENT, reply)
		user.connection.Write([]byte(SERVER_IDENT + reply + MESSAGE_TERMINATOR))
	}
}

// Get's the user's ident
func (user *IRCUser) Ident() string {
	return user.Nick + "!" + user.User + "@" + user.connection.RemoteAddr().String()
}


// Gets the last parameter from a slice (including spaces)
// Removes the prepended :
func lastParam(params []string) string {
	for i := range params {
		if params[i][0] == ':' {
			return strings.Replace(
				strings.Join(params[i:], " "),
				":", "", 1)
		}
	}

	return params[len(params)-2]
}
