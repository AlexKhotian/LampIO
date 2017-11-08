package Networking

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

// TCPServer creates and handles listener, later forwards data
type TCPServer struct {
	port int32
}

// CreateTCPServerOnPort creating a server on the given port
func CreateTCPServerOnPort(port int32) *TCPServer {
	server := new(TCPServer)
	server.port = port
	return server
}

// Run start a server and wait for incoming connection
func (server *TCPServer) Run() bool {
	portString := ":" + strconv.Itoa(int(server.port))
	listener, err := net.Listen("tcp", portString)
	if err != nil {
		log.Println("Failed to listen with error:", err.Error())
		return false
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Println("Failed to accept with error:", err.Error())
		return false
	}
	server.handleConnection(&conn)
	return true
}

func (server *TCPServer) handleConnection(conn *net.Conn) {
	message, _ := bufio.NewReader(*conn).ReadString('\n')
	// output message received
	fmt.Print(string(message))
}

// Shutdown server and clean up resources
func (server *TCPServer) Shutdown() bool {
	return true
}
