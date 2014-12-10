//
// GoIRCd
//
// IRC Message Struct
//

package main

import (
	"strings"
)

type Message struct {
	Command string
	Arguments []string
}


//
// Parse a message from a string
//
func ParseMessage(str string) Message {
	parts := strings.Fields(str)

	return Message{
		Command: parts[0],
		Arguments: parts[1:],
	}
}


