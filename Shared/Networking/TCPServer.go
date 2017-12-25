package Networking

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strconv"
)

// TCPServer creates and handles listener, later forwards data
type TCPServer struct {
	port             int32
	gracefulShutdown bool
	listener         net.Listener
}

// CreateTCPServerOnPort creating a server on the given port
func CreateTCPServerOnPort(port int32) *TCPServer {
	server := new(TCPServer)
	server.port = port
	server.gracefulShutdown = false
	return server
}

// Run start a server and wait for incoming connection
func (server *TCPServer) Run() bool {
	if !server.initListnerWithTLS() {
		log.Println("Failed to init a listener")
		return false
	}
	for !server.gracefulShutdown {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Failed to accept with error:", err.Error())
			return false
		}
		go server.handleConnection(&conn)
	}
	return true
}

func (server *TCPServer) initListnerWithTLS() bool {
	cert, err := tls.LoadX509KeyPair("../certs/server.pem", "../certs/server.key")
	if err != nil {
		log.Println("Failed to load keys: ", err)
		return false
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader

	service := "0.0.0.0:" + strconv.Itoa(int(server.port))

	listener, err := tls.Listen("tcp", service, &config)
	server.listener = listener
	if err != nil {
		log.Println("Failed to listen with error:", err.Error())
		return false
	}
	return true
}

func (server *TCPServer) handleConnection(conn *net.Conn) {
	message, _ := bufio.NewReader(*conn).ReadString('\n')
	// output message received
	fmt.Print(string(message))
}

// Shutdown server and clean up resources
func (server *TCPServer) Shutdown() {
	server.gracefulShutdown = false
	server.listener.Close()
}
