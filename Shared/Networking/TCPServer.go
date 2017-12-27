package Networking

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"log"
	"net"
	"strconv"
)

// TCPServer creates and handles listener, later forwards data
type TCPServer struct {
	port             int32
	gracefulShutdown bool
	listener         net.Listener
	incomingCmd      chan Command
}

// CreateTCPServerOnPort creating a server on the given port
func CreateTCPServerOnPort(port int32) *TCPServer {
	server := new(TCPServer)
	server.port = port
	server.gracefulShutdown = false
	server.incomingCmd = make(chan Command)
	return server
}

// GetCommandsChan returns channel for handling of incomming commands
func (server *TCPServer) GetCommandsChan() chan Command {
	return server.incomingCmd
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
	decoder := json.NewDecoder(*conn)

	cmd := &Command{}

	err := decoder.Decode(cmd)
	if err != nil {
		log.Println("Failed to decode incoming command")
		return
	}
	server.incomingCmd <- *cmd
}

// Shutdown server and clean up resources
func (server *TCPServer) Shutdown() {
	server.gracefulShutdown = false
	server.listener.Close()
}
