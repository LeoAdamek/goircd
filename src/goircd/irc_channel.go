//
// IRC Channel
//
package main

import (
//	log "github.com/Sirupsen/logrus"
)

type IRCChannel struct{
	name string
	connections IRCConnectionMap
	owner *IRCConnection
	events chan *Event

	Topic string
}


func CreateIRCChannel(chanName string, chanOwner *IRCConnection) IRCChannel {
	channel := IRCChannel{
		name: chanName,
		owner: chanOwner,
		events: make(chan *Event),
		connections: make(IRCConnectionMap),
		Topic: "Example Topic",
	}

	return channel
}

// We have to protect the name from changing
func (c *IRCChannel) GetName() string {
	return c.name
}


// Add a connection
func (c *IRCChannel) AddConnection(conn *IRCConnection) error {
	c.connections[conn.String()] = conn
	return nil
}

// Remove a connection
func (c *IRCChannel) RemoveConnection(conn *IRCConnection) error {
	c.connections[conn.String()] = nil

	return nil
}
