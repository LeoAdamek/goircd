//
// GoIRCd IRC Connection
//
package main

import (
	"net"
	"strings"
	log "github.com/Sirupsen/logrus"
)

type IRCConnection struct {

	/* Public */
	User string
	Nick string
	HasBeenWelcomed bool

	/* Private */
	conn net.Conn
	events chan *Event
	callbacker Callbacker
	server *IRCServer
	channels IRCChannelMap
}


func NewIRCConnection(connection net.Conn, server *IRCServer) IRCConnection {
	return IRCConnection{
		conn: connection,
		events: make(chan *Event, 32),
		server: server,
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
	
	c.callback(event)

}

// Send a message to a client
func (c *IRCConnection) SendMessage(parts ...string) {
	var reply string

	reply = ":localhost "
		
	if len(parts) > 1 {
		reply += strings.Join(parts[:len(parts)-1], " ")
		reply += " :" + parts[len(parts)-1]
	} else {
		reply += parts[0]
	}
	
	reply += MSG_TERM

	log.WithFields(
		log.Fields{
			"ident" : c.String(),
			"message" : reply,
		}).Info("Sent Reply")
	
	c.conn.Write([]byte(reply))

}

////////// IRC MESSAGES

//
// Pong
func (c *IRCConnection) Pong(message string) {

	if len(message) > 0 {
		c.SendMessage(PONG, message)
	} else {
		c.SendMessage(PONG)
	}
}

//
// Welcome
func (c *IRCConnection) Welcome() {
	c.SendMessage(REP_WELCOME, c.User, "Welcome to the GoIRCd Server")
}

//
// Info
//
// Sends the client their connection info
func (c *IRCConnection) Info() {
	c.SendMessage(NOTICE, "Connection Information for " + c.String() + ": ")
}

//
// Join
//
// Joins (or creates) a channel
func (c *IRCConnection) Join(e *Event) {
	chanName := e.Message()
	channel := c.server.GetChannel(chanName)

	if channel == nil {
		channel = NewIRChannel(chanName)
		c.server.AddChannel(channel)
	}

	c.SendMessage(JOIN, ' ', chanName)

	c.channels[chanName] = channel
}

//
// Stringify Function for an IRCConnection
//
// Returns a string representation of the IRCConnection
func (c *IRCConnection) String() string {
	return c.Nick + "!" + c.User + "@" + c.conn.RemoteAddr().String()
}
