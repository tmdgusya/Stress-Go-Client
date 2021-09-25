package connection

import (
	"log"
	"net"
)

const SERVER_DOMAIN = "http://localhost"
const SOCKET_SERVER_PORT = ":8080"

func ConnectFactory() (net.Conn, error) {
	conn, err := net.Dial("tcp", SERVER_DOMAIN+SOCKET_SERVER_PORT)
	if err != nil {
		log.Println(err)
	}
	return conn, err
}