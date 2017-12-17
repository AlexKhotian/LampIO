package Server

import (
	"LampIO/Shared/Networking"

	"os"
	"os/signal"
	"syscall"
)

func StartServer() {
	server := Networking.CreateTCPServerOnPort(7777)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(server *Networking.TCPServer) {
		<-sigs
		server.Shutdown()
	}(server)

	server.Run()
}
