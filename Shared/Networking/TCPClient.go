package Networking

import (
	"crypto/tls"
	"log"
	"net"
	"strconv"
)

// TCPClient creates connection with monitoring server, holds main connection
type TCPClient struct {
	port    int32
	address string
	conn    net.Conn
}

func CreateTCPClient(port int32, address string) *TCPClient {
	return &TCPClient{port, address, nil}
}

func (client *TCPClient) Run() bool {
	cert, err := tls.LoadX509KeyPair("../certs/client.pem", "../certs/client.key")
	if err != nil {
		log.Println("Failed to load certs", err)
		return false
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", client.address+":"+strconv.Itoa(int(client.port)), &config)
	if err != nil {
		log.Println("Failed to connect with error:", err.Error())
		return false
	}
	client.conn = conn
	return true
}

func (client *TCPClient) SendCommand(command []byte) bool {
	if client.conn != nil {
		_, err := client.conn.Write(command)
		if err != nil {
			log.Println("Failed to send command with error:", err.Error())
			return false
		}
		return true
	}
	return false
}

func (client *TCPClient) Shutdown() {
	client.conn.Close()
}
