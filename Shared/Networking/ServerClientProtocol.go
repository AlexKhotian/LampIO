package Networking

import (
	"LampIO/Client/Versioning"
	"LampIO/Helper"

	"time"
)

type CommandType int32

const (
	NewClientCommandRequestType CommandType = 0
	NewModuleCommandRequestType CommandType = 1
	NewModuleDataRequestType    CommandType = 2
)

type Command struct {
	CommandType   CommandType
	MarshaledData []byte
}

type NewClientCommandRequest struct {
	Version Versioning.Version `json:"Version"`
	Name    string             `json:"Name"`
}

type NewModuleCommandRequest struct {
	Name  string      `json:"Name"`
	UUID  Helper.UUID `json:"UUID"`
	Units string      `json:"Units"`
}

type ModuleDataCommandRequest struct {
	UUID     Helper.UUID `json:"UUID"`
	NewValue float64     `json:"NewValue"`
	Time     time.Time   `json:"Time"`
}
