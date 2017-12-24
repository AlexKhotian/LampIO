package Networking

import (
	"LampIO/Helper"

	"time"
)

type CommandType int32

const (
	NewClientCommandType CommandType = 0
	NewModuleCommandType CommandType = 1
	NewModuleDataType    CommandType = 2
)

type Command struct {
	CommandType    CommandType
	MarshalledData []byte
}

type NewClientCommand struct {
}

type NewModuleCommand struct {
	Name  string      `json:"Name"`
	UUID  Helper.UUID `json:"UUID"`
	Units string      `json:"Units"`
}

type ModuleDataCommand struct {
	UUID     Helper.UUID `json:"UUID"`
	NewValue float64     `json:"NewValue"`
	Time     time.Time   `json:"Time"`
}
