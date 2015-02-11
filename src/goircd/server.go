//
// GoIRCd Server
//
// Main Server Component for GoIRCd
//

package main


import (
	log "github.com/Sirupsen/logrus"
	"net"
)

type IRCConnectionMap map[string]IRCConnection
type IRCChannelMap map[string]IRCChannel

type IRCServer struct{
	connections IRCConnectionMap;
	channels IRCChannelMap
	service_socket Socket
}

// Create a new IRCServer Instance
func CreateIRCServer() IRCServer {
	server := IRCServer{}

	return server;
}

//
// Serve Forever
func (s *IRCServer) ServeForever() {

	s.service_socket, err = net.Listen("tcp",":6667")

	if err != nil {
		log.Errorln("Couldn't listen on *:6667.", err)
	}

	log.Infoln("Listening on *6667")

	for {
		client, err := service_socket.Accept()

		// If there's any kind of client error.
		// We don't care. Continue on.
		if err != nil {
			continue
		}

		user := NewIRCConnection(client, s)

		s.connections[user.Ident()] = user

		// Each user is handled by a goroutine.
		go user.HandleConnection()
	}

}
	
//
// Find a channel
//
func (s *IRCServer) FindChannel(chanName string) *IRCChannel {
	return s.channels[chanName]
}

// Add a Channel
func (s *IRCServer) AddChannel(channel *IRCChannel) {
	s.channels[channel.GetName()] = channel
}

func (s *IRCServer) RemoveChannel(channel *IRCChannel) {
	s.channels[channel.GetName()] = nil
}
