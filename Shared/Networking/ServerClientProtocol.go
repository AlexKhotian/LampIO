package Networking

import (
	"LampIO/Client/Versioning"
	"LampIO/Helper"

	"time"
)

// CommandType type of the command. which sent betweem server and client
type CommandType int32

// Types of command
const (
	NewClientCommandRequestType CommandType = 0
	NewModuleCommandRequestType CommandType = 1
	NewModuleDataRequestType    CommandType = 2
)

// Command - top level command for communication
type Command struct {
	CommandType   CommandType `json:"CommandType"`
	MarshaledData []byte      `json:"MarshaledData"`
}

// NewClientCommandRequest signals, that we have new client
type NewClientCommandRequest struct {
	Version Versioning.Version `json:"Version"`
	Name    string             `json:"Name"`
}

// NewModuleCommandRequest signals, that we have new monitoring module
type NewModuleCommandRequest struct {
	Name  string      `json:"Name"`
	UUID  Helper.UUID `json:"UUID"`
	Units string      `json:"Units"`
}

// ModuleDataCommandRequest signals, that client sent monitoring data
type ModuleDataCommandRequest struct {
	UUID     Helper.UUID `json:"UUID"`
	NewValue float64     `json:"NewValue"`
	Time     time.Time   `json:"Time"`
}
