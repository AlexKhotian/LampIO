package Monitoring

import (
	"LampIO/Helper"
	"LampIO/Shared/Networking"

	"encoding/json"
	"errors"
)

// INewModuleHandshake establish connection between client and server
// If we have a new service
type INewModuleHandshake interface {
	CreateAndSendNewModuleCommand(
		moduleName *string,
		uuid Helper.UUID,
		units *string) error
}

type newModuleHandshake struct {
	sendCommand func(newModuleData *Networking.Command)
}

func INewModuleHandshakeFactory(sendCommand func(newModuleData *Networking.Command)) INewModuleHandshake {
	instance := new(newModuleHandshake)
	instance.sendCommand = sendCommand
	return instance
}

func (instance *newModuleHandshake) CreateAndSendNewModuleCommand(
	moduleName *string,
	uuid Helper.UUID,
	units *string) error {
	if *moduleName == "" {
		return errors.New("Failed to create new module cmd, empty module name")
	}

	newModuleCmd := &Networking.NewModuleCommandRequest{
		Name: *moduleName, UUID: uuid, Units: *units}

	marshalledData, err := json.Marshal(newModuleCmd)
	if err != nil {
		return err
	}

	cmd := &Networking.Command{
		CommandType:    Networking.NewModuleCommandRequestType,
		MarshalledData: marshalledData}

	instance.sendCommand(cmd)

	return nil
}
