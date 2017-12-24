package Networking

type IMonitoringDataSink interface {
	CreatePacket() bool
	SendData() bool
}

type monitoringDataSink struct {
	sendCommand func()
}
