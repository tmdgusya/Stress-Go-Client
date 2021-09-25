package connection

import (
	"log"

	"github.com/go-stomp/stomp"
)

const SERVER_DOMAIN = "http://localhost"
const SOCKET_SERVER_PORT = ":8080"

func ConnectFactory() (*stomp.Conn, error) {
	conn, err := stomp.Dial("tcp", SERVER_DOMAIN+SOCKET_SERVER_PORT)
	if err != nil {
		log.Println(err)
	}
	return conn, err
}