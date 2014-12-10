//
// Signal Handling
//
package main

import (
	"os"
	"os/signal"
	"syscall"
	log "github.com/Sirupsen/logrus"
)


func handleSignals() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)

	for {
		s := <- sigChan

		if s == syscall.SIGINT {
			log.Infoln("Interrupt Received. Press Again to terminate.")

			s := <- sigChan

			if s == syscall.SIGINT {
				log.Infoln("2x Interrupt Received. Shutting Down.")
				os.Exit(0)
			}

		}

	}

}
