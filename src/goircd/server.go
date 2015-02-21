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

type IRCConnectionMap map[string]*IRCConnection
type IRCChannelMap map[string]*IRCChannel

type IRCServer struct{
	connections IRCConnectionMap;
	channels IRCChannelMap
	service_socket net.Listener
}

// Create a new IRCServer Instance
func CreateIRCServer() IRCServer {
	server := IRCServer{
		channels: make(IRCChannelMap),
		connections: make(IRCConnectionMap),
	}

	return server;
}

//
// Serve Forever
func (s *IRCServer) ServeForever() {

	service_socket, err := net.Listen("tcp",":6667")

	if err != nil {
		log.Errorln("Couldn't listen on *:6667.", err)
	}

	s.service_socket = service_socket

	log.Infoln("Listening on *6667")

	for {
		client, err := service_socket.Accept()

		// If there's any kind of client error.
		// We don't care. Continue on.
		if err != nil {
			continue
		}

		user := NewIRCConnection(client, s)

		// s.connections[user.String()] = &user

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

	chanName := channel.GetName()

	log.WithFields(
		log.Fields{
			"chanName": chanName,
		}).Info("Adding Channel")

	s.channels[chanName] = channel
}

func (s *IRCServer) RemoveChannel(channel *IRCChannel) {
	s.channels[channel.GetName()] = nil
}

func (s *IRCServer) ListChannels(client *IRCConnection) {

	client.SendMessage(REP_LIST_START, "Channel", "Topic")
	
	
	for i := range s.channels {
		client.SendMessage(REP_LIST, i, s.channels[i].Topic)
	}

	client.SendMessage(REP_LIST_END, "End", "of", "/LIST")
}
