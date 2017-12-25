package Monitoring

import (
	"LampIO/Client/Versioning"
	"LampIO/Shared/Networking"
	"encoding/json"
	"errors"
)

type INewClientHandshake interface {
	CreateAndSendNewClientCommand(name *string) error
}

type newClientHandshake struct {
	sendCommand func(command *Networking.Command)
}

func INewClientHandshakeFactory(sendFunc func(command *Networking.Command)) INewClientHandshake {
	instance := new(newClientHandshake)
	instance.sendCommand = sendFunc
	return instance
}

func (instance *newClientHandshake) CreateAndSendNewClientCommand(name *string) error {
	if *name == "" {
		return errors.New("Name of new client can not be empty")
	}

	version := &Versioning.Version{
		Minor: Versioning.ClientVersionMinor,
		Major: Versioning.ClientVersionMajor}

	newClientHandshakeCmd := &Networking.NewClientCommandRequest{
		Version: *version,
		Name:    *name}

	marshaledCmd, err := json.Marshal(newClientHandshakeCmd)
	if err != nil {
		return err
	}

	cmd := &Networking.Command{
		CommandType:   Networking.NewClientCommandRequestType,
		MarshaledData: marshaledCmd}

	instance.sendCommand(cmd)

	return nil
}
