package Server

import (
	"LampIO/Shared/Networking"
)

type ICommandsDispatcherServer interface {
	RegisterForCommands(incomingCommandSignal chan Networking.Command)
	HandleCommand(cmd *Networking.Command)
	Shutdown()
}

type commandsDispatcherServer struct {
	incomingCommandSignal chan Networking.Command
}

func (instance *commandsDispatcherServer) RegisterForCommands(incomingCommandSignal chan Networking.Command) {

	instance.incomingCommandSignal = incomingCommandSignal
	select {
	case cmd := <-instance.incomingCommandSignal:
		{
			instance.HandleCommand(&cmd)
		}
	}
}

func (instance *commandsDispatcherServer) HandleCommand(cmd *Networking.Command) {

	// TODO: impl handling
	switch cmd.CommandType {
	case Networking.NewClientCommandRequestType:
	case Networking.NewModuleCommandRequestType:
	case Networking.NewModuleDataRequestType:
	}
}

func (instance *commandsDispatcherServer) Shutdown() {
	close(instance.incomingCommandSignal)
}
