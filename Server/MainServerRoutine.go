package Server

import (
	"LampIO/Shared/Networking"

	"os"
	"os/signal"
	"syscall"
)

// StartServer spawns new server
// main go routine on server side
func StartServer() {
	server := Networking.CreateTCPServerOnPort(7777)

	dispatcher := ICommandsDispatcherServerFactory()
	dispatcher.RegisterForCommands(server.GetCommandsChan())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(server *Networking.TCPServer) {
		<-sigs
		dispatcher.Shutdown()
		server.Shutdown()
	}(server)

	server.Run()
}
