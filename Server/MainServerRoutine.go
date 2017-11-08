package Server

import "LampIO/Shared/Networking"

func StartServer() {
	server := Networking.CreateTCPServerOnPort(7777)
	server.Run()
}
