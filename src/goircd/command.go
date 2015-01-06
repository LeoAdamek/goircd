//
// GoIRCd Command Interface
//
// Implements commands which may be recieved by a user
//

package main


type Command interface {
	Execute(*IRCConnection, *Event)
}
