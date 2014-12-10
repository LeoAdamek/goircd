//
// Commentary:
//
// Handles the loading of the command line flags
//
package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
)

type FlagOptions struct {
	ConfigFile string
}

var flags FlagOptions

//
// Load the flags
//
// Returns:
// true - Success
// false - Error
func (f *FlagOptions) LoadFlags() bool {

	log.Infoln("Loading Options")
	
	flag.StringVar(&f.ConfigFile, "config", "config.json", "Configuration File")

	flag.Parse()

	return true
}

