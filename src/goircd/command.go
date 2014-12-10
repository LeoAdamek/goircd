//
// GoIRCd Command Interface
//

package main


type Command interface {
	Execute(*IRCConnection, *Event)
}

