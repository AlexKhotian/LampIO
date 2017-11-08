package Client

import "LampIO/Shared/Networking"

func StartClient() {
	client := Networking.CreateTCPClient(7777, "127.0.0.1")
	client.Run()
	client.SendCommand([]byte("Test"))
}
