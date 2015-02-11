//
// IRC Channel
//
package goircd


type IRCChannel struct{
	string name
	connections IRCConnectionMap
	events chan *Event
}


func CreateIRChannel(chanName string, chanOwner *IRCConnection) IRCChannel {
	channel := IRCChannel{}
}

// We have to protect the name from changing
func (c *IRCChannel) GetName() string {
	return c.name
}


// Add a connection
func (c *IRCChannel) AddConnection(conn *IRCConnection) error {
	c.connections[conn.Ident()] = conn

	return nil
}

// Remove a connection
func (c *IRCChannel) RemoveConnection(conn *IRCConnection) error {
	c.connections[conn.Ident()] = nil

	return nil
}
