//
// GoIRCd
//
// Callback Interface
//
package main

import(
	log "github.com/Sirupsen/logrus"
)

type Callbacker interface {
	Callback(*IRCConnection, *Event)
}


type CallbackFunc func(*IRCConnection, *Event)

// Callback Function Executor
func (cf CallbackFunc) Callback(c *IRCConnection, e *Event) {
	cf(c,e)
}


func (c *IRCConnection) Callbackerfunc(f func(*IRCConnection, *Event)) {
	c.callbacker = CallbackFunc(f)
}

func (c *IRCConnection) Callbacker(cf CallbackFunc) {
	c.callbacker = cf
}

func (c *IRCConnection) callback(e *Event) {

	log.WithFields(
		log.Fields{
			"client" : c.conn.RemoteAddr().String(),
			"raw" : e.Raw,
			"args" : e.Args,
			"message" : e.Message(),
		}).Debug("Event")
	
	switch e.Code {
	case PING:
		c.Pong(e.Message())
	case NICK:
		if len(e.Args) > 2 {

			if e.Args[1] == "USER" {
				c.User = e.Args[2]
			}
		}

	        c.Nick = e.Args[0]

		if !c.HasBeenWelcomed {
			c.Welcome()
		}
		
	case USER:
		c.User = e.Message()
		
	case OPER:
		c.SendMessage(NOTICE, "Oper Status Requested...")

	case JOIN:
		c.Join(e)
		
	case INFO:
		c.Info()

	case PRIVMSG:
		c.PrivateMessage(e)

	case LIST:
		Server.ListChannels(c)
	}

}
