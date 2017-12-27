package Networking

// IMonitoringDataSink handles transfer of monitoring data to server
type IMonitoringDataSink interface {
	CreatePacket() bool
	SendData() bool
}

type monitoringDataSink struct {
	sendCommand func()
}
