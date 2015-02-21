//
// GoIRCd
//
// Gobal Constants
//

package main

const (

	// Max Messae Size (512B)
	MAX_MSG_SIZE = 512

	// Message Terminator
	MSG_TERM = "\r\n"

	IDENT_SERVER = 0
	IDENT_USER   = 1
	
	// IRC Commands
	JOIN = "JOIN"
	PRIVMSG = "PRIVMSG"

	QUIT = "QUIT"
	PART = "PART"
	NICK = "NICK"
	USER = "USER"
	INFO = "INFO"

	PING = "PING"
	PONG = "PONG"

	LIST = "LIST"

	OPER = "OPER"

	NOTICE = "NOTICE"

	// Numeric Server Response Codes
	REP_WELCOME = "001"
	REP_TOPIC = "332"
	
	REP_LIST_START = "321"
	REP_LIST_END = "322"
	REP_LIST = "323"

)

	
