package Networking

// TCPServer creates and handles listener, later forwards data
type TCPServer struct {
	port int32
}

// Run start a server and wait for incoming connection
func (server *TCPServer) Run() bool {
	return true
}

// Shutdown server and clean up resources
func (server *TCPServer) Shutdown() bool {
	return true
}
