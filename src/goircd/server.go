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

//
// Serve Forever
func ServeForever() {

	server, err := net.Listen("tcp",":6667")

	if err != nil {
		log.Errorln("Couldn't listen on *:6667.", err)
	}

	log.Infoln("Listening on *6667")

	for {
		client, err := server.Accept()

		// If there's any kind of client error.
		// We don't care. Continue on.
		if err != nil {
			continue
		}

		user := NewIRCConnection(client)

		// Each user is handled by a goroutine.
		user.HandleConnection()
	}

}
	
