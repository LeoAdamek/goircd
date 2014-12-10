// Commentary:
//
//   Configuration file loading
//

package main

import(
	"sync"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os/signal"
	"os"
	"encoding/json"
	"syscall"
)

type Configuration struct {
	Mode string
	CacheSize int
}


var (
	config *Configuration
	configLock = new(sync.RWMutex)
)

func (config *Configuration) LoadFile(filePath string) bool {

	data , err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Errorln("Error Loading Config:", err)
		return false
	}

	tmpConfig := new(Configuration)

	if err = json.Unmarshal(data, tmpConfig); err != nil {
		log.Errorln("Error Parsing Config: ", err)
		return false
	}

	configLock.Lock()
	config = tmpConfig
	configLock.Unlock()

	return true
}

//
// Trigger a configuration reload by a SIGUSR1
// To be called as a goroutine
//
func (config *Configuration) ReloadBySignal(filePath string) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR1)

	for {
		<- s
		config.LoadFile(filePath)
		log.Infoln("Reloaded Configuration")
	}
}

