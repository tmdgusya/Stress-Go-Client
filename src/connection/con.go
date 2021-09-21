package connection

import (
	"log"
	"net"
)

const SOCKET_SERVER_PORT = ":8080"

func ConnectFactory() net.Conn {
	conn, err := net.Dial("tcp", SOCKET_SERVER_PORT)
	if err != nil {
		log.Println(err)
	}
	return conn
}
