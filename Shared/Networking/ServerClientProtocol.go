package Networking

import "LampIO/Helper"

type NewClientCommand struct {
}

type NewModuleCommand struct {
	Name  string      `json:"Name"`
	UUID  Helper.UUID `json:"UUID"`
	Units string      `json:"Units"`
}

type ModuleDataCommand struct {
	UUID Helper.UUID `json:"UUID"`
}
