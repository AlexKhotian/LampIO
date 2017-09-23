package Networking

import (
	"log"
	"net"
	"strconv"
)

// TCPClient creates connection with monitoring server, holds main connection
type TCPClient struct {
	port    int32
	address *string
}

func CreateTCPClient(port int32, address *string) *TCPClient {
	return &TCPClient{port, address}
}

func (client *TCPClient) Run() bool {
	conn, err := net.Dial("tcp", *client.address+":"+strconv.Itoa(int(client.port)))
	if err != nil {
		log.Println("Failed to connect with error:", err.Error())
		return false
	}

	return true
}
