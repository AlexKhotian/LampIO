package Client

import (
	"LampIO/Client/Monitoring"
	"LampIO/Shared/Networking"
)

func StartClient() {
	client := Networking.CreateTCPClient(7778, "127.0.0.1")
	client.Run()

	newClientCommandHandshake := Monitoring.INewClientHandshakeFactory(client.SendCommand)
	newClientCommandHandshake.CreateAndSendNewClientCommand("TestClient")
	client.Shutdown()
}
