package Networking

import (
	"crypto/tls"
	"encoding/json"
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

// CreateTCPClient - factory for TCPClient
func CreateTCPClient(port int32, address string) *TCPClient {
	return &TCPClient{port, address, nil}
}

// Run runs tcp client
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

// SendCommand sends command to server
func (client *TCPClient) SendCommand(cmd *Command) bool {
	if client.conn != nil {
		encodedData, err := json.Marshal(cmd)
		if err != nil {
			log.Println("Failed to encode command with error:", err.Error())
		}

		_, err = client.conn.Write(encodedData)
		if err != nil {
			log.Println("Failed to send command with error:", err.Error())
			return false
		}
		return true
	}
	return false
}

// Shutdown - cleanups and shutdowns client tcp
func (client *TCPClient) Shutdown() {
	if client.conn != nil {
		client.conn.Close()
	}
}
