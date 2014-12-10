//
// GoIRCd IRC Connection
//
package main

import (
	"net"
	log "github.com/Sirupsen/logrus"
)

type IRCConnection struct {
	conn net.Conn
	User string
	Nick string
	events chan *Event
}


func NewIRCConnection(connection net.Conn) IRCConnection {
	return IRCConnection{
		conn: connection,
		events: make(chan *Event, 32),
	}
}

func (c *IRCConnection) HandleConnection() {
	client_recv := make([]byte, 8192)

	var client_recv_len int
	var err error

	for {
		// Read data
		client_recv_len, err = c.conn.Read(client_recv)

		// Check if data was received
		if err != nil && client_recv_len == 0 {

			log.Infoln("Client ", c, " Disconnected")
			
			// No data Received. Disconnect
			break
		}

		// Handle the message...
		go c.HandleMessage(string(client_recv))

		// Reset the recv buffer.
		client_recv = make([]byte, 8192)
	}

	// When the loop finishes
	// Close the connection
	c.conn.Close()
}


//
// Handle a message
//
// Handles a single message
func (c *IRCConnection) HandleMessage(message string) {

	event, err := ParseMessage(message)

	if err != nil {
		log.Errorln("Error Parsing Message: ", err)
		return
	}

	log.WithFields(
		log.Fields{
			"command" : event.Code,
			"user" : event.User,
			"args" : event.Args,
		}).Info("Client Command")
	
	c.events <- event

}

//
// Stringify Function for an IRCConnection
//
// Returns a string representation of the IRCConnection
func (c *IRCConnection) String() string {
	return c.Nick + "!" + c.User + "@" + c.conn.RemoteAddr().String()
}
