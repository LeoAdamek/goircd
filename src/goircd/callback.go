//
// GoIRCd
//
// Callback Interface
//
package main

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
	switch e.Code {
	case PING:
		c.Pong(e.Message())
	case NICK:
		if !c.HasBeenWelcomed {
			c.Welcome()
		}
	}
}
