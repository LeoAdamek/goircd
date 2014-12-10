//
// GoIRCd Event System
//
// Shamelessly copied from goirc
//

package main

import (
	"fmt"
	"strings"
	"errors"
)

type Event struct {
	Code string
	Raw string
	Nick string
	Host string
	Source string
	User string
	Args []string
}


func ParseMessage(msg string) (*Event, error) {
	msg = msg[:len(msg)-2] // Trim the \r\n from the message

	event := &Event{Raw: msg}

	// Set the source if it was provided
	if msg[0] == ':' {
		if i := strings.Index(msg, " "); i > -1 {
			event.Source = msg[1:i]
			msg = msg[i+1 : len(msg)]
		} else {
			return nil, errors.New(fmt.Sprintf("Malformed message: %s", msg))
		}

		// Parse the Ident string
		if i, j := strings.Index(event.Source, "!"), strings.Index(event.Source, "@");
		i > -1 && j > -1 {
			event.Nick = event.Source[0:i]
			event.User = event.Source[i+1:j]

			event.Host = event.Source[j+1 : len(event.Source)]
		}
	}

	split := strings.SplitN(msg, " :", 2)
	args := strings.Fields(split[0])
	event.Code = strings.ToUpper(args[0])
	event.Args = args[1:]

	if len(split) > 1 {
		event.Args = append(event.Args, split[1])
	}

	return event, nil
}


// Get the "Message" argument
// Always the last argument of an IRC frame
func (e *Event) Message() string {
	if len(e.Args) == 0 {
		return ""
	}

	return e.Args[len(e.Args)-1]
}
