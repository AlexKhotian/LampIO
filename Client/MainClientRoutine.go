package Client

import (
	"LampIO/Client/Monitoring"
	"LampIO/Helper"
	"LampIO/Shared/Networking"
)

func StartClient() {
	client := Networking.CreateTCPClient(7778, "127.0.0.1")
	client.Run()

	newClientCommandHandshake := Monitoring.INewClientHandshakeFactory(client.SendCommand)
	newClientCommandHandshake.CreateAndSendNewClientCommand("TestClient")

	newModuleHandshake := Monitoring.INewModuleHandshakeFactory(client.SendCommand)
	moduleName := "TestModule"
	units := "min"
	uuid, _ := Helper.GenerateNewUUID()
	newModuleHandshake.CreateAndSendNewModuleCommand(&moduleName, uuid, &units)
	client.Shutdown()
}
