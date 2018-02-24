package Networking

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"
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
	server.incomingCmd = make(chan Command, 256)
	return server
}

// GetCommandsChan returns channel for handling of incomming commands
func (server *TCPServer) GetCommandsChan() chan Command {
	return server.incomingCmd
}

// Run start a server and wait for incoming connection
func (server *TCPServer) Run() bool {
	log.Println("Spawn routine")
	if !server.initListnerWithTLS() {
		log.Println("Failed to init a listener")
		return false
	}
	log.Println("Finished init")
	for !server.gracefulShutdown {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Println("Failed to accept with error:", err.Error())
			return false
		}
		log.Println("Got new connection")
		go server.handleConnection(&conn)
	}
	return true
}

func (server *TCPServer) initListnerWithTLS() bool {
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Println("Failed to load keys: ", err)
		return false
	}
	log.Println("Loaded certs")

	config := tls.Config{Certificates: []tls.Certificate{cert}}
	config.Rand = rand.Reader

	service := "0.0.0.0:" + strconv.Itoa(int(server.port))

	listener, err := tls.Listen("tcp", service, &config)
	server.listener = listener
	if err != nil {
		log.Println("Failed to listen with error:", err.Error())
		return false
	}
	log.Println("Created TLS listener")
	return true
}

func (server *TCPServer) handleConnection(conn *net.Conn) {
	decoder := json.NewDecoder(*conn)
	(*conn).SetReadDeadline(time.Now().Add(5 * time.Second))
	for {
		// TODO: Improve wait
		// Check if we have something really to read
		cmd := &Command{}
		err := decoder.Decode(cmd)
		if err != nil {
			if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
				log.Println("TCP timeout:", err.Error())
				return
			}
			continue
		}
		log.Println("Received new command")
		if err != nil {
			log.Println("Failed to decode incoming command", err.Error())
			continue
		}
		server.incomingCmd <- *cmd
		log.Println("Command Passed")
	}
}

// Shutdown server and clean up resources
func (server *TCPServer) Shutdown() {
	server.gracefulShutdown = false
	server.listener.Close()
}
