//
// GoIRCd -- A distributable IRC server in Go
//

// Commentary:
//
// Handles the initialization of the IRC daemon

package main

import (
	log "github.com/Sirupsen/logrus"
)
//
// Main Function
//
// Loads the command flags
// Loads the configuration file
// Merges the flags and the config file
// Initializes Logging
//
// Forks of the daemon process (if daemonizing is enabled)
//
func main() {

	log.Infoln("Application Starting")

	// Load the command line flags
	flags.LoadFlags()

	// Load the configuration
	// Get the configuration to reload on SIGUSR1
	config.LoadFile(flags.ConfigFile)
	go config.ReloadBySignal(flags.ConfigFile)

	// Handle OS Signals
	go handleSignals()

	// Serve Forever
	ServeForever()

}
